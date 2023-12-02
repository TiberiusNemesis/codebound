package data_structures_impl

import (
	"errors"
	"fmt"
)

// Represents a node object from a doubly linked list, with pointers to previous and next ListNode elements.
// ListNode can contain any type of data.
type ListNode struct {
	data            interface{}
	nextPointer     *ListNode
	previousPointer *ListNode
}

// Represents a linked list data structure.
// Provides a head and a tail node to be used as reference of where the list begins and ends.
type LinkedList struct {
	head *ListNode
	tail *ListNode
}

// Aggregates methods used by ListNode objects.
type ListNoder interface {
	getData() (interface{}, error)
	getPreviousNode() (*ListNode, error)
	getNextNode() (*ListNode, error)
}

// Aggregates methods used by Linked List objects.
type LinkedLister interface {
	getNode(index int) (*ListNode, error)
	removeNode(node *ListNode) error
	replaceNode(index int, newNode *ListNode) error
	addNode(index int, newNode *ListNode) error
	appendNode(newNode *ListNode) error
}

// ListNoder method implementations

// Getter for the data object. Probably unnecessary.
func (node *ListNode) getData() (interface{}, error) {
	if node.data != nil {
		return node.data, nil
	}
	return nil, errors.New("node is empty")
}

// Getter for the previous node.
func (node *ListNode) getPreviousNode() (*ListNode, error) {
	if node.previousPointer != nil {
		return node.previousPointer, nil
	}
	return nil, errors.New("no previous node available")
}

// Getter for the next node.
func (node *ListNode) getNextNode() (*ListNode, error) {
	if node.nextPointer != nil {
		return node.nextPointer, nil
	}
	return nil, errors.New("no next node available")
}

// Helper method(s?)

// Links two nodes together, given their previous and next nodes.
func linkNodes(previousNode, newNode, nextNode *ListNode) {
	if previousNode != nil {
		previousNode.nextPointer = newNode
	}

	if nextNode != nil {
		nextNode.previousPointer = newNode
	}

	newNode.previousPointer = previousNode
	newNode.nextPointer = nextNode
}

// Linked List methods implemented

// Gets the node from a list based on a given index.
// Using the head of the linked list, we go over n index elements
func (list *LinkedList) getNode(index int) (*ListNode, error) {
	if index >= 0 {
		currentNode := list.head
		for i := 0; i < index; i++ {
			nextNode, err := currentNode.getNextNode()
			if err != nil {
				return nil, err
			}
			currentNode = nextNode
		}
		if currentNode == nil {
			return nil, fmt.Errorf("index out of bounds: %d", index)
		}
		return currentNode, nil
	}
	return nil, errors.New("index must be greater than or equal to 0")
}

// Removes a node from the list.
func (list *LinkedList) removeNode(node *ListNode) error {
	previousNode := node.previousPointer
	nextNode := node.nextPointer

	if previousNode == nil && nextNode == nil {
		// In this case, the node is the only one in the list
		list.head = nil
		list.tail = nil
	} else if previousNode == nil {
		// Here, our node is the head of the list
		nextNode.previousPointer = nil
		list.head = nextNode
	} else if nextNode == nil {
		// In this one, it's the tail
		previousNode.nextPointer = nil
		list.tail = previousNode
	} else {
		// And if the node is somewhere in the middle
		previousNode.nextPointer = nextNode
		nextNode.previousPointer = previousNode
	}

	return nil
}

// Replaces a node in the list with a new one.
func (list *LinkedList) replaceNode(index int, newNode *ListNode) error {
	nodeToReplace, err := list.getNode(index)
	if err != nil {
		return err
	}

	previousNode := nodeToReplace.previousPointer
	if previousNode == nil {
		// In this case, the node is probably the head of the list
		if nodeToReplace.nextPointer == nil {
			// If it's the only node, then it's also the new tail
			list.tail = newNode
		}
		// Otherwise, we just replace it and update the head
		list.head = newNode
		return nil
	} else {
		// If it's not the head, we just link the previous node to the new one
		previousNode.nextPointer = newNode
	}

	nextNode := nodeToReplace.nextPointer
	
	linkNodes(previousNode, newNode, nextNode)
	return nil
}

// Adds a new node to the list at a chosen index.
// Handles scenarios where the given index is invalid, or when the insertion is at the head.
func (list *LinkedList) addNode(index int, newNode *ListNode) error {
	if index < 0 {
		return errors.New("index must be greater than or equal to 0")
	}

	// if the index is 0, we're adding the node at the head of the list
	if index == 0 {
		newNode.nextPointer = list.head
		if list.head != nil {
			list.head.previousPointer = newNode
		}
		list.head = newNode
		if list.tail == nil {
			list.tail = newNode
		}
		return nil
	}

	previousNode, err := list.getNode(index - 1)
	if err != nil {
		return err
	}
	
	// else, we're adding the node somewhere in the middle or at the end of the list
	linkNodes(previousNode, newNode, previousNode.nextPointer)

	// If it was at the end, we update the tail
	if previousNode.nextPointer == nil {
		list.tail = newNode
	}

	return nil
}

// Adds a new node to the end of the list.
func (list *LinkedList) appendNode(newNode *ListNode) {
	if list.tail != nil {
		linkNodes(list.tail, newNode, nil)
		list.tail = newNode
	} else {
		list.head = newNode
		list.tail = newNode
	}
}
