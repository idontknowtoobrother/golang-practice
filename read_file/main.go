package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func getSourceDir() string {
	_, sourceFile, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("Failed to get source file location")
	}
	return filepath.Dir(sourceFile)
}

func getFileInSourceDir(fileName string) string {
	sourceDir := getSourceDir()
	return filepath.Join(sourceDir, fileName)
}

func readFile(filename string) (string, error) {
	fmt.Printf("Searching file %s ... \n", filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	var fileName string
	fmt.Print("Please input child file name: ")
	fmt.Scanf("%s", &fileName)

	if len(fileName) == 0 {
		fmt.Println("File name cannot be empty.")
		return
	}

	fullOpenFilePath := getFileInSourceDir(fileName)
	fileContent, err := readFile(fullOpenFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("----- %s -----\n", fileName)
	fmt.Println(fileContent)
}
