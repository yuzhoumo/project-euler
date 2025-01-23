package lib

func Fibonacci[T GenericInt](n T) T {
	var i, prev, curr T = 0, 0, 1
	for ; i < n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

func NewFiberator[T GenericInt]() func() T {
	var prev, curr T = 0, 1
	return func() T {
		prev, curr = curr, prev+curr
		return curr
	}
}
