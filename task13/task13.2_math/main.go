package main

import "fmt"

// числа меняются местами без временной переменной при использовании арифметических действий
func changeNumbers(a, b int64) (int64, int64) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func main() {
	var (
		a int64 = 157
		b int64 = 234
	)
	fmt.Printf("before change: a = %d, b = %d\n", a, b)
	a, b = changeNumbers(a, b)
	fmt.Printf("after change: a = %d, b = %d\n", a, b)
}
