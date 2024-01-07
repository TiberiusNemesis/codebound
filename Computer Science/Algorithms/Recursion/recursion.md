# Recursion

A process of defining a problem (or a solution to a problem) as a simpler version of itself, or many simpler versions of itself. In Computer Science, recursion is a fundamental concept which provides a very clean and elegant way to write code that solves tasks which require an undefined number of steps to reach the desired result.

A recursive function, for example, is one that may call itself during its execution to acquire a necessary value for its final result. A fibonacci sequence can be mathematically represented as $F(i) = F(i - 1) + F(i - 2)$; an equivalent function in Python can be expressed as:

```python
def fibonacci(i):
    if i in {0, 1}:
        return i
    return fibonacci(i - 1) + fibonacci(i - 2)  
```

Building a recursive algorithm requires three main steps to be defined:

1. A base case, i.e. when to stop looping/making recursive calls
2. Step(s) toward the base case
3. The recursive call.
In the above example, our base case is returning `i` once the value equals `1` or `0`. The step toward the base case would be our subtractions, `(i - 1)` and `(i - 2)`, which are embedded in the recursive calls, `fibonacci(i-1) + fibonacci(i-2)`.

## Types of Recursion

### Direct Recursion

When a function calls itself directly -- this is the simplest and most common form of recursion:

```python
def countdown(n):
    if n <= 0:
        print("Done!")
    else:
        print(n)
        countdown(n-1)
```

### Indirect Recursion

When a function calls another function, which eventually calls the first one. This forms a circular chain of calls which will eventually lead back to the original function call:

```python
def function_A(n):
    if n > 0:
        print(n)
        function_B(n-1)

def function_B(n):
    if n > 1:
        print(n)
        function_A(n/2)
```

### Nested Recursion

Occurs when a recursive function makes a recursive call acting as one of the parameters of a function that is, itself, a recursive call:

```python
def nested_recursion(n):
    if n <= 10:
        return n + 5
    else:
        return nested_recursion(nested_recursion(n-1))
```

Two other important types to be distinguished are the tail recursion and the non-tail recursion.

A common disadvantage of recursion is that it may lead to high memory usage due to extensive or multiple call stacks for each call (especially in non-tail recursion). They may also be less efficient than iterative solutions in problems that are not so easily broken down. Finally, in scenarios where proper termination (i.e. a base case) is lacking, we can encounter infinite recursion scenarios, resulting in a stack overflow error.

## Sources

- [Fireship, Recursion in 100 Seconds](https://www.youtube.com/watch?v=rf60MejMz3E)
- [wikipedia, Recursion (computer science)](https://en.wikipedia.org/wiki/Recursion_(computer_science))
- [GeeksForGeeks, Introduction to Recursion](https://www.geeksforgeeks.org/introduction-to-recursion-data-structure-and-algorithm-tutorials/)
- [GeeksForGeeks, Types of Recursions](https://www.geeksforgeeks.org/types-of-recursions/)
- [University of Utah Computer Science, Introduction - Recursion](https://users.cs.utah.edu/~germain/PPS/Topics/recursion.html#:~:text=Recursion%20means%20%22defining%20a%20problem,%2B%20F(i%2D2))
