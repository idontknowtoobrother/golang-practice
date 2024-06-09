package main

import (
	"errors"
	"fmt"
	"log"
)

func validateAge(age int) error {
	if age < 0 || age > 150 {
		errInfo := fmt.Sprintf("invalid age: %d", age)
		return errors.New(errInfo)
	}

	return nil
}

func main() {
	if err := validateAge(-1); err != nil {
		log.Fatal(err)
		return
	}
}
