# Web Scraper with Go

A simple web scraper written in Go that fetches and displays cryptocurrency market data from the [CoinGecko API](https://www.coingecko.com/).

## 📦 Package Structure

- **main**: The entry point of the program.

## 📚 Imports

- `fmt`: For formatted I/O.
- `net/http`: To make HTTP requests.
- `encoding/json`: To decode JSON responses.
- `io`: General I/O operations.
- `github.com/olekukonko/tablewriter`: To format output in a tabular layout.
<!-- If you’re using cheerio in the future -->
<!-- - `gopkg.in/cheerio.v1`: Intended for HTML parsing (not currently used). -->

## 📐 Struct Definition

### `Coin`

```go
type Coin struct {
    ID                 string  `json:"id"`
    Symbol             string  `json:"symbol"`
    Name               string  `json:"name"`
    Image              string  `json:"image"`
    CurrentPrice       float64 `json:"current_price"`
    MarketCap          float64 `json:"market_cap"`
    MarketCapRank      int     `json:"market_cap_rank"`
    TotalVolume        float64 `json:"total_volume"`
    High24h            float64 `json:"high_24h"`
    Low24h             float64 `json:"low_24h"`
    PriceChange24h     float64 `json:"price_change_24h"`
    PriceChangePercent float64 `json:"price_change_percentage_24h"`
    CirculatingSupply  float64 `json:"circulating_supply"`
    TotalSupply        float64 `json:"total_supply"`
    ATH                float64 `json:"ath"`
    ATL                float64 `json:"atl"`
    LastUpdated        string  `json:"last_updated"`
}

🚀 Main Function

The program:

    Sends a GET request to the CoinGecko API.

    Parses the JSON response into a slice of Coin structs.

    Displays the data in a readable table format using tablewriter.

Example Output

+-----------+--------+--------------+---------------+----------+---------+----------+------------+---------------+-------------+-------------+-------------------------+
| NAME      | SYMBOL | PRICE (USD)  | MARKET CAP    | HIGH 24H | LOW 24H | CHANGE $ | CHANGE %   | CIRCULATING   | TOTAL SUP.  | ATH         | LAST UPDATED            |
+-----------+--------+--------------+---------------+----------+---------+----------+------------+---------------+-------------+-------------+-------------------------+
| Bitcoin   | BTC    | $93055.00    | $1.84T        | $96386   | $92844  | -$3053   | -3.18%     | 19800175.00   | 21000000.00 | $108135.00  | 2024-12-23T18:45:19Z    |
+-----------+--------+--------------+---------------+----------+---------+----------+------------+---------------+-------------+-------------+-------------------------+

⚠️ Error Handling

The code gracefully handles:

    HTTP request failures.

    Non-200 HTTP responses.

    JSON decoding errors.

If any error occurs, a descriptive message is printed and the program exits.
🛠 Requirements

    Go 1.18+

    Internet connection (to access the CoinGecko API)

Install dependencies:

go get github.com/olekukonko/tablewriter

🧪 Run the Project

go run main.go

