package main

import (
	_ "awesomeProject/module"
	module "awesomeProject/module"
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/jackc/pgx/v5"
	_ "time"
)

type OrderHandler struct {
	db *pgx.Conn
}

func (o OrderHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (o OrderHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (o *OrderHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var order module.Order
		if err := json.Unmarshal(message.Value, &order); err != nil {
			panic(err)
		}
		if err := o.putSql(order); err != nil {
			panic(err)
		}
		session.MarkMessage(message, "")
	}
	return nil
}
func (o *OrderHandler) putSql(order module.Order) error {
	ctx := context.Background()
	trans, err := o.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer trans.Rollback(ctx)

	_, err = trans.Exec(context.Background(), `INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature,customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		panic(err)
	}
	_, err = trans.Exec(context.Background(), `INSERT INTO deliveries (name, phone, zip, city, address, region, email)
    values ($1, $2, $3, $4, $5, $6, $7)`,
		order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		panic(err)
	}
	_, err = trans.Exec(context.Background(), `INSERT INTO payments (transaction, request_id, currency, provider,amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
    values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		panic(err)
	}
	for _, item := range order.Items {
		_, err = trans.Exec(context.Background(), `INSERT INTO items (chrt_id, track_number, price, rid, name,
					sale, size, total_price, nm_id, brand, status
				) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID,
			item.Brand, item.Status,
		)
		if err != nil {
			panic(err)
		}
	}
	if err := trans.Commit(context.Background()); err != nil {
		panic(err)
	} else {
		fmt.Println("Made write")
	}
	return nil
}

func main() {
	conf := sarama.NewConfig()
	conf.Consumer.Offsets.Initial = sarama.OffsetNewest
	conf.Consumer.Offsets.AutoCommit.Enable = false
	consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "groupTest", conf)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	ctx := context.Background()
	psql, err := pgx.Connect(ctx, "postgres://postgres:pass@localhost:4040/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	handler := &OrderHandler{db: psql}

	for {
		err := consumer.Consume(ctx, []string{"TransferData"}, handler)
		if err != nil {
			panic(err)
		}

	}

}
