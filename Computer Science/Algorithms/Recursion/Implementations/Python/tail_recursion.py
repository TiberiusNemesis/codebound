def tail_factorial(x):
    return aux_factorial(1, x);        

def aux_factorial(y, x):
    if x == 0:
        return y
    return aux_factorial(y*x, x-1)

### The recursive tree for tail_factorial(5) is
### aux_factorial(1*5, 4)
### aux_factorial(1*4, 3)
### aux_factorial(1*3, 2)
### aux_factorial(1*2, 1)
### aux_factorial(1*1, 0)