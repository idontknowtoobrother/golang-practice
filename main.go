package main

import "fmt"

func changeValueViaPointer(ref *int) {
	*ref = 200
}

func main() {
	a := 10
	fmt.Println(a)
	changeValueViaPointer(&a)
	fmt.Println(a)
}
