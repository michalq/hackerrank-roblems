package main

import (
	"fmt"
	"math"
)

func NewPrime() *Prime {
	return &Prime{primes: []int64{2, 3}, lastPrimeNo: 1}
}

type Prime struct {
	primes      []int64
	lastPrimeNo int64
}

func (p *Prime) getLastPrime() int64 {
	return p.primes[p.lastPrimeNo]
}

func (p *Prime) FindNthPrime(n int64) int64 {
	if n <= p.lastPrimeNo {
		return p.primes[n]
	}
	var i, lastPrime int64
	lastPrime = p.getLastPrime()
	for i = lastPrime + 2; ; i += 2 {
		if !p.isPrime(lastPrime, i) {
			continue
		}
		p.primes = append(p.primes, i)
		p.lastPrimeNo += 1
		if p.lastPrimeNo == n {
			return i
		}
	}
}

func (p *Prime) IsPrime(n int64) bool {
	// 1. First usage will generate whole array of primes to sqrt(n)
	// 2. Next usage we will have to find density of primes density(n) which will give d
	// 3. d means that number say 5005 have x primes
	// 4. we can get primes[x-1] and if it is
	return true
}

func (p *Prime) isPrime(startOn, n int64) bool {
	if n%2 == 0 {
		return false
	}
	if startOn >= n {
		startOn = 3
	}
	var i int64
	for i = startOn; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (p *Prime) density(n int64) int64 {
	return 0
}

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

// Find prime divisors of n
func findDivisors(primeProcessor *Prime, n int64) []int64 {
	var divisors []int64
	var resultOfDivising int64
	resultOfDivising = n
	var i int64
	for i = 0; float64(i) < math.Sqrt(float64(resultOfDivising)); i++ {
		if resultOfDivising%primeProcessor.FindNthPrime(i) == 0 {
			resultOfDivising /= primeProcessor.FindNthPrime(i)
			divisors = append(divisors, primeProcessor.FindNthPrime(i))
			i = 0
		}
	}
	divisors = append(divisors, resultOfDivising)

	return divisors
}

func main() {

	var n, m, k int64
	n = 15
	m = 9
	k = 1

	var binCoefficient int64
	binCoefficient = binomialCoefficient(n, m)
	fmt.Println(binCoefficient)

	primeProcessor := NewPrime()
	divisors := findDivisors(primeProcessor, binCoefficient)

	var i int64
	var sum int64
	for i = 0; i < k; i++ {
		sum = 0
		for _, prime := range divisors {
			fmt.Printf("%d \n", prime)
			sum += prime
			// TODO find all combinations for k > 1
		}

		fmt.Println(sum)
	}
}
