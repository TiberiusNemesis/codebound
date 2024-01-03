package data_structures_impl

import (
	"testing"
)

func TestGetLastNode(t *testing.T) {
	tests := []struct {
		elements []int
		size     int
		expected int
	}{
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 2, 2},
		{[]int{1, 2, 3}, 3, 3},
		{[]int{1, 2, 3, 5}, 4, 5},
		{[]int{1, 2, 3, 8, 11}, 5, 11},
		{[]int{8, 42, 3, 11, 4, 555}, 6, 555},
		{[]int{42, 2, 1337, 4, 12, 6, 2252}, 7, 2252},
	}

	for _, test := range tests {
		tree := BuildCompleteBinaryTree(test.elements)
		lastNode := GetLastNode(tree, test.size)
		if lastNode == nil || lastNode.Key != test.expected {
			t.Errorf("GetLastNode failed for tree size %d, expected %d, got %d",
				test.size, test.expected, lastNode.Key)
		}
	}
}

func TestGetLastNode_NilRoot(t *testing.T) {
	tree := &BinaryTreeNode{}
	lastNode := GetLastNode(tree, 0)
	if lastNode != nil {
		t.Errorf("GetLastNode failed for tree size 0, expected nil, got %d",
			lastNode.Key)
	}
}

func TestRemoveLastNode(t *testing.T) {
	tests := []struct {
		elements []int
		size     int
		expected int
	}{
		{[]int{2}, 1, 0},
		{[]int{2, 5}, 2, 1},
		{[]int{2, 5, 55}, 3, 2},
		{[]int{2, 5, 55, 10}, 4, 3},
		{[]int{2, 5, 55, 10, 24}, 5, 4},
		{[]int{2, 5, 55, 10, 24, 33}, 6, 5},
		{[]int{2, 5, 55, 10, 24, 33, 1}, 7, 6},
	}

	for _, test := range tests {
		tree := BuildCompleteBinaryTree(test.elements)
		tree = RemoveLastNode(tree, test.size)
		btree := BinaryTree{Root: tree}
		newSize := btree.Size()

		if newSize != test.expected {
			t.Errorf("RemoveLastNode failed for tree size %d, expected new size %d, got %d", test.size, test.expected, newSize)
		}
	}
}

func TestRemoveLastNode_NilRoot(t *testing.T) {
	tree := &BinaryTreeNode{}
	lastNode := RemoveLastNode(tree, 0)
	if lastNode != nil {
		t.Errorf("RemoveLastNode failed for tree size 0, expected nil, got %d",
			lastNode.Key)
	}
}