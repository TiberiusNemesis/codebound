package data_structures_impl

type BinaryTree struct {
	Root *BinaryTreeNode
}

type BinaryTreeNode struct {
    Key   int
    Value interface{}
    Left  *BinaryTreeNode
    Right *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree {
    return &BinaryTree{}
}

// Size returns the number of nodes in the tree
func (t *BinaryTree) Size() int {
    return sizeNode(t.Root)
}

func sizeNode(node *BinaryTreeNode) int {
    if node == nil {
        return 0
    }
    return 1 + sizeNode(node.Left) + sizeNode(node.Right)
}


// Insert adds a new node to the tree
func (t *BinaryTree) Insert(key int, value interface{}) {
    t.Root = insertNode(t.Root, key, value)
}

func insertNode(node *BinaryTreeNode, key int, value interface{}) *BinaryTreeNode {
    if node == nil {
        return &BinaryTreeNode{Key: key, Value: value}
    }
    if key < node.Key {
        node.Left = insertNode(node.Left, key, value)
    } else if key > node.Key {
        node.Right = insertNode(node.Right, key, value)
    }
    return node
}


// Delete removes a node with a given key
func (t *BinaryTree) Delete(key int) {
    t.Root = deleteNode(t.Root, key)
}

func deleteNode(node *BinaryTreeNode, key int) *BinaryTreeNode {
    if node == nil {
        return nil
    }

    // First, find the node to be deleted.
    if key < node.Key {
        node.Left = deleteNode(node.Left, key)
    } else if key > node.Key {
        node.Right = deleteNode(node.Right, key)
    } else {
        // Node with only one child or no child
        if node.Left == nil {
            return node.Right
        } else if node.Right == nil {
            return node.Left
        }

        // Node with two children: Get the in-order successor (smallest in the right subtree)
        node.Key = minValue(node.Right)

        // Delete the in-order successor
        node.Right = deleteNode(node.Right, node.Key)
    }

    return node
}

// minValue returns the minimum key value found in that tree
func minValue(node *BinaryTreeNode) int {
    minv := node.Key
    for node.Left != nil {
        minv = node.Left.Key
        node = node.Left
    }
    return minv
}


// Find searches for a node with a given key
func (t *BinaryTree) Find(key int) (*BinaryTreeNode, bool) {
    return findNode(t.Root, key)
}

func findNode(node *BinaryTreeNode, key int) (*BinaryTreeNode, bool) {
    if node == nil {
        return nil, false
    }
    if key == node.Key {
        return node, true
    }
    if key < node.Key {
        return findNode(node.Left, key)
    }
    return findNode(node.Right, key)
}
