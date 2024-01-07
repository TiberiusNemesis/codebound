# Non-Tail Recursion

Basically, whenever the recursive call is *not* the last action of the function, we have an example of a non-tail recursion. This will most often lead to deeper call stacks, which can affect performance quite heavily, as seen in this non-tail recursive factorial function:

```python
def factorial(n):
    if n == 0:
        return 1
    else:
        return n * factorial(n-1)  
```

## Sources

- [GeeksForGeeks, Why is Tail Recursion optimization faster than normal Recursion?](https://www.geeksforgeeks.org/why-is-tail-recursion-optimization-faster-than-normal-recursion/)
