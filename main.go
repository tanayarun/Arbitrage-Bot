package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	
	var prices map[string]map[string]float64

	if err := json.NewDecoder(resp.Body).Decode(&prices); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	
	for coin, data := range prices {
		fmt.Printf("%s price in USD: $%.2f\n", coin, data["usd"])
	}
	
	//reader := io.NewSectionReader()
}
