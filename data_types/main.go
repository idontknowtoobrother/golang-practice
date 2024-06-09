package main

import (
	"fmt"
)

func main() {

	fmt.Println("Data Types in Go!!")
	var name string = "Jakkrit"
	var age uint8 = 24
	var isStudent bool = false // I'm a employee already!

	var dayInTheMonths uint8 = 31

	var valueOfPi float64 = 3.14159265359 // float64 is more precise than float32

	var bankBalance float64 = 1000000.00

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Is Student: ", isStudent)
	fmt.Println("Day in the months: ", dayInTheMonths)
	fmt.Println("Value of Pi: ", valueOfPi)
	fmt.Println("Bank Balance: ", bankBalance)

}
