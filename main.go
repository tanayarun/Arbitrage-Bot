package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	fmt.Println("✅ Loaded API Key and Secret (not used yet)")
	fmt.Println("API_KEY:", apiKey)
	fmt.Println("API_SECRET:", apiSecret)

	
	baseURL := "https://testnet.binancefuture.com"
	symbol := "BTCUSDT"

	client := resty.New()
	resp, err := client.R().
		Get(baseURL + "/fapi/v1/ticker/price?symbol=" + symbol)

	if err != nil {
		log.Fatal("Error fetching price:", err)
	}

	fmt.Println("✅ Latest Price of", symbol + ":")
	fmt.Println(resp.String())
}
