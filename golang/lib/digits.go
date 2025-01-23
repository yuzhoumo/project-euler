package lib

func IsPalindromeNumber[T GenericInt](n T) bool {
    if n < 0 {
        return false
    }

    var tmp, rev T = n, 0

    for tmp > 0 {
        rev *= 10
        rev += tmp % 10
        tmp /= 10
    }

    return rev == n
}
