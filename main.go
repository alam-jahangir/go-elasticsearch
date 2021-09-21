package main

import (
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	// "github.com/elastic/go-elasticsearch/v8" // Depends on ES version
	// "github.com/elastic/go-elasticsearch/v8/esutil"
)

func main() {
	// Ref: https://pkg.go.dev/github.com/elastic/go-elasticsearch/esapi
	// Ref: https://pkg.go.dev/github.com/elastic/go-elasticsearch
	/*
		cfg := elasticsearch.Config{
			Addresses: []string{
				"http://localhost:9200",
				// "http://localhost:9201",
			},
		}
		es, err := elasticsearch.NewClient(cfg)
	*/
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	// log.Println(elasticsearch.Version)

	data := make(map[string]interface{})
	data["category"] = "Test"
	data["currency"] = "USD"
	data["customer_first_name"] = "Hello"
	data["customer_last_name"] = "Go"
	data["customer_full_name"] = "Hello Go"
	data["customer_gender"] = "Male"
	data["customer_id"] = 2345
	data["customer_phone"] = "P-12346"
	data["product_name"] = "Dell Inspire"
	data["quantity"] = 2
	data["price"] = 50.87
	data["min_price"] = 50.00
	data["order_date"] = time.Now()

	// res, err := es.Info()
	// if err != nil {
	// 	log.Fatalf("Error getting response: %s", err)
	// }
	// defer res.Body.Close()

	res, err := es.Index("data_ecommerce", esutil.NewJSONReader(&data), es.Index.WithDocumentID("test_data_p12346"), es.Index.WithRefresh("true"))
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	log.Println(res)
}
