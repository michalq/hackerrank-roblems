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

func factorial(n int64) int64 {
	if n < 0 {
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
func binomialCoefficient(n, k int64) int64 {
	var res int64
	res = 1

	// Since C(n, k) = C(n, n-k)
	if k > n-k {
		k = n - k
	}

	// Calculate value of [n * (n-1) *---* (n-k+1)] / [k * (k-1) *----* 1]
	var i int64
	for i = 0; i < k; i++ {
		res *= (n - i)
		res /= (i + 1)
	}

	return res
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

func inArray(needle int64, haystack []int64) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}

	return false
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
			if !inArray(primeProcessor.FindNthPrime(i), divisors) {
				divisors = append(divisors, primeProcessor.FindNthPrime(i))
			}

			i = -1
		}
	}

	if !inArray(resultOfDivising, divisors) {
		divisors = append(divisors, resultOfDivising)
	}
	return divisors
}

func debug(format string, a ...interface{}) {
	// return
	fmt.Printf(format, a...)
}

func main() {
	var n, m, k int64

	// BINOMIAL COEFFICIENT LESS THAN 0 PROBLEM
	// n = 22
	// m = 6
	// k = 4

	n = 987105656
	m = 487251344
	k = 15
	/*
		304813266
		178539329
		329465650
		740794633
		570415638
		776892677
		744488030
		786963472
		27544794
		48712826
		766309953
		433855695
		442980294
		613302929
		773472163
	*/

	// fmt.Scanf("%d", &n)
	// fmt.Scanf("%d", &m)
	// fmt.Scanf("%d", &k)

	var binCoefficient int64
	binCoefficient = binomialCoefficient(n, m)
	debug("Binomial Coefficient: %d\n", binCoefficient)

	primeProcessor := NewPrime()

	var divisors [][]int64
	divisors = make([][]int64, k)
	divisors[0] = findDivisors(primeProcessor, binCoefficient)
	// Finding divisors for k(i).
	var ki int64
	var j, l, start, primeDivisorsLen int64
	var x, y, z int64
	primeDivisorsLen = int64(len(divisors[0]))
	for ki = 1; ki < k; ki++ {
		// First iteration from first to (last - 1) element.
		// First iteration use divisors from k(i - 1) and prime divisors k(0) to generate divisors for k(i).

		x, y, z = 0, 0, 0
		for j = 0; j < int64(len(divisors[ki-1])-1); j++ {
			// Second iteration from j + 1 to last element.

			start = y + ki + x
			for l = start; l < primeDivisorsLen; l++ {
				debug("M %d * %d = %d (j-%d y-%d) [Start at %d] \n", divisors[ki-1][j], divisors[0][l], divisors[ki-1][j]*divisors[0][l], j, y, start)
				divisors[ki] = append(divisors[ki], divisors[ki-1][j]*divisors[0][l])
			}

			if primeDivisorsLen == start {
				if primeDivisorsLen-ki == x {
					z++
					x = z
				} else {
					x++
				}

				debug("HOP j=%d x=%d y=%d ki=%d\n", j, x, y, ki)

				y = 0
			} else {
				y++
				debug("inc y=%d\n", y)
			}

		}
	}

	// Summing and printing.
	debug("Summing up\n")
	var i, sum int64
	for i = 0; i < k; i++ {
		sum = 0
		for _, divisor := range divisors[i] {
			// debug("S %d \n", divisor)
			sum += divisor
		}

		fmt.Println(sum)
	}
}
