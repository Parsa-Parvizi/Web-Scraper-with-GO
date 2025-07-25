package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/olekukonko/tablewriter"
	"net/http"
	"os"
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

func SaveCoinsToPostgres(coins []Coin) error {
	connStr := "host=localhost port=5432 user=coinuser password=coinpassword dbname=coindb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	createTable := `
	CREATE TABLE IF NOT EXISTS coins (
		id TEXT PRIMARY KEY,
		symbol TEXT,
		name TEXT,
		current_price REAL,
		market_cap REAL,
		high_24h REAL,
		low_24h REAL,
		price_change_percent REAL
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	for _, c := range coins {
		_, err := db.Exec(`
			INSERT INTO coins (id, symbol, name, current_price, market_cap, high_24h, low_24h, price_change_percent)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			ON CONFLICT (id) DO UPDATE SET
				current_price = EXCLUDED.current_price,
				market_cap = EXCLUDED.market_cap,
				high_24h = EXCLUDED.high_24h,
				low_24h = EXCLUDED.low_24h,
				price_change_percent = EXCLUDED.price_change_percent;
		`, c.ID, c.Symbol, c.Name, c.CurrentPrice, c.MarketCap, c.High24h, c.Low24h, c.PriceChangePercent)

		if err != nil {
			return fmt.Errorf("failed to insert data: %v", err)
		}
	}

	return nil
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
	err = SaveCoinsToPostgres(coins)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Database Error:", err)
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
