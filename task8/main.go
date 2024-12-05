// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

// InverseBit инвертирует n-ый бит числа
func InverseBit(num int64, n int) int64 {
	if GetBit(num, n) == 0 {
		// Устанавливаем i-й бит в 1 с помощью операции OR.
		num |= (1 << n)
	} else {
		// Устанавливаем i-й бит в 0 с помощью операции AND с инверсией маски.
		num &= ^(1 << n)
	}
	return num
}

// GetBit определяет n-ый бит числа num
func GetBit(num int64, n int) int {
	return int((num >> n) & 1)
}

func main() {
	var num int64 = 451
	var bitPosition int = 5

	fmt.Printf("Исходное число: %d (%b)\n", num, num)
	num = InverseBit(num, bitPosition)
	fmt.Printf("Результат: %d (%b)\n", num, num)
}
