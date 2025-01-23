package solutions

import "strconv"

func E1() string {
    tot := 0
    for i := 0; i < 1000; i++ {
        if (i % 3 == 0) || (i % 5 == 0) {
            tot += i
        }
    }
    return strconv.Itoa(tot)
}
