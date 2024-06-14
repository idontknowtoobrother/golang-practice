package main

import "fmt"

func createCounterPtr() *int {
	counter := 0
	return &counter

}

func incrementCounterPtr(counter *int) {
	*counter++
}

func main() {

	var ptrCounter *int = createCounterPtr()
	fmt.Println(*ptrCounter)
	incrementCounterPtr(ptrCounter)
	fmt.Println(*ptrCounter)
	incrementCounterPtr(ptrCounter)
	fmt.Println(*ptrCounter)
}
