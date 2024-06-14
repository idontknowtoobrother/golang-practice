package main

import "fmt"

func updateValeByPtr(a *int) {
	*a = 40
}

func main() {
	a := 10
	fmt.Println("before: ", a)
	updateValeByPtr(&a)
	fmt.Println("after: ", a)
}
