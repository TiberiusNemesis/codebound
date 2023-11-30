package data_structures_impl

import (
	"testing"
)

func TestLinkedList_GetNode(t *testing.T) {
	list := LinkedList{}
	firstNode := &ListNode{data: "use"}
	secondNode := &ListNode{data: "the"}
	thirdNode := &ListNode{data: "force"}
	list.appendNode(firstNode)
	list.appendNode(secondNode)
	list.appendNode(thirdNode)

	// Test for getting nodes at valid indices
	expectedNode := firstNode
	actualNode, err := list.getNode(0)
	if err != nil {
		t.Errorf("Error occurred while getting node at index 0: %v", err)
	}
	if actualNode != expectedNode {
		t.Errorf("Node at index 0 should've been %v, but was %v", expectedNode, actualNode)
	}

	secondExpectedNode := secondNode
	secondActualNode, err := list.getNode(1)
	if err != nil {
		t.Errorf("Error occurred while getting node at index 1: %v", err)
	}
	if secondActualNode != secondExpectedNode {
		t.Errorf("Expected node at index 1 to be %v, but was %v", secondExpectedNode, secondActualNode)
	}

	thirdExpectedNode := thirdNode
	thirdActualNode, err := list.getNode(2)
	if err != nil {
		t.Errorf("Error occurred while getting node at index 2: %v", err)
	}
	if thirdActualNode != thirdExpectedNode {
		t.Errorf("Expected node at index 2 to be %v, but was %v", thirdExpectedNode, thirdActualNode)
	}

	// Test to get node at an invalid index
	_, err = list.getNode(3)
	if err == nil {
		t.Error("Expected error when getting node at index 3, but no error was thrown")
	}
}

func TestLinkedList_RemoveNode(t *testing.T) {
	// TODO: Write test cases for the removeNode method
}

func TestLinkedList_ReplaceNode(t *testing.T) {
	// TODO: Write test cases for the replaceNode method
}

func TestLinkedList_AddNode(t *testing.T) {
	// TODO: Write test cases for the addNode method
}

func TestLinkedList_AppendNode(t *testing.T) {
	// TODO: Write test cases for the appendNode method
}

func TestListNode_GetData(t *testing.T) {
	// TODO: Write test cases for the getData method
}

func TestListNode_GetPreviousNode(t *testing.T) {
	// TODO: Write test cases for the getPreviousNode method
}

func TestListNode_GetNextNode(t *testing.T) {
	// TODO: Write test cases for the getNextNode method
}

func TestLinkNodes(t *testing.T) {
	// TODO: Write test cases for the linkNodes function
}
