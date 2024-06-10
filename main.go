package main

import "fmt"

func main() {
	fmt.Println("Enter class and number of students (e.g., Math 30). Type 'end' to stop:")

	mapStudentsCount := make(map[string]int)
	for {
		fmt.Print("[example: Math 30] - ")
		var class string
		var studentsCount int
		_, err := fmt.Scanf("%s %d", &class, &studentsCount)
		if err != nil && class != "end" {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if class == "end" {
			break
		}

		_, exist := mapStudentsCount[class]
		if exist {
			mapStudentsCount[class] += studentsCount
			continue
		}

		mapStudentsCount[class] = studentsCount
	}

	fmt.Println("Class and number of students:")
	for class, studentsCount := range mapStudentsCount {
		fmt.Printf("%s: %d\n", class, studentsCount)
	}
}
