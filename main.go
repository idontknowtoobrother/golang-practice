package main

import (
	"fmt"
	"log"

	"github.com/idontknowtoobrother/exchange_currency_rate/db"
)

func main() {
	var amountInUSD float64
	fmt.Print("Enter amount in USD: ")
	fmt.Scanf("%f", &amountInUSD)

	var exchangeToCurrency string
	fmt.Print("Enter currency to convert to (EUR, JPY, GBP): ")
	fmt.Scanf("%s", &exchangeToCurrency)

	currencyRates, err := db.GetCurrencyRate("USD")

	if err != nil {
		log.Fatal(err)
		return
	}

	if currencyRates[exchangeToCurrency] == 0 {
		log.Fatalf("Currency rate %s not found", exchangeToCurrency)
		return
	}

	fmt.Printf("Converted amount: %.2f %s\n", amountInUSD*currencyRates[exchangeToCurrency], exchangeToCurrency)
}
