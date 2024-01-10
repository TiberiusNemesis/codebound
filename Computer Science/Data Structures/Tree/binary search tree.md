# Binary Search Tree

A type of [binary tree](Computer%20Science/Data%20Structures/Tree/binary%20tree.md) which is sorted in a manner where for each node, the value of all the nodes in the left subtree have a lesser (or equal) value, and values in the right subtree always have a greater value.  

![binary-search-tree](/Assets/binary-search-tree.png)

This allows us to perform the [binary search](Computer%20Science/Algorithms/Searching/binary%20search.md) [algorithm](Computer%20Science/Algorithms/algorithm.md) very efficiently; a [binary search](Computer%20Science/Algorithms/Searching/binary%20search.md) [tree](Computer%20Science/Data%20Structures/tree.md), when also a [balanced tree](Computer%20Science/Data%20Structures/Tree/balanced%20tree.md), has a time complexity of $O(log n)$ for its insert/find/delete operations (on average).

## Operations and Time Complexity

### Worst-case

If it is not a [balanced tree](Computer%20Science/Data%20Structures/Tree/balanced%20tree.md), the BST will not have the ideal structure to perform the operations in the most efficient manner:

- Access: O(n)
- Search: O(n)
- Insertion: O(n)
- Deletion: O(n)

### Average case

If it is balanced, however, the time for operations drops considerably:

- Access: O(log n)
- Search: O(log n)
- Insertion: O(log n)
- Deletion: O(log n)

## Sources

- [Tree | Illustrated Data Structures](https://www.youtube.com/watch?v=S2W3SXGPVyU)
- [BST implementation - memory allocation in stack and heap](https://www.youtube.com/watch?v=hWokyBoo0aI&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=30)
- [Find Min and Max Element in Binary Search Tree](https://www.youtube.com/watch?v=Ut90klNN264&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=31)
- [YouTube, Data structures: Binary Search Tree](https://www.youtube.com/watch?v=pYT9F8_LFTM)
