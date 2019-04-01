package main

import (
	"fmt"
	"math/big"
)

func NewFibonnaci() *Fibonnaci {
	var tmp []*big.Int
	tmp = append(tmp, big.NewInt(1))
	tmp = append(tmp, big.NewInt(1))
	return &Fibonnaci{seq: tmp, current: 1}
}

type Fibonnaci struct {
	seq     []*big.Int
	current int64
}

func (f *Fibonnaci) next() {
	if f.current < 1 {
		f.seq = append(f.seq, big.NewInt(1))
	}

	f.seq = append(f.seq, big.NewInt(0).Add(f.seq[f.current-1], f.seq[f.current]))
	f.current += 1
}

func (f *Fibonnaci) get(n int64) *big.Int {
	for f.current < n {
		f.next()
	}

	return f.seq[n]
}

func find(fib *Fibonnaci, A, B string, searchingLocation *big.Int) byte {
	lenA := big.NewInt(int64(len(A)))
	lenB := big.NewInt(int64(len(B)))
	lenAB := big.NewInt(0).Add(lenA, lenB)
	if searchingLocation.Cmp(lenA) <= 0 { // searchingLocation <= lenA
		return A[searchingLocation.Uint64()-1]
	} else if searchingLocation.Cmp(lenB) <= 0 { // searchingLocation <= lenB
		return B[searchingLocation.Uint64()-1]
	} else if searchingLocation.Cmp(lenAB) <= 0 { // searchingLocation <= lenAB
		return (A + B)[searchingLocation.Uint64()-1]
	}

	// Conditions above fulfill seq 0, 1 and 2,
	// lets start with 3.
	var length big.Int
	var pos int64
	pos = 3
	for {
		// fmt.Printf(
		//   "pos: %d, fib: %d, a*: %d, b*: %d\n",
		//   pos,
		//   (fib.get(pos-2)*lenA + fib.get(pos-1)*lenB),
		//   fib.get(pos-2),
		//   fib.get(pos-1),
		// )
		length.Add(
			big.NewInt(0).Mul(fib.get(pos-2), lenA),
			big.NewInt(0).Mul(fib.get(pos-1), lenB),
		)
		// Find next sequence bigger than searchingLocation
		if length.Cmp(searchingLocation) == -1 { // length < searchingLocation
			pos += 1
			continue
		}
		break
	}

	// fmt.Printf("found pos: %d\n", pos)
	var decrease *big.Int
	decrease = big.NewInt(0)
	for {
		// It definitely is in pos = pos -2
		// Lets also correct 'searchingLocation' to proper 'pos'
		decrease.Add(
			big.NewInt(0).Mul(fib.get(pos-3), lenA),
			big.NewInt(0).Mul(fib.get(pos-2), lenB),
		)

		// fmt.Printf("pos: %d \tsearching loc.: %d \t", pos, searchingLocation)
		if decrease.Cmp(searchingLocation) >= 0 { // decrease >= searchingLocation
			pos -= 1
			// fmt.Printf("BACK decrease: %d \t\t fib-pos: %d %d\n", decrease, fib.get(pos-2), fib.get(pos-2))
			continue
		}

		searchingLocation.Sub(searchingLocation, decrease)
		// fmt.Printf("decrease: %d \t\t fib-pos: %d %d\n", decrease, fib.get(pos-3), fib.get(pos-2))
		pos -= 2
		if searchingLocation.Cmp(lenAB) <= 0 {
			// fmt.Printf("Found at location: %d\n", searchingLocation.Uint64()-1)
			if pos%2 == 0 {
				return (A + B)[searchingLocation.Uint64()-1]
			} else {
				return (B + A)[searchingLocation.Uint64()-1]
			}
		}
	}

	return 1
}

func main() {
	// Global
	fib := NewFibonnaci()
	debug := false
	if !debug {
		// IO
		var quantity int
		var A, B string
		var n string
		np := big.NewInt(0)

		fmt.Scanf("%d", &quantity)
		for i := 0; i < quantity; i++ {
			fmt.Scanf("%s", &A)
			fmt.Scanf("%s", &B)
			fmt.Scanf("%s", &n)

			np.SetString(n, 10)
			// Algorithm
			fmt.Println(string(find(fib, A, B, np)))
		}
	} else {
		np := big.NewInt(0)
		np.SetString("633825300114114700748351602688", 10) // 500
		fmt.Printf("%s\n", string(find(
			fib,
			"a",
			"b",
			np,
		)))
	}
}
