package repository

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func GetRandomCity() (map[string]interface{}, error) {
	response, err := http.Get("http://elasticsearch1:9200/countries/_search?size=100")
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer response.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	hits, ok := data["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected format of hits")
	}

	cities := make([]map[string]interface{}, len(hits))
	for i, hit := range hits {
		cities[i] = hit.(map[string]interface{})["_source"].(map[string]interface{})
	}

	rand.Seed(time.Now().UnixNano())
	randomCity := cities[rand.Intn(len(cities))]
	return randomCity, nil
}
