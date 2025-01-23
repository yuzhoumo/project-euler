package solutions

import "strconv"

func E6() string {
    const n uint64 = 100
    const sumUpTo = (n * (n + 1) / 2)
    const squareOfSums = sumUpTo * sumUpTo
    const sumOfSquares = (n * (n + 1) * (2 * n + 1)) / 6
    return strconv.FormatUint(squareOfSums - sumOfSquares, 10)
}
