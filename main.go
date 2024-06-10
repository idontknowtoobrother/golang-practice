package main

import (
	"fmt"
	"strings"
)

func main() {
	var sentence string
	fmt.Print("Enter sentence: ")
	fmt.Scanln(&sentence)
	words := strings.Fields(sentence)

	freqWord := make(map[string]int)

	for _, word := range words {
		freqWord[word]++
	}

	for word, freq := range freqWord {
		fmt.Println(word, freq)
	}

}
