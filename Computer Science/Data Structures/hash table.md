# Hash Table

A hash table, map, hashmap, dictionary or *associative array* are all different names that refer to the same [data structure](Computer%20Science/Data%20Structures/data%20structure.md) that handles relationships of key-value pairs and stores them into an array-like structure -- it allows for very fast access of elements since the lookup is performed by using the key. Since each key is unique and get(key) will return the value associated to that key, that makes accessing an element a constant operation (when referring to time complexity).

## How does it work?

### Key-Value Pairs

The elements in a hash table are stored in key-value pairs. These keys are unique, and are used to access the associated value.  
Here's an example:

```js
// we create a new Hash Table of size 10
var hashTable = new HashTable(10);

// then insert into it the key "my_phone" and the associated value "Samsung Galaxy"
hashTable.insert("my_phone", "Samsung Galaxy");

// to get our value, we use the get operation and pass the key
var phone = hashTable.get("my_phone")

console.log(phone) // output will be "Samsung Galaxy"
```

### Hash Function

This is a function that takes in a key and creates an index where the key-value pair will be stored in the array.  
A good hash function distributes keys uniformly across the array in order to minimize collisions (which are situations where multiple keys hash to the same index).

Example hash function to determine in which index the key-value pair should be inserted:

```js
   function hash(key) {
        if (key === null || key === undefined) {
            return 0; // Could be any other default hash value. 0 was a personal pick
        }
    
        let total = 0;
        for (let i = 0; i < key.length; i++) {
            total += key.charCodeAt(i);
        }
    
        return total % this.cells.length;
    } 
```

### Handling Collisions

When two keys hash to the same index, the hash table needs to have a strategy to handle this.  
The most common strategy is called chaining, where a linked list is used at each index.  
Another option is open addressing, where the hash table searches for the next available slot and puts the element there.

## Operations and Time Complexity

This depends on whether or not collisions will occur, which is often directly correlated to the quality of the hash function and the load factor of the hash table (i.e. the ratio of the number of elements to the size of the hash table, higher load = more collisions)

### Ideal scenario with no collisions

- Access: O(1)
- Search: O(1)
- Insertion: O(1)
- Deletion: O(1)

### Real-world scenario with collisions

- Access: Average O(1), worst O(n)
- Search: Average O(1), worst O(n)
- Insertion: Average O(1), worst O(n)
- Deletion: Average O(1), worst O(n)

## Sources

- [Hash Table | Illustrated Data Structures](https://www.youtube.com/watch?v=jalSiaIi8j4)
- [Hash Table in 4 Minutes](https://youtu.be/knV86FlSXJ8)
