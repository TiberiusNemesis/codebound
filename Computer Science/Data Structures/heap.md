# Heap

A heap is a [complete binary tree](Computer%20Science/Data%20Structures/Tree/complete%20binary%20tree.md) sorted in one of two specific manners, each one representing the main types of heap: the *min heap* and the *max heap*. It's commonly used to store data in a way that allows us to get the min or max element in a very efficient manner.

Sometimes this [data structure](Computer%20Science/Data%20Structures/data%20structure.md) is called "binary heap", "min/max heap [binary tree](Computer%20Science/Data%20Structures/Tree/binary%20tree.md)" or "priority [queue](Computer%20Science/Data%20Structures/queue.md)".

Regarding the name priority [queue](Computer%20Science/Data%20Structures/queue.md): this is because a heap can also be technically considered a format of the [queue](Computer%20Science/Data%20Structures/queue.md) structure, but instead of ordering the [queue](Computer%20Science/Data%20Structures/queue.md) by order of insertion, it orders elements by using a priority value. A*, Prim's minimum [spanning tree](Computer%20Science/Data%20Structures/Graph/spanning%20tree.md) [algorithm](Computer%20Science/Algorithms/algorithm.md), Djikstra's [algorithm](Computer%20Science/Algorithms/algorithm.md) and the heap sort are some of the most common algorithms used in priority queues/heaps.

Insertion operations in heaps are always done by placing the newest element on the last position from top to bottom, left to right. Then, we follow up by comparing the node with its parent and based on which type of heap is being used and the difference in values from one node to another, we may swap/bubble up the child node. This operation is repeated until no more swapping is required.

## Min Heap

The root node contains the smallest (or *lowest value*) element.

The rule for a min heap's nodes is simple: the parent node's value must be lower or equal to the child node's value.

```js
// Min heap rule
value(child) >= value(parent)
```

Insert and delete operations take this into account.

When inserting an element in a min heap, we compare the node being inserted with the parent node. If it is smaller, the inserted node swaps positions with the parent node. This is repeated until the node is placed at the correct sorted position where its children have a higher value, or until it is the root node.

When deleting, we delete the smallest value (the root node) and move the last node (*i.e.* the rightmost node in the last level) of the [tree](Computer%20Science/Data%20Structures/tree.md) to the root. This acts as a pseudo-insertion, causing the now potentially misplaced elements to bubble up (if applicable), forcing a sort on the [tree](Computer%20Science/Data%20Structures/tree.md).

## Max Heap

Contrary to the min heap, in this case the root node will contain the greatest (or *highest value*) element.

Rule for a max heap's nodes is also the opposite of the min heap: the parent node's value must be **higher** or equal to the child node's value.

```js
// Max heap rule
value(parent) >= value(child)
```

Inserting into a max heap again requires the comparison of the inserted node with the parent nodes, but in this case, if the inserted node is greater than the nodes above, then it is swapped. The operation again continues until the inserted node is at the root or has a parent with a value larger than itself.

In deletion, the largest value (i.e. the root node) is removed and the last node (from top to bottom, left to right) is moved to the root. The sorting then occurs normally, as if the element was inserted in the top. The largest numbers will bubble up and the element will end up in a position where it is either a leaf node, or it is greater than its child nodes.

## Time Complexity

### Worst-case

This applies to both max and min heaps:

- Deletion: O(log n)
- Insertion: O(log n)

## Sources

- [Heap | Illustrated Data Structures](https://www.youtube.com/watch?v=F_r0sJ1RqWk)
- [Priority Queue - Introduction](https://www.coursera.org/lecture/data-structures/introduction-2OpTs)
- [Heaps and Heap Sort](https://www.youtube.com/watch?v=B7hVxCmfPtM&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=5)
