# Web-Scraper-with-GO

Package:

main: This code is written in the main package, which is the entry point for the program.
Imports:

fmt: This package is used for formatted printing to the console.
net/http: This package is used for making HTTP requests to web services.
encoding/json: This package is used for encoding and decoding JSON data.
io: This package provides general input/output functionality.
gopkg.in/cheerio.v1: This package (although not used in this code) is likely intended for parsing HTML documents, but it's not relevant for the current functionality that deals with JSON data.

#Struct Definition:

Coin:This struct defines the structure of a cryptocurrency with the following fields:

ID: The unique identifier for the coin (string).
Symbol: The symbol of the coin (string).
Name: The name of the coin (string).
Image: The URL of the coin's image (string).
CurrentPrice: The current price of the coin in USD (float64).
MarketCap: The market capitalization of the coin in USD (float64).
MarketCapRank: The market cap rank of the coin (int).
TotalVolume: The total trading volume of the coin in the last 24 hours (float64).
High24h: The highest price of the coin in the last 24 hours (float64).
Low24h: The lowest price of the coin in the last 24 hours (float64).
PriceChange24h: The absolute price change of the coin in the last 24 hours (float64).
PriceChangePercent: The percentage price change of the coin in the last 24 hours (float64).
CirculatingSupply: The circulating supply of the coin (float64).
TotalSupply: The total supply of the coin (float64).
ATH: The all-time high price of the coin (float64).
ATL: The all-time low price of the coin (float64).
LastUpdated: The timestamp of the last data update (string).
Main Function:

main: This is the main function where the program execution begins.
It performs the following steps:
Fetches data from the CoinGecko API endpoint using an HTTP GET request.
Checks for errors during the HTTP request and exits if an error occurs.
Verifies the HTTP status code and exits if it's not 200 (OK).
Decodes the JSON response into a slice of Coin structs.
Handles errors during JSON decoding and exits if an error occurs.
Prints a header message "Cryptocurrency Market Data".
Loops through each Coin struct in the slice.
Prints the coin's name, symbol, current price, market cap, 24h high/low, price change (absolute and percentage), circulating supply, total supply, all-time high/low, and last updated timestamp.
Prints a separator line between each coin entry.

#Error Handling:
The code incorporates error handling for both HTTP requests and JSON decoding. If an error occurs during either step, the program prints an informative message and exits.
Overall, this code demonstrates how to fetch cryptocurrency market data from the CoinGecko API using Go, parse the JSON response, and display the information in a structured format.


# Example Output
---------------------------------------------------
Cryptocurrency Market Data:
Name: Bitcoin (btc)
Current Price: $93055.00
Market Cap: $1846576757596.00
24h High: $96386.00, 24h Low: $92844.00
Price Change (24h): -$3053.31 (-3.18%)
Circulating Supply: 19800175.00
Total Supply: 21000000.00
All-Time High: $108135.00, All-Time Low: $67.81
Last Updated: 2024-12-23T18:45:19.236Z
---------------------------------------------------
