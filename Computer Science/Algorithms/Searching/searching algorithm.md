# Searching

A method to find a particular element (or a group of elements) from a collection of items, or in layman's terms, getting a specific thing (or things) somewhere inside a whole lot more things.

Searching algorithms can be divided into two major categories based on their search strategy:

- Sequential search (e.g. linear search, sentinel linear search) is a type of algorithm which involves scanning each element of the collection one by one, until the desired value is found, or the entire collection has been searched. It is straightforward, but quite inefficient when used in large datasets, with a time complexity of $O(n)$, aka linear runtime.

- Interval search (e.g. binary search, ternary search, or interpolation search) represents a searching algorithm which "skips" parts of the collection in order to reach the desired element faster.  Binary Search, for example, repeatedly divides the sorted collection in half to narrow down the search area, with a time complexity of $O(log n)$. However, interval searches require the dataset to be sorted, as an unsorted collection would not allow for proper calculation of which sections/intervals the algorithm should skip.

Some other advanged searching techniques may include hashing (which allows for near-constant time searching under certain conditions) and algorithms designed for a specific data structure, like a binary search tree, a hash table, or a graph.

## Sources

- [GeeksForGeeks, Searching Algorithms](https://www.geeksforgeeks.org/searching-algorithms/)
- [freeCodeCamp, Search Algorithms – Linear Search and Binary Search Code Implementation and Complexity Analysis](https://www.freecodecamp.org/news/search-algorithms-linear-and-binary-search-explained/)
