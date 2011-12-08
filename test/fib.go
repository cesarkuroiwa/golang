package test

func Fibonacci(n int) int {
	if (n == 1 || n == 2) {
		return 1
	}
	
	a, b := 1, 1
	for i := 0; i < n - 1; i++ {
		a, b = b, a + b
	}

	return a
}