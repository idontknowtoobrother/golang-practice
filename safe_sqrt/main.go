package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func safeSqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New(fmt.Sprintf("cannot calculate square root of a negative number: %.2f", value))
	}

	return math.Sqrt(value), nil
}

func main() {
	result, err := safeSqrt(-1)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(result)
}
