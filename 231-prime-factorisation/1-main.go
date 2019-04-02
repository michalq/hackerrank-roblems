package main

import (
	"fmt"
)

//
func factorial(n int64) int64 {
	if n < 0 {
		fmt.Println("n must be bigger than 0")
		return 0
	}

	var i, fact int64
	fact = 1
	for i = 1; i <= n; i++ {
		fact *= i
	}

	return fact
}

// See https://en.wikipedia.org/wiki/Binomial_coefficient
func binomialCoefficient(n, m int64) int64 {
	return factorial(n) / (factorial(m) * factorial(n-m))
}

func main() {
	fmt.Println("Hello")
	fmt.Println(binomialCoefficient(15, 9))
}
