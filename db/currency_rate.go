package db

import (
	"errors"
	"fmt"
)

func GetCurrencyRate(fromCurrencyPrefix string) (map[string]float64, error) { // This is simulate api response for currency rate
	if fromCurrencyPrefix == "USD" {
		return map[string]float64{
			"USD": 1.0,
			"EUR": 0.8,
			"JPY": 110.0,
			"GBP": 0.7,
		}, nil
	}
	return map[string]float64{}, errors.New(fmt.Errorf("CURRENCY RATE %s NOT FOUND", fromCurrencyPrefix).Error())
}
