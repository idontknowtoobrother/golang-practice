package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (link *LinkedList) InsertBegin(value int) {
	newNode := &Node{Value: value}
	newNode.Next = link.Head
	link.Head = newNode
}

func (link *LinkedList) InsertEnd(value int) {
	newNode := &Node{Value: value}
	// LinkedList has only head
	// if not have head we can't assign Next. Purpose of InsertEnd meaning insert to tail
	// Imaging we use LinkedList but we didn't have Head how tf we possible can assign to the end
	if link.Head == nil {
		link.Head = newNode
		return
	}
	current := link.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (link *LinkedList) Log() {
	cur := link.Head
	for cur != nil {
		fmt.Print(cur.Value, "-->")
		cur = cur.Next
	}
	fmt.Println("nil")
}

func main() {
	linkedList := &LinkedList{}
	linkedList.InsertBegin(10)
	linkedList.InsertBegin(30)
	linkedList.Log()

}
