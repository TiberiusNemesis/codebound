package main

import (
    "fmt"
    ds "github.com/TiberiusNemesis/codebound/ComputerScience/DataStructures/Implementations/Go"
)

// buildMaxHeap creates a max heap from an unsorted array
func buildMaxHeap (arr []int) {
    n := len(arr)
    for i := (n/2)-1; i >= 0; i-- {
        heapify(arr, n, i)
    }
}

// heapify is similar to build max heap, but assumes part of the array is already sorted
func heapify (arr []int, n int, i int) {
    left := 2 * i + 1
    right := 2 * i + 2
    largest := i

    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
    }
}

// Heapify recursively maintains the max-heap property starting from a given node
func Heapify(node *ds.BinaryTreeNode) {
    if node == nil {
        return
    }

    largest := node
    left := node.Left
    right := node.Right

    if left != nil && left.Key > largest.Key {
        largest = left
    }

    if right != nil && right.Key > largest.Key {
        largest = right
    }

    if largest != node {
        node.Key, largest.Key = largest.Key, node.Key
        Heapify(largest)
    }
}

// BuildMaxHeap builds a max-heap recursively using a node (ideally starting from a root node)
func BuildMaxHeap(node *ds.BinaryTreeNode) {
    if node == nil {
        return
    }
    // Perform a post-order traversal starting from the leaves and moving up.
    // This ensures that the subtree under each node is already a max-heap.
    BuildMaxHeap(node.Left)
    BuildMaxHeap(node.Right)

    // Heapify the current node to maintain the max-heap property.
    Heapify(node)
}

// InorderTraversal performs an inorder traversal of the binary tree for testing.
func InorderTraversal(node *ds.BinaryTreeNode) {
    if node == nil {
        return
    }
    InorderTraversal(node.Left)
    fmt.Printf("%d ", node.Key)
    InorderTraversal(node.Right)
}

func main() {
    tree := &ds.BinaryTree{
        Root: &ds.BinaryTreeNode{
            Key: 2,
            Left: &ds.BinaryTreeNode{
                Key: 8,
                Left: &ds.BinaryTreeNode{
                    Key: 3,
                },
                Right: &ds.BinaryTreeNode{
                    Key: 9,
                },
            },
            Right: &ds.BinaryTreeNode{
                Key: 5,
                Left: &ds.BinaryTreeNode{
                    Key: 1,
                },
            },
        },
    }

    fmt.Println("Original Tree (Inorder Traversal):")
    InorderTraversal(tree.Root)
    fmt.Println()

    BuildMaxHeap(tree.Root)

    fmt.Println("Max-Heap (Inorder Traversal after BuildMaxHeap):")
    InorderTraversal(tree.Root)
    fmt.Println()
}