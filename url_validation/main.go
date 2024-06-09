package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
)

func validateUrl(input string) error {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return errors.New("invalid URL format")
	}

	return nil
}

func main() {
	inputUrl := "https://chatgpt.com/c/3eead42e-bd05-4390-810d-db2aa12c26cb"
	err := validateUrl(inputUrl)
	if err != nil {
		log.Fatalf("%s is not valid url", inputUrl)
	}

	fmt.Printf("%s is valid url !! :D", inputUrl)
}
