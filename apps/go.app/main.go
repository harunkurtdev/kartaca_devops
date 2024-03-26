package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func staj(ctx *fiber.Ctx) error {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	req := esapi.SearchRequest{
		Index: []string{"countries"},
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error sending the request to Elasticsearch: %s", err)
	}

	defer res.Body.Close()
	if res.IsError() {
		log.Fatalf("Elasticsearch error: %s", res.Status())
	}

	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("Error parsing the Elasticsearch response: %s", err)
	}

	rand.Seed(time.Now().UnixNano())
	hits := response["hits"].(map[string]interface{})["hits"].([]interface{})
	randomIndex := rand.Intn(len(hits))
	randomCountry := hits[randomIndex].(map[string]interface{})["_source"]

	return ctx.JSON(randomCountry)
}

func hello(ctx *fiber.Ctx) error {
	// return ctx.SendString("Merhaba Go!")
	return ctx.Send([]byte("Merhaba Go!"))
}
func main() {

	var app *fiber.App = fiber.New()
	app.Use(cors.New())
	app.Get("/", hello)

	app.Get("/staj", staj)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	log.Fatal(app.Listen(":" + port))
}
