package solutions

import (
	"project-euler/golang/lib"
	"strconv"
)

func E4() string {
	hi := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			product := i * j
			if lib.IsPalindromeNumber(product) && product > hi {
				hi = product
			}
		}
	}
	return strconv.Itoa(hi)
}
