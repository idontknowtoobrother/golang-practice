package main

import "fmt"

func main() {

	var sumOfNaturalNumbers uint16
	for i := 1; i <= 100; i++ {
		sumOfNaturalNumbers += uint16(i)
	}

	fmt.Println("Sum of first 100 natural numbers is: ", sumOfNaturalNumbers)
}
