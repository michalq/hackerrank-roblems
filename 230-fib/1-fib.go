package main

import "fmt"

func NewFibonnaci() *Fibonnaci {
	return &Fibonnaci{seq: []uint64{1, 1}, current: 1}
}

type Fibonnaci struct {
	seq     []uint64
	current uint64
}

func (f *Fibonnaci) next() {
	if f.current < 1 {
		f.seq = append(f.seq, 1)
	}

	f.seq = append(f.seq, f.seq[f.current-1]+f.seq[f.current])
	f.current += 1
}

func (f *Fibonnaci) get(n uint64) uint64 {
	for f.current < n {
		f.next()
	}

	return f.seq[n]
}

func find(fib *Fibonnaci, A, B string, searchingLocation uint64) byte {
	lenA := uint64(len(A))
	lenB := uint64(len(B))
	lenAB := lenA + lenB
	if searchingLocation <= lenA {
		return A[searchingLocation-1]
	} else if searchingLocation <= lenB {
		return B[searchingLocation-1]
	} else if searchingLocation <= lenAB {
		return (A + B)[searchingLocation-1]
	}

	// Conditions above fulfill seq 0, 1 and 2,
	// lets start with 3.
	var pos uint64
	pos = 3
	for {
		// fmt.Printf(
		//  "pos: %d, fib: %d, a*: %d, b*: %d\n",
		//  pos,
		//  (fib.get(pos-2)*lenA + fib.get(pos-1)*lenB),
		//  fib.get(pos-2),
		//  fib.get(pos-1),
		// )
		// Find next sequence bigger than searchingLocation
		if (fib.get(pos-2)*lenA + fib.get(pos-1)*lenB) < searchingLocation {
			pos++
			continue
		}
		break
	}

	// fmt.Printf("found pos: %d\n", pos)
	var decrease uint64
	for {
		// It definitely is in pos = pos -2
		// Lets also correct 'searchingLocation' to proper 'pos'
		decrease = (fib.get(pos-2-1)*lenA + fib.get(pos-1-1)*lenB)
		fmt.Printf("pos: %d \tsearching loc.: %d \t", pos, searchingLocation)
		if decrease > searchingLocation {
			pos -= 1
			fmt.Printf("BACK to pos: %d \t\t fib-pos: %d %d\n", pos, fib.get(pos-2-1)*lenA, fib.get(pos-1-1)*lenB)

			continue
		}

		searchingLocation = searchingLocation - decrease
		fmt.Printf("decrease by: %d \t\t fib-pos: %d %d\n", decrease, fib.get(pos-2-1)*lenA, fib.get(pos-1-1)*lenB)
		pos -= 2
		if searchingLocation <= lenAB {
			if pos%2 == 0 {
				return (A + B)[searchingLocation-1]
			} else {
				return (B + A)[searchingLocation-1]
			}
		}
	}

	return 0
}

func main() {
	// Global
	fib := NewFibonnaci()
	debug := true
	if !debug {
		// IO
		var quantity int
		var A, B string
		var n uint64

		fmt.Scanf("%d", &quantity)
		for i := 0; i < quantity; i++ {
			fmt.Scanf("%s", &A)
			fmt.Scanf("%s", &B)
			fmt.Scanf("%d", &n)

			// Algorithm
			fmt.Println(string(find(fib, A, B, n)))
		}
	} else {
		// fmt.Printf("%s\n", string(find(fib, "abc", "def", 49)))
		// fmt.Printf("%s\n", string(find(
		//     fib,
		//     "1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679",
		//     "8214808651328230664709384460955058223172535940812848111745028410270193852110555964462294895493038196",
		//     104683731294243150,
		// )))

		fmt.Printf("%s\n", string(find(
			fib,
			"a",
			"b",
			50,
		)))
	}
}
