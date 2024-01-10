# Linear Search

A quite straightforward approach to searching, linear search involves going through each element of the collection until the target element is found. Logically, the worst-case scenario would be when the element is at the end of the array, and we'd have to go through the entire collection in order to find it. This would result in linear runtime, aka $O(n)$.

In pseudocode, the linear search algorithm can be represented as follows:

- Start with index `0`, comparing each element with the target
- If the target is equal to the element, return its index
- Else if there are still elements left, move forward
- If the end of the list has been reached and the target is not found, return `-1`

## Sources

- [GeeksForGeeks, Linear Search Algorithm](https://www.geeksforgeeks.org/linear-search/)
- [freeCodeCamp, Search Algorithms – Linear Search and Binary Search Code Implementation and Complexity Analysis
](https://www.freecodecamp.org/news/search-algorithms-linear-and-binary-search-explained/)
