package data_structures_impl

import (
	"fmt"
)

type Node struct {
	data any 
	pointerToNext *Node
	pointerToPrevious *Node	
}

type LinkedList struct {
	head *Node
	tail *Node
}

type Noder interface {
	getData() any
	getPreviousNode() *Node
	getNextNode() *Node
}

type LinkedLister interface {
	getNode(index int) *Node
	removeNode(index int) bool
	replaceNode(index int) *Node
	addNode(index int) // Node at selected index will be shifted towards the end of the list
}

// Node methods implemented
func (node *Node) getData() any {
	return node.data;
}

func (node *Node) getPreviousNode() *Node {
	fmt.Println("Printing previous node (*Node): ", node.pointerToPrevious);
	return node.pointerToPrevious;	
}

func (node *Node) getNextNode() *Node {
	fmt.Println("Printing next node (*Node): ", node.pointerToNext);
	return node.pointerToNext;	
}