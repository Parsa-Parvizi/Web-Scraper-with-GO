package main

import (
        "fmt"
        "net/http"
        "encoding/json"
		"io"
        "gopkg.in/cheerio.v1"
)

type Coin struct {
	ID                string  `json:"id"`
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	Image             string  `json:"image"`
	CurrentPrice      float64 `json:"current_price"`
	MarketCap         float64 `json:"market_cap"`
	MarketCapRank     int     `json:"market_cap_rank"`
	TotalVolume       float64 `json:"total_volume"`
	High24h           float64 `json:"high_24h"`
	Low24h            float64 `json:"low_24h"`
	PriceChange24h    float64 `json:"price_change_24h"`
	PriceChangePercent float64 `json:"price_change_percentage_24h"`
	CirculatingSupply  float64 `json:"circulating_supply"`
	TotalSupply        float64 `json:"total_supply"`
	ATH                float64 `json:"ath"`
	ATL                float64 `json:"atl"`
	LastUpdated        string  `json:"last_updated"`
}

func main() {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return
	}

	var coins []Coin
	err = json.NewDecoder(resp.Body).Decode(&coins)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Cryptocurrency Market Data:")
	for _, coin := range coins {
		fmt.Printf("Name: %s (%s)\n", coin.Name, coin.Symbol)
		fmt.Printf("Current Price: $%.2f\n", coin.CurrentPrice)
		fmt.Printf("Market Cap: $%.2f\n", coin.MarketCap)
		fmt.Printf("24h High: $%.2f, 24h Low: $%.2f\n", coin.High24h, coin.Low24h)
		fmt.Printf("Price Change (24h): $%.2f (%.2f%%)\n", coin.PriceChange24h, coin.PriceChangePercent)
		fmt.Printf("Circulating Supply: %.2f\n", coin.CirculatingSupply)
		fmt.Printf("Total Supply: %.2f\n", coin.TotalSupply)
		fmt.Printf("All-Time High: $%.2f, All-Time Low: $%.2f\n", coin.Ath, coin.Atl)
		fmt.Printf("Last Updated: %s\n", coin.LastUpdated)
		fmt.Println("---------------------------------------------------")
	}
}
