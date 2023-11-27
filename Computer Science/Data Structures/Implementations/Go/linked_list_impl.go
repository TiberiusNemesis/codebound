package data_structures_impl

import (
	"fmt"
	"errors"
)

// Represents a node object from a doubly linked list, with pointers to previous and next ListNode elements.
// ListNode can contain any type of data.
type ListNode struct {
	data              interface{}
	pointerToNext     *ListNode
	pointerToPrevious *ListNode
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
	if node.pointerToPrevious != nil {
		return node.pointerToPrevious, nil
	}
	return nil, errors.New("No previous node available.")
}

// Getter for the next node.
func (node *ListNode) getNextNode() (*ListNode, error) {
	if node.pointerToNext != nil {
		return node.pointerToNext, nil
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
			currentNode, _ = currentNode.getNextNode()
		}
		return currentNode, nil
	}
	return nil, errors.New("Index must be greater than or equal to 0.")
}

// Removes a node from a list,
func (list *LinkedList) removeNode(node *ListNode) error {
	previousNode, previousNodeErr := node.getPreviousNode()
	if previousNodeErr != nil {
		return fmt.Errorf("Error removing node from linked list: %v", previousNodeErr)
	}	
	
	nextNode, nextNodeErr := node.getNextNode()
	if nextNodeErr != nil {
		return fmt.Errorf("Error removing node from linked list: %v", nextNodeErr)
	}

	previousNode.pointerToNext = nextNode
	nextNode.pointerToPrevious = previousNode
	return nil
}

func (list *LinkedList) replaceNode(index int, newNode *ListNode) error {
	nodeToReplace, err := list.getNode(index)
	if err != nil {
		return err
	}

	previousNode, previousNodeErr := nodeToReplace.getPreviousNode()
	if previousNodeErr != nil {
		return fmt.Errorf("Error removing node from linked list: %v", previousNodeErr)
	}	
	
	nextNode, nextNodeErr := nodeToReplace.getNextNode()
	if nextNodeErr != nil {
		return fmt.Errorf("Error removing node from linked list: %v", nextNodeErr)
	}

	previousNode.pointerToNext = newNode
	newNode.pointerToPrevious = previousNode
	newNode.pointerToNext = nextNode
	nextNode.pointerToPrevious = newNode
	return nil
}

// Adds a new node to the list at a chosen index.
// Handles scenarios where the given index is invalid, or when the insertion is at the head.
func (list *LinkedList) addNode(index int, newNode *ListNode) error {
	if index < 0 {
		return errors.New("Index must be greater than or equal to 0.")
	}

	if index == 0 {
		newNode.pointerToNext = list.head
		if list.head != nil {
			list.head.pointerToPrevious = newNode
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

	nextNode := previousNode.pointerToNext

	previousNode.pointerToNext = newNode
	newNode.pointerToPrevious = previousNode
	newNode.pointerToNext = nextNode
	if nextNode != nil {
		nextNode.pointerToPrevious = newNode
	} else {
		list.tail = newNode
	}

	return nil
}

func (list *LinkedList) appendNode(newNode *ListNode) {
	if list.tail != nil {
		list.tail.pointerToNext = newNode
		newNode.pointerToPrevious = list.tail
		list.tail = newNode
	} else {
		list.head = newNode
		list.tail = newNode
	}
}