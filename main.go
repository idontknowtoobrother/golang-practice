package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 123, 321
	fmt.Printf("\na: %d, b: %d", a, b)
	swap(&a, &b)
	fmt.Printf("\na: %d, b: %d", a, b)

}
