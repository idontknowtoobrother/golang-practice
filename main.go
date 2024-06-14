package main

import "fmt"

func main() {

	var arayyInteger [5]int = [5]int{
		10,
		20,
		30,
		40,
		50,
	}

	arrayPtrInteger := [5]*int{}

	for i := range arayyInteger {
		arrayPtrInteger[i] = &arayyInteger[i]
	}

	for i := range arrayPtrInteger {
		fmt.Printf("Value at index %d: %d, Address: %p\n", i, *arrayPtrInteger[i], arrayPtrInteger[i])
	}
}
