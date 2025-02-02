package solutions

import (
	"project-euler/golang/lib"
	"strconv"
)

func E2() string {
	tot := 0
	fib := lib.NewFiberator[int]()
	curr := fib()
	for curr < 4000000 {
		curr = fib()
		if curr&1 == 1 {
			continue
		}
		tot += curr
	}
	return strconv.Itoa(tot)
}
