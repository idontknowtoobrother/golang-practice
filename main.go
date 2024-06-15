package main

import "fmt"

func createArray() *[5]int {
	return &[5]int{}
}

func fillArray(arrs *[5]int) {
	for i := range arrs {
		arrs[i] = i + 1
	}
}

func main() {
	arrs := createArray()
	fillArray(arrs)

	fmt.Println("Array values:", *arrs)
}
