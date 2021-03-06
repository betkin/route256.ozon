package main

import (
	"log"
	"math/big"
	"time"
)

// for timing fixation

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %vns\n", msg, time.Since(start).Nanoseconds())
}

// Simple Method

func SimpleFact(n int) *big.Int {
	defer duration(track("SimpleFact"))

	x := new(big.Int)
	x.SetInt64(1)
	xi := big.NewInt(2)

	for i := 2; i <= n; i++ {
		x.Mul(x, xi)
		xi.Add(xi, big.NewInt(1))
	}
	return x
}

// Tree method

func TreeFact(n int) *big.Int {
	defer duration(track("TreeFact"))
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}
	if n == 2 {
		return big.NewInt(int64(n))
	}
	mid := (2 + n) * 7 / 13 // to speed up the final multiple
	return new(big.Int).Mul(TreeFactRange(2, mid), TreeFactRange(mid+1, n))
}

func TreeFactRange(l int, r int) *big.Int {

	if l == r {
		return big.NewInt(int64(l))
	}
	mid := (l + r) / 2
	return new(big.Int).Mul(TreeFactRange(l, mid), TreeFactRange(mid+1, r))
}

// Sieve of Eratosthenes

func StieveFact(n int) *big.Int {
	defer duration(track("StieveFact"))
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}
	var (
		k       = 0
		simples []bool
		powers  []int64
	)

	for i := 0; i <= n; i++ {
		simples = append(simples, true)
	}
	simples[0] = false
	simples[1] = false

	for i := 2; i <= n; i++ {
		if simples[i] {
			powers = append(powers, 0)
			powers[k] += 1
			for j := i + i; j <= n; j += i {
				simples[j] = false
				for x := j; x%i == 0; x /= i {
					powers[k] += 1
				}
			}
			k++
		}
	}
	k = 0
	f := big.NewInt(1)
	for i := 2; i <= n; i++ {
		if simples[i] {
			f.Mul(f, new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(powers[k]), big.NewInt(0)))
			k++
		}
	}
	return f
}

func main() {
	const N = 50000
	SimpleFact(N)
	TreeFact(N)
	StieveFact(N)
}
