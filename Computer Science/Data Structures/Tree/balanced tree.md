# Balanced Tree

A balanced [tree](Computer%20Science/Data%20Structures/tree.md) is a type of [data structure](Computer%20Science/Data%20Structures/data%20structure.md) -- specifically, a type of [binary tree](Computer%20Science/Data%20Structures/Tree/binary%20tree.md) that maintains the elements (i.e. the nodes) in a sorted order, ensuring the [tree](Computer%20Science/Data%20Structures/tree.md) remains "balanced" by not allowing large height discrepancies.

A [tree](Computer%20Science/Data%20Structures/tree.md) is considered balanced if, for every node in that [tree](Computer%20Science/Data%20Structures/tree.md), the height of the left subtree and the right subtree have a difference of *at most* one. This condition ensures that the [tree](Computer%20Science/Data%20Structures/tree.md)'s height is O(log n), where n is the number of nodes in that [tree](Computer%20Science/Data%20Structures/tree.md).

This is useful because it allows operations to be performed in logarithmic time, making them very efficient even when using large sets of data.

Balanced [tree](Computer%20Science/Data%20Structures/tree.md) subtypes include AVL trees, Red-Black trees, B-Trees, B+ Trees, and Splay Trees.

## Sources

- [Programiz, Balanced Binary Tree](https://www.programiz.com/dsa/balanced-binary-tree)
- [GeeksForGeeks, Balanced Binary Tree](https://www.geeksforgeeks.org/balanced-binary-tree/)
