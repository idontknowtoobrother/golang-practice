package main

import (
	"fmt"
)

func panicButStillExecute() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in doWork:", r)
			err = fmt.Errorf("panic occured: %v", r)
		}
	}()

	panic("Simulate panic!")
	return nil
}

func tryCatch(fn func()) (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occured: %v", r)
		}
	}()

	fn()

	return nil
}

func main() {

	fmt.Println("Start of work")
	err := panicButStillExecute()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("End of work")
	a := []int{}

	err = tryCatch(func() {
		fmt.Println(a[20])
	})
	if err != nil {
		fmt.Println("Caught error:", err)
	}

	fmt.Println(a)
}
