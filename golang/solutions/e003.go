package solutions

import (
	"math"
	"strconv"

	"project-euler/golang/lib"
)

func E3() string {
	target := 600851475143
	limit := uint64(math.Sqrt(float64(target))) + 1
	primes := lib.SieveOfEratosthenes[uint64](limit)
	for i := len(primes) - 1; i >= 0; i-- {
		if primes[i] && (target%i == 0) {
			return strconv.Itoa(i)
		}
	}
	return "none"
}
