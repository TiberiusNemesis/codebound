# Data Structures

At its core, you could say a data structure is an organized, repeatable arrangement of data in the computer's memory. We can use these to process, retrieve and store our data in more efficient ways.

The most primitive data structures are arrangements of bytes in the memory in order to represent a value; most often, we call these _data types._ Your typical booleans, integers, floats, strings, et cetera. More complex kinds of data structures are called non-primitive. These can be further split up into two big categories: linear and non-linear.

There are many kinds of data structures, with varying capabilities, constraints and rules -- most have a unique scenario (or set of conditions) in which they can be used (or are ideally suited for), such as when implementing a specific type of [algorithm](Computer%20Science/Algorithms/algorithm.md).

A _binary [tree](Computer%20Science/Data%20Structures/tree.md),_ for example, is a non-linear data structure formed by nodes, where each of the nodes contain data, and can point to up to two other nodes. While it might be good for applying search algorithms, when we increase the size of the [tree](Computer%20Science/Data%20Structures/tree.md), suddenly it consumes a very large amount of memory. This is opposed to the [array](Computer%20Science/Data%20Structures/array.md), which would use a lower amount of memory while managing to store a greater amount of elements.

## Linear Data Structures

As the name suggests, a _linear_ data structure stores data in the computer's memory in a linear or sequential manner, where each element is attached to the next element in the structure.

Some of the most common linear data structures are

- the [array](Computer%20Science/Data%20Structures/array.md)
- the [linked list](Computer%20Science/Data%20Structures/linked%20list.md)
- the [queue](Computer%20Science/Data%20Structures/queue.md)
- the [stack](Computer%20Science/Data%20Structures/stack.md)

## Non-Linear Data Structures

A non-linear data structure, on the other hand, does not need its elements to be in a sequence. Instead, in most cases, each element in a non-linear data structure will use a _pointer_ to describe the location of another element inside the computer's memory.

Some of the most common non-linear data structures are

- the [graph](Computer%20Science/Data%20Structures/graph.md)
- the [tree](Computer%20Science/Data%20Structures/tree.md)
- the [hash table](Computer%20Science/Data%20Structures/hash%20table.md)
- the [heap](Computer%20Science/Data%20Structures/heap.md)

## Sources

- [What are Data Structures?](https://www.geeksforgeeks.org/data-structures)
- [Data Structures and Algorithms](https://www.javatpoint.com/data-structure-tutorial)
- [Data Structures Illustrated](https://www.youtube.com/watch?v=9rhT3P1MDHk&list=PLkZYeFmDuaN2-KUIv-mvbjfKszIGJ4FaY)
