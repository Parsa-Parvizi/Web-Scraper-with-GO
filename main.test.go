package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestParseJSON(t *testing.T) {
	// Sample JSON for one coin
	const jsonData = `[{
		"id": "bitcoin",
		"symbol": "btc",
		"name": "Bitcoin",
		"current_price": 30000,
		"market_cap": 600000000000,
		"high_24h": 31000,
		"low_24h": 29500,
		"price_change_percentage_24h": 2.5
	}]`

	var coins []Coin
	err := json.NewDecoder(strings.NewReader(jsonData)).Decode(&coins)
	if err != nil {
		t.Errorf("Failed to parse JSON: %v", err)
	}

	if len(coins) != 1 {
		t.Errorf("Expected 1 coin, got %d", len(coins))
	}

	if coins[0].Name != "Bitcoin" {
		t.Errorf("Expected Bitcoin, got %s", coins[0].Name)
	}

	if coins[0].CurrentPrice != 30000 {
		t.Errorf("Expected price 30000, got %.2f", coins[0].CurrentPrice)
	}
}
