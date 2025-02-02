package lib

import "math"

func NewPrimeIterator[T GenericInt](size T) func() T {
	values := make([]bool, size)

	var i T
	for i = 2; i < size; i++ {
		values[i] = true
	}

	var prime T = 2

	return func() T {
		for !values[prime] {
			prime++
		}

		for curr := prime + prime; curr < size; curr += prime {
			values[curr] = false
		}

		tmp := prime
		prime++

		return tmp
	}
}

func SieveOfEratosthenes[T GenericInt](size T) []bool {
	if size <= 2 {
		return make([]bool, size)
	}

	values := make([]bool, size)

	var i T
	for i = 2; i < size; i++ {
		values[i] = true
	}

	var prime T
	for prime = 2; prime < T(math.Sqrt(float64(size)))+1; prime++ {
		if !values[prime] {
			continue
		}
		for curr := prime + prime; curr < size; curr += prime {
			values[curr] = false
		}
	}

	return values
}

func IsPrime[T GenericInt](n T) bool {
	if n <= 2 {
		return n == 2
	}

	var i T
	for i = 2; i < T(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
