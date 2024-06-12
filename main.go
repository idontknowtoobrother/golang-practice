package main

import "fmt"

func swap(value *int, toSwapValue *int) {
	*value, *toSwapValue = *toSwapValue, *value
}

func main() {
	a, b := 10, 30
	fmt.Printf("\na=%d, b=%d", a, b)
	swap(&a, &b)
	fmt.Printf("\na=%d, b=%d", a, b)

}
