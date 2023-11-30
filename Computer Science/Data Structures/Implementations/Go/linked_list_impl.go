package data_structures_impl

import (
	"fmt"
	"errors"
)

// Represents a node object from a doubly linked list, with pointers to previous and next ListNode elements.
// ListNode can contain any type of data.
type ListNode struct {
	data              interface{}
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
	return nil, errors.New("Node is empty.")
}

// Getter for the previous node. 
func (node *ListNode) getPreviousNode() (*ListNode, error) {
	if node.previousPointer != nil {
		return node.previousPointer, nil
	}
	return nil, errors.New("No previous node available.")
}

// Getter for the next node.
func (node *ListNode) getNextNode() (*ListNode, error) {
	if node.nextPointer != nil {
		return node.nextPointer, nil
	}
	return nil, errors.New("No next node available.")
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
			return nil, fmt.Errorf("Index out of bounds: %d", index)
		}
		return currentNode, nil
	}
	return nil, errors.New("Index must be greater than or equal to 0.")
}

// Removes a node from the list.
func (list *LinkedList) removeNode(node *ListNode) error {
	previousNode := node.previousPointer
	if previousNode == nil {
		return fmt.Errorf("Error removing node from linked list: no previous node available")
	}	
	
	nextNode := node.nextPointer
	if nextNode == nil {
		return fmt.Errorf("Error removing node from linked list: no next node available")
	}

	previousNode.nextPointer = nextNode
	nextNode.previousPointer = previousNode
	return nil
}

func (list *LinkedList) replaceNode(index int, newNode *ListNode) error {
	nodeToReplace, err := list.getNode(index)
	if err != nil {
		return err
	}

	previousNode := nodeToReplace.previousPointer
	if previousNode == nil {
		return fmt.Errorf("Error removing node from linked list: no previous node available")
	}	
	
	nextNode := nodeToReplace.nextPointer
	if nextNode == nil {
		return fmt.Errorf("Error removing node from linked list: no next node available")
	}

	previousNode.nextPointer = newNode
	newNode.previousPointer = previousNode
	newNode.nextPointer = nextNode
	nextNode.previousPointer = newNode
	return nil
}

// Adds a new node to the list at a chosen index.
// Handles scenarios where the given index is invalid, or when the insertion is at the head.
func (list *LinkedList) addNode(index int, newNode *ListNode) error {
	if index < 0 {
		return errors.New("Index must be greater than or equal to 0.")
	}

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

	nextNode := previousNode.nextPointer

	previousNode.nextPointer = newNode
	newNode.previousPointer = previousNode
	newNode.nextPointer = nextNode
	if nextNode != nil {
		nextNode.previousPointer = newNode
	} else {
		list.tail = newNode
	}

	return nil
}

func (list *LinkedList) appendNode(newNode *ListNode) {
	if list.tail != nil {
		list.tail.nextPointer = newNode
		newNode.previousPointer = list.tail
		list.tail = newNode
	} else {
		list.head = newNode
		list.tail = newNode
	}
}