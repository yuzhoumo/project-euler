package solutions

import (
	"project-euler/golang/lib"
	"strconv"
)

func E7() string {
	iter := lib.NewPrimeIterator[uint64](150000)
	for i := 0; i < 10000; i++ {
		iter()
	}
	return strconv.FormatUint(iter(), 10)
}
