package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
)

// Coin represents the structure of each cryptocurrency data from the API
type Coin struct {
	ID                 string  `json:"id"`
	Symbol             string  `json:"symbol"`
	Name               string  `json:"name"`
	CurrentPrice       float64 `json:"current_price"`
	MarketCap          float64 `json:"market_cap"`
	High24h            float64 `json:"high_24h"`
	Low24h             float64 `json:"low_24h"`
	PriceChangePercent float64 `json:"price_change_percentage_24h"`
}

// FetchCoins sends HTTP request to the API and returns the list of coins or error
func FetchCoins() ([]Coin, error) {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10&page=1"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var coins []Coin
	err = json.NewDecoder(resp.Body).Decode(&coins)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}
	return coins, nil
}

func main() {
	coins, err := FetchCoins()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Use tablewriter to print data in table format
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Symbol", "Price ($)", "Market Cap ($)", "24h High", "24h Low", "24h Change (%)"})

	for _, c := range coins {
		table.Append([]string{
			c.Name,
			c.Symbol,
			fmt.Sprintf("%.2f", c.CurrentPrice),
			fmt.Sprintf("%.0f", c.MarketCap),
			fmt.Sprintf("%.2f", c.High24h),
			fmt.Sprintf("%.2f", c.Low24h),
			fmt.Sprintf("%.2f", c.PriceChangePercent),
		})
	}

	table.Render()
}
