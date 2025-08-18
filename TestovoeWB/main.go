package main

// import (
//
//	"bufio"
//	"bytes"
//	"compress/gzip"
//	"context"
//	"fmt"
//	"github.com/PuerkitoBio/goquery"
//	"github.com/chromedp/chromedp"
//	"io"
//	"log"
//	"net/http"
//	"os"
//	"strings"
//	"time"
//
// )
//
//	type tovar struct {
//		name       string
//		price      string
//		club_price string
//		url        string
//	}
func main() {
	//	greensparkResponse()

}

//func greensparkResponse() {
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	opts := append(chromedp.DefaultExecAllocatorOptions[:],
//		chromedp.Flag("headless", true), // можно false для отладки
//		chromedp.Flag("disable-blink-features", "AutomationControlled"),
//		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
//	)
//
//	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
//	defer cancel()
//
//	browserCtx, cancel := chromedp.NewContext(allocCtx)
//	defer cancel()
//
//	url := "https://green-spark.ru/search/?q=iphone+13&sectionId=0"
//
//	var html string
//	err := chromedp.Run(browserCtx,
//		chromedp.Navigate(url),
//		chromedp.WaitReady("body", chromedp.ByQuery),
//		chromedp.Sleep(3*time.Second),
//		chromedp.OuterHTML("html", &html),
//	)
//
//	if err != nil {
//		log.Fatalf("Failed to get page: %v", err)
//	}
//
//	fmt.Println("Successfully got page content")
//	fmt.Printf(html)
//}
//func memsResponce(responce string) []tovar {
//
//	conn, err := http.NewRequest("GET", "https://memstech.ru/catalog/?q="+responce+"&set_filter=y&arrFilter_286_1205029988=Y&AJAX_CALL=N", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	client := &http.Client{
//		CheckRedirect: func(req *http.Request, via []*http.Request) error {
//			return http.ErrUseLastResponse
//		},
//	}
//	o, err := client.Do(conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//	doc, err := goquery.NewDocumentFromReader(o.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	mas := []tovar{}
//	doc.Find(".catalog-section-item-wrapper").Each(func(i int, s *goquery.Selection) {
//		if i != 0 {
//			hrefinfo, _ := s.Find("a").Attr("href")
//			a := tovar{
//				name:       strings.TrimSpace(s.Find(".intec-cl-text-hover").Text()),
//				price:      strings.TrimSpace(s.Find(".catalog-section-item-price-discount").Text()),
//				club_price: strings.TrimSpace(s.Find(".pr-promo").Text()),
//				url:        "https://memstech.ru" + hrefinfo,
//			}
//			mas = append(mas, a)
//		}
//	})
//	return mas
//}
//func masterResponce(responce string) []tovar {
//	req1, err := http.NewRequest("GET", "https://master-mobile.ru/catalog/?q="+responce, nil)
//	masterclient := &http.Client{
//		CheckRedirect: func(req *http.Request, via []*http.Request) error {
//			return http.ErrUseLastResponse
//		},
//	}
//	req1.Header.Set("Accept-Encoding", "gzip")
//	resp, err := masterclient.Do(req1)
//	if err != nil {
//		log.Fatal(err)
//	}
//	gz, err := gzip.NewReader(resp.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//fmt.Println(resp.StatusCode)
//	doc, err := goquery.NewDocumentFromReader(gz)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	mas := []tovar{}
//	doc.Find(".preview-item").Each(func(i int, s *goquery.Selection) {
//		href, _ := s.Find("a").Attr("href")
//		name := s.Find(".preview-item__name").Text()
//		price := s.Find(".payment-price__value--basic").Contents().Not("span").Text()
//		clubPrice := s.Find(".payment-price__value--club").Text()
//		url := "https://master-mobile.ru" + href
//		mas = append(mas, tovar{
//			name:       strings.TrimSpace(name),
//			price:      strings.TrimSpace(price),
//			club_price: strings.TrimSpace(clubPrice),
//			url:        strings.TrimSpace(url),
//		})
//	})
//
//	return mas
//
//}
//func mobaResponce(responce string) []tovar {
//	req, err := http.NewRequest("GET", "https://moba.ru/catalog/?q="+responce+"&s=Поиск", nil)
//	fmt.Println("Search is " + responce)
//	if err != nil {
//		log.Fatal(err)
//	}
//	headers := map[string]string{
//		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
//		"Accept-Encoding":           "gzip, deflate, br, zstd",
//		"Accept-Language":           "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7",
//		"Cache-Control":             "max-age=0",
//		"Connection":                "keep-alive",
//		"Cookie":                    "_ga= GA1.2.123456789.123456789;  _ym_visorc= w",
//		"Host":                      "moba.ru",
//		"Referer":                   "https://moba.ru/",
//		"Sec-Ch-Ua":                 `"Chromium";v="134", "Not:A-Brand";v="24", "Opera GX";v="119"`,
//		"Sec-Ch-Ua-Mobile":          "?0",
//		"Sec-Ch-Ua-Platform":        "Windows",
//		"Sec-Fetch-Dest":            "document",
//		"Sec-Fetch-Mode":            "navigate",
//		"Sec-Fetch-Site":            "same-origin",
//		"Sec-Fetch-User":            "?1",
//		"Upgrade-Insecure-Requests": "1",
//		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 OPR/119.0.0.0",
//	}
//	for key, value := range headers {
//		req.Header.Set(key, value)
//	}
//
//	client := &http.Client{
//		CheckRedirect: func(req *http.Request, via []*http.Request) error {
//			return http.ErrUseLastResponse
//		},
//	}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Body.Close()
//
//	fmt.Println("Status Code:", resp.StatusCode, resp.Cookies())
//	body, _ := io.ReadAll(resp.Body)
//	gz, err := gzip.NewReader(bytes.NewReader(body))
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer gz.Close()
//
//	doc, _ := goquery.NewDocumentFromReader(gz)
//	mas := []tovar{}
//	doc.Find(".item.main_item_wrapper").Each(func(i int, s *goquery.Selection) {
//		if i != 0 {
//			href, _ := s.Find("a").Attr("href")
//			mas = append(mas, tovar{
//				name:       strings.TrimSpace(s.Find(".item-name-cell").Find(".title").Text()),
//				price:      strings.TrimSpace(s.Find(".price-cell").Find(".price_group").Find(".price").First().Text()),
//				club_price: strings.TrimSpace(s.Find(".price-cell").Find(".price_group.min").Find(".price_value").Text() + s.Find(".price-cell").Find(".price_group.min").Find(".price_currency").Text()),
//				url:        ("https://moba.ru") + href,
//			})
//		}
//
//	})
//	return mas
//}
//func internet_цыган() {
//	results := []tovar{}
//	scanner := bufio.NewScanner(os.Stdin)
//	fmt.Print("Введите поисковый запрос: ")
//	scanner.Scan()
//	responce := scanner.Text()
//	var biolder strings.Builder
//	for _, i := range responce {
//		if i == ' ' {
//			biolder.WriteRune('+')
//		} else {
//			biolder.WriteRune(i)
//		}
//	}
//	responce = biolder.String()
//	fmt.Println("Вы ввели: ", responce)
//	results = append(results, masterResponce(responce)...)
//	results = append(results, mobaResponce(responce)...)
//	results = append(results, memsResponce(responce)...)
//	for i := 0; i < len(results); i++ {
//		fmt.Println(results[i])
//	}
//
//}
