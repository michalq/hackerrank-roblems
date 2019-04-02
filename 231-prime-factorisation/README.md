# Problem

https://www.hackerrank.com/contests/projecteuler/challenges/euler231/problem

# Dictionary

 - [Binomial Coefficient](https://en.wikipedia.org/wiki/Binomial_coefficient)
 - [Prime omega function](https://en.wikipedia.org/wiki/Prime_omega_function)

# Alghorithm

1. Calculate binomial coefficient of (N, M)
2. Find all positive divisors of bin. coefficient of (N, M) that match omega(d) = k (k = 1 .. K)
3. And the part with modulo, not so important for now

Only three steps looks pretty simple…

## Input

```N M K```

## Example

Having input

```
15 9 3
```

Steps will looks like follow

1. Binomial coefficient of (N, M) = 5005
2. for k = 1
    something1 * … * somethingx = 5005
    && omega(something1) == 1 && … && omega(somethingx) == 1
3. something1 + … + somethingx => first line of input
4. repeat from point 2 for k = 2 and k = K
