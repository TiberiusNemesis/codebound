def factorial(n):
    if n == 0:
        return 1
    else:
        return n * factorial(n-1)

### The recursive tree for factorial(5) looks something like:
###         5  factorial (4)
###       4        factorial (3)
###     3              factorial(2)
###   2                    factorial(1)
### 1                          factorial(0)