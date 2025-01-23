package solutions

import "strconv"

func allDivisible(n uint64, d uint64) bool {
	var i uint64
	for i = 2; i <= d; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func E5() string {
	const upTo uint64 = 20
	var check uint64 = upTo

	for {
		if allDivisible(check, upTo) {
			return strconv.FormatUint(check, 10)
		}
		check += upTo
	}
}
