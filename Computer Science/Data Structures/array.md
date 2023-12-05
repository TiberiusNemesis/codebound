# Array
A [data structure](Computer%20Science/Data%20Structures/data%20structure.md) which represents a fixed amount of elements that are stored sequentially in memory. Elements must have the same type in a single array, *i.e.* we can have an array which stores only numbers, or only characters, or only boolean values, and so on.

Example of an array with five elements: 
```python
[0, 1, 2, 3, 4]
```

When referring to an array through a variable, an element is accessed by including the index between brackets, like so:
```python
# array x with 5 elements
x = [10, 20, 30, 40, 50]

# selecting the third element
x[2]
```

Arrays can be *multidimensional.* This occurs when an array stores other arrays inside of it, i.e. each element of the array represents another array. Formally, it is said that *each dimension represents one index that must be specified in order to access an element* -- so when we have a 2D array, we need to specify two indexes. For a 3D array, we'd need to specify 3, and so on. Here's an example:
```python
# array a contains three arrays
a = [[10, 20, 30],[40, 50, 60],[70, 80, 90]]

# selecting the second element in the third array, aka 80
second_element_third_array = a[2][1]
```

They are very useful whenever:
- repeated or frequent access of elements by their indices is required
- the maximum number of elements is known in advance
## How does it work?
When creating an array, we allocate a fixed amount of memory for our defined number of elements. Having the same data type, the same amount of memory is allocated for each element (*e.g.* ten 32-bit integers will have each integer occupy 4 bytes, while an array of ASCII characters would have each element occupy 1 byte). 

Since these elements are sequential in memory, and they have the same size, accessing them is much faster. However, inserting a new element is a more costly operation, since it usually requires reallocating memory and moving all other elements.
## Operations and Time Complexity
### Worst-case
- Access: O(1)
- Search: O(n)
- Insertion: O(n)
- Deletion: O(n)
### Best-case
- Access: O(1)
- Search: O(1)
- Insertion: O(1)
- Deletion: O(1)
## Sources
- (YouTube) [Array Data Structure | Illustrated Data Structures](https://www.youtube.com/watch?v=QJNwK2uJyGs)
- [wiki: Array (data structure)](https://en.wikipedia.org/wiki/Array_(data_structure))
- (YouTube) [Jagged Arrays](https://www.youtube.com/watch?v=1jtrQqYpt7g)
- [Arrays, GeeksForGeeks](https://www.geeksforgeeks.org/array-data-structure/)