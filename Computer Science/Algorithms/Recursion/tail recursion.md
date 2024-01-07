# Tail Recursion

When the recursive call is the last operation in the function, we call it a tail [recursion](Computer%20Science/Algorithms/Recursion/recursion.md). Some languages support tail call optimization (aka TCO), which leads to tail recursive functions being as efficient as iterative loops. Here's an example of a factorial function with tail [recursion](Computer%20Science/Algorithms/Recursion/recursion.md):

```python
def tail_factorial(x):
    return aux_factorial(1, x);        

def aux_factorial(y, x):
    if x == 0:
        return y
    return aux_factorial(y*x, x-1)
```

## Sources

- [GeeksForGeeks, What is Tail Recursion?](https://www.geeksforgeeks.org/tail-recursion/)
- [StackOverflow, What is tail recursion?](https://stackoverflow.com/questions/33923/what-is-tail-recursion)
- [GeeksForGeeks, Why is Tail Recursion optimization faster than normal Recursion?](https://www.geeksforgeeks.org/why-is-tail-recursion-optimization-faster-than-normal-recursion/)
