package data_structures_impl

import (
	"testing"
	"errors"
)

func TestGetData(t *testing.T) {
	// Create a new linked list
	ll := LinkedList{}
	// Add some nodes to the linked list
	ll.addNode(0, &ListNode{0, nil, nil})
	ll.addNode(1, &ListNode{1, nil, nil})
	ll.addNode(2, &ListNode{2, nil, nil})

	// Test getting data from the nodes
	node, err := ll.getNode(0)
	if err != nil {
		t.Errorf("Error getting node from list: %v", err)
	}
	data, err := node.getData()
	if data != 0 {
		t.Errorf("Expected data to be 0, got %v", data)
	}

	data, err = ll.getNode(1)
	if err != nil {
		t.Errorf("Error getting node from list: %v", err)
	}

	if data != 1 {
		t.Errorf("Expected data to be 1, got %v", data)
	}

	data, err = ll.getNode(2).getData()
	if err != nil {
		t.Errorf("Error getting data from node: %v", err)
	}
	if data != 2 {
		t.Errorf("Expected data to be 2, got %v", data)
	}
}

func TestRemoveNode(t *testing.T) {
	// Create a new linked list
	ll := LinkedList{}

	// Add some nodes to the linked list
	ll.addNode(0)
	ll.addNode(1)
	ll.addNode(2)

	// Remove a node from the linked list
	removed := ll.removeNode(1)
	if !removed {
		t.Errorf("Expected node to be removed")
	}

	// Test getting data from the remaining nodes
	data, err := ll.getNode(0).getData()
	if err != nil {
		t.Errorf("Error getting data from node: %v", err)
	}
	if data != 0 {
		t.Errorf("Expected data to be 0, got %v", data)
	}

	data, err = ll.getNode(1).getData()
	if err != nil {
		t.Errorf("Error getting data from node: %v", err)
	}
	if data != 2 {
		t.Errorf("Expected data to be 2, got %v", data)
	}
}

func TestReplaceNode(t *testing.T) {
	// Create a new linked list
	ll := LinkedList{}

	// Add some nodes to the linked list
	ll.addNode(0)
	ll.addNode(1)
	ll.addNode(2)

	// Replace a node in the linked list
	replaced := ll.replaceNode(1)
	if replaced == nil {
		t.Errorf("Expected node to be replaced")
	}

	// Test getting data from the replaced node
	data, err := replaced.getData()
	if err != nil {
		t.Errorf("Error getting data from node: %v", err)
	}
	if data != 1 {
		t.Errorf("Expected data to be 1, got %v", data)
	}
}

func TestAddNode(t *testing.T) {
	// Create a new linked list
	ll := LinkedList{}

	// Add some nodes to the linked list
	ll.addNode(0)
	ll.addNode(1)
	ll.addNode(2)

	// Add a new node to the linked list
	ll.addNode(1)

	// Test getting data from the nodes
	data, err := ll.getNode(0).getData()
	if err != nil {
		t.Errorf("Error getting data from node: %v", err)
	}
	if data != 0 {
		t.Errorf("Expected data to be 0, got %v", data)
	}

	data, err = ll.getNode(1).getData()
	if err != nil {
		t.Errorf("Error getting data from node: %v", err)
	}
	if data != 1 {
		t.Errorf("Expected data to be 1, got %v", data)
	}

	data, err = ll.getNode(2).getData()
	if err != nil {
		t.Errorf("Error getting data from node: %v", err)
	}
	if data != 3 {
		t.Errorf("Expected data to be 3, got %v", data)
	}

	data, err = ll.getNode(3).getData()
	if err != nil {
		t.Errorf("Error getting data from node: %v", err)
	}
	if data != 2 {
		t.Errorf("Expected data to be 2, got %v", data)
	}
}