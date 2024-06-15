package main

import "fmt"

func increamentCount(mapItem map[string]*int, itemName string) {
	if _, exist := mapItem[itemName]; !exist {
		count := 0
		mapItem[itemName] = &count
	}

	*mapItem[itemName]++
}

func main() {
	mapItem := make(map[string]*int)

	increamentCount(mapItem, "apple")
	increamentCount(mapItem, "apple")
	increamentCount(mapItem, "apple")
	increamentCount(mapItem, "apple")
	increamentCount(mapItem, "orange")
	increamentCount(mapItem, "orange")

	for itemName, count := range mapItem {
		fmt.Printf("\nitem: %s, count: %d", itemName, *count)
	}
}
