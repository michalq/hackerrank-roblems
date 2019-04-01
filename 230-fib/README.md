# Problem

https://www.hackerrank.com/contests/projecteuler/challenges/euler230/problem

# Files
 - 1-fib.go Made on uint64
 - 2-fib.go Same algorithm made on big.Int

# Alghoritm

Looking position: 49

```
abc
def
abcdef
defabcdef
abcdefdefabcdef
defabcdefabcdefdefabcdef
         .
abcdefdefabcdefdefabcdefabcdefdefabcdef
defabcdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdef < STOP at n, CUT using n-1, MOVE LOC to new
                                                .
abcdefdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdef
defabcdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdef
abcdefdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdefabcdefdefabcdefdefabcdefabcdefdefabcdef
```

# Tests
```
a
b // And we should end up here?
ab
bab // decrease < searchLoc, fine?
  .
abbab    // ERROR!
  .
bababbab // ERROR! decrease > searchingLoc, instead of searchingLoc -= decrease we should pos -= 2 and searchingLoc = decrease
  .
abbabbababbab
bababbababbabbababbab // decrease < searchLoc, fine
               .
abbabbababbabbababbababbabbababbab
bababbababbabbababbababbabbababbabbababbababbabbababbab // decrease < searchLoc, fine
                                                 .
abbabbababbabbababbababbabbababbabbababbababbabbababbababbabbababbabbababbababbabbababbab
```
---
```
a
b
ab
bab
abbab
bababbab
abbabbababbab
bababbababbabbababbab
             .
abbabbababbabbababbababbabbababbab // Cannot decrease pos by 2, lets dec by 1 and check again
             .
bababbababbabbababbababbabbababbabbababbababbabbababbab
abbabbababbabbababbababbabbababbabbababbababbabbababbababbabbababbabbababbababbabbababbab  // decrease < searchLoc, fine
                                                                    .
```