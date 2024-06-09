package main

import (
	"fmt"
	"time"
)

func main() {
	// IF ELSE
	var inputYear int
	fmt.Print("Enter a year: ")
	fmt.Scanf("%d", &inputYear)

	// Get this year
	year, _, _ := time.Now().Date()

	if inputYear == year {
		fmt.Printf("You're in present year!!")
	} else {
		fmt.Printf("You're not in present year!!")
	}

	// IF-ELSEIF
	fmt.Print("\nEnter a year again: ")
	fmt.Scanf("%d", &inputYear)

	if inputYear == year {
		fmt.Printf("%d is present year!!", inputYear)
	} else if inputYear > year {
		fmt.Printf("%d is future year!! You're mtfk teleporter!!", inputYear)
	} else {
		fmt.Printf("%d is past year!!", inputYear)
	}

	// IF
	var inputPresentYear int
	fmt.Print("\nEnter a present year only: ")
	fmt.Scanf("%d", &inputPresentYear)

	if inputPresentYear == year {
		fmt.Printf("%d is correct!!", inputPresentYear)
		return // Exit the program
	}

	fmt.Printf("%d is not correct!! this year is %d", inputPresentYear, year)
}
