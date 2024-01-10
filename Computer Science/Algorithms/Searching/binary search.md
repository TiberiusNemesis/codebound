# Binary Search

A search algorithm where, in a sorted array, the search interval is repeatedly divided in half, continuously approximating the area of the element to be queried until it is found. With this approach, each iteration means half of the remaining values are eliminated, leading to a 99% reduction of the search scope within just 7 steps.

Binary search's worst-case execution is $O(log n)$, aka logarithmic runtime. Interestingly, this can be seen in a visual example when transforming an array into a binary tree, with the steps to find a certain element being visualized as the average nodes to either the left or the right of the middle element.

Below is the array `[1, 4, 5, 7, 9, 12, 15, 18, 19, 22, 25, 29, 40, 50]`, represented as a binary tree. Each level down represents one more step that the binary search algorithm would need to take in order to find the element in that node.

![Binary search tree representation of an array.](/Assets/array-as-bst.png)

## Sources

- [Michael Sambol (YouTube), Binary search in 4 minutes](https://www.youtube.com/watch?v=fDKIpRe8GW4)
- [wikipedia, Binary search algorithm](https://en.wikipedia.org/wiki/Binary_search_algorithm)
- [GeeksForGeeks, Binary Search](https://www.geeksforgeeks.org/binary-search/)
