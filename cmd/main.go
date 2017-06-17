package main

import (
    "fmt"
    "github.com/mvaude/exchange/pkg/go-bittrex"
)

const (
    API_KEY    = "API_KEY"
    API_SECRET = "API_SECRET"
)

func main() {
    // Bittrex client
    bittrex := bittrex.New(API_KEY, API_SECRET)

    // Get markets
    markets, err := bittrex.GetMarkets()
    fmt.Println(err, markets)
}
