# Searching

A method to find a particular element (or a group of elements) from a collection of items, or in layman's terms, getting a specific thing (or things) somewhere inside a whole lot more things.

Searching algorithms can be divided into two major categories based on their search strategy:

- Sequential search (e.g. [linear search](Computer%20Science/Algorithms/Searching/linear%20search.md), sentinel [linear search](Computer%20Science/Algorithms/Searching/linear%20search.md)) is a type of [algorithm](Computer%20Science/Algorithms/algorithm.md) which involves scanning each element of the collection one by one, until the desired value is found, or the entire collection has been searched. It is straightforward, but quite inefficient when used in large datasets, with a time complexity of $O(n)$, aka [linear runtime](Computer%20Science/Asymptotic%20Notation/Common%20Runtimes/linear%20runtime.md).

- Interval search (e.g. [binary search](Computer%20Science/Algorithms/Searching/binary%20search.md), ternary search, or interpolation search) represents a searching [algorithm](Computer%20Science/Algorithms/algorithm.md) which "skips" parts of the collection in order to reach the desired element faster.  Binary Search, for example, repeatedly divides the sorted collection in half to narrow down the search area, with a time complexity of $O(log n)$. However, interval searches require the dataset to be sorted, as an unsorted collection would not allow for proper calculation of which sections/intervals the [algorithm](Computer%20Science/Algorithms/algorithm.md) should skip.

Some other advanged searching techniques may include hashing (which allows for near-constant time searching under certain conditions) and algorithms designed for a specific [data structure](Computer%20Science/Data%20Structures/data%20structure.md), like a [binary search tree](Computer%20Science/Data%20Structures/Tree/binary%20search%20tree.md), a [hash table](Computer%20Science/Data%20Structures/hash%20table.md), or a [graph](Computer%20Science/Data%20Structures/graph.md).

## Sources

- [GeeksForGeeks, Searching Algorithms](https://www.geeksforgeeks.org/searching-algorithms/)
- [freeCodeCamp, Search Algorithms – Linear Search and Binary Search Code Implementation and Complexity Analysis](https://www.freecodecamp.org/news/search-algorithms-linear-and-binary-search-explained/)
