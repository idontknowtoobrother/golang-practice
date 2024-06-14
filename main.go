package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {

	aNode := &Node{Value: 10}
	bNode := &Node{Value: 20}
	cNode := &Node{Value: 30}
	aNode.Next = bNode
	bNode.Next = cNode

	fmt.Println(aNode.Value)
	if aNode.Next != nil {
		fmt.Println(aNode.Next.Value)
		if bNode.Next != nil {
			fmt.Println(bNode.Next.Value)
		}
	}
}
