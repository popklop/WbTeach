package main

import (
	_ "awesomeProject/docs"
	models "awesomeProject/module"
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var lastResponceChache = make(map[string]models.Order)
var queue = make([]string, 0, 5)
var mu = &sync.Mutex{}
var chache = make(map[string]models.Order)

// @title Orders API
// @version 1.0
// @description API для управления заказами
// @host localhost:8080
// @BasePath /
func main() {
	db, err := sql.Open("pgx", "postgres://postgres:pass@localhost:4040/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.GET("/order/:id", getOrder(db))
	router.Static("/static", "./UserInterface")
	router.POST("/createorder", createOrder(db))

	go func() {
		time.Sleep(100 * time.Millisecond)
		callendpints()
	}()

	log.Fatal(router.Run(":8080"))
}

// GetOrder godoc
// @Summary Получить заказ по ID
// @Description Получить детали заказа по его уникальному идентификатору
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order UID"
// @Success 200 {object} models.Order
// @Failure 404 {object} object{error=string}
// @Failure 500 {object} object{error=string}
// @Router /order/{id} [get]
func getOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var order models.Order
		mu.Lock()
		order, exists := lastResponceChache[id]
		mu.Unlock()
		if !exists {
			mu.Lock()
			order, exists = chache[id]
			mu.Unlock()
			if !exists {
				err := db.QueryRow("SELECT order_uid, track_number FROM orders WHERE order_uid = $1", id).Scan(
					&order.OrderUID,
					&order.TrackNumber,
				)
				if err != nil {
					c.JSON(404, gin.H{"error": "Order not found"})
					return
				}
			}
			mu.Lock()
			defer mu.Unlock()
			if len(queue) >= 5 {
				oldestID := queue[0]
				delete(lastResponceChache, oldestID)
				queue = queue[1:]
			}
			lastResponceChache[id] = order
			queue = append(queue, id)
			chache[id] = order
		}
		endpoint := c.Request.URL.Path
		writecache(endpoint)
		c.JSON(200, gin.H{
			"order_uid":    order.OrderUID,
			"track_number": order.TrackNumber,
		})
	}
}

// CreateOrder godoc
// @Summary Создать новый заказ
// @Description Добавляет заказ в очередь Kafka
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Данные заказа"
// @Success 201 {object} models.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /createorder [post]
func createOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(400, gin.H{"error": "Невалидные данные"})
			return
		}

		trans, err := db.Begin()
		if err != nil {
			panic(err)
		}
		defer trans.Rollback()

		_, err = trans.Exec(

			`INSERT INTO orders (
				order_uid, track_number, entry, locale, internal_signature,
				customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			order.OrderUID,
			order.TrackNumber,
			order.Entry,
			order.Locale,
			order.InternalSignature,
			order.CustomerID,
			order.DeliveryService,
			order.ShardKey,
			order.SmID,
			order.DateCreated,
			order.OofShard,
		)
		if err != nil {
			panic(err)
		}
		_, err = trans.Exec(`INSERT INTO deliveries (name, phone, zip, city, address, region, email)
    values ($1, $2, $3, $4, $5, $6, $7)`,
			order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
		if err != nil {
			panic(err)
		}
		_, err = trans.Exec(`INSERT INTO payments (transaction, request_id, currency, provider,amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
    values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
			order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
		if err != nil {
			panic(err)
		}
		for _, item := range order.Items {
			_, err = trans.Exec(`INSERT INTO items (chrt_id, track_number, price, rid, name,
					sale, size, total_price, nm_id, brand, status
				) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
				item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID,
				item.Brand, item.Status)
			if err != nil {
				panic(err)
			}
		}
		if err := trans.Commit(); err != nil {
			panic(err)
		} else {
			fmt.Println("Made write")
		}
		if err != nil {
			c.JSON(500, gin.H{"error": "Ошибка при создании заказа"})
			log.Printf("500", err)
			return
		}

		c.JSON(201, order) // 201 Created
	}
}
func writecache(endpoint string) {
	existingEndpoints := make(map[string]bool)
	if file, err := os.Open("InitData.txt"); err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			existingEndpoints[scanner.Text()] = true
		}
		file.Close()
	} else if !os.IsNotExist(err) {
		log.Printf("Ошибка чтения файла: %v", err)
		return
	}
	if existingEndpoints[endpoint] {
		return
	}
	file, err := os.OpenFile("InitData.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Ошибка открытия файла: %v", err)
		return
	}
	defer file.Close()
	if _, err := file.WriteString(endpoint + "\n"); err != nil {
		log.Printf("Ошибка записи в файл: %v", err)
	}
}
func callendpints() {
	wg := sync.WaitGroup{}
	file, err := os.Open("InitData.txt")
	if err != nil {
		panic(err)
	}
	endpoints := make([]string, 0)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		endpoints = append(endpoints, scanner.Text())
	}
	for _, endpoint := range endpoints {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get("http://127.0.0.1:8080" + url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			order := &models.Order{}
			if err := json.Unmarshal(body, order); err != nil {
				panic(err)

			}
			mu.Lock()
			chache[order.OrderUID] = *order
			mu.Unlock()
		}(endpoint)
	}
	wg.Wait()
	log.Println("c", chache)
}
