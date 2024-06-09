package main

import (
	"errors"
	"fmt"
	"log"
)

func safeDivide(numerator float64, denominator float64) (float64, error) {
	if denominator <= 0.0 {
		return 0.0, errors.New("division by zero is not allowed")
	}

	return numerator / denominator, nil
}

func main() {

	result, err := safeDivide(20, 0.0)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("The result is %.2f\n", result)
}
