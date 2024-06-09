package main

import (
	"errors"
	"fmt"
	"log"
)

func validateUsername(username string) error {

	usernameLength := len(username)
	if usernameLength < 5 || usernameLength > 15 {
		return errors.New(fmt.Sprintf("username must be loger than 5 and not more than 15 charactes: %s", username))
	}

	return nil
}

func main() {

	validateUsername("morethanfivem")             // not error
	err := validateUsername("morethanfifteenJAA") // should not pass
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	if err = validateUsername("less"); err != nil { // should not pass
		log.Fatalf("Error: %s", err)
	}
}
