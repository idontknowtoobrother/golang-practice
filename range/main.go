package main

import "fmt"

func main() {

	var fiveNumbers [5]int

	fiveNumbers[0] = 1
	fiveNumbers[1] = 2
	fiveNumbers[2] = 3
	fiveNumbers[3] = 4
	fiveNumbers[4] = 5

	var sumOfNumbers int = 0

	for _, value := range fiveNumbers {
		sumOfNumbers += value
	}

	var averageOfFiveNumbers float64 = float64(sumOfNumbers) / float64(len(fiveNumbers))

	fmt.Println("The sum of the numbers is: ", sumOfNumbers)
	fmt.Printf("The average of the numbers is: %.2f\n", averageOfFiveNumbers)

}
