// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	var (
		wg sync.WaitGroup
	)

	finalArray := make([]int, len(array))
	for i, el := range array {
		wg.Add(1)
		go func(i int) {
			// запись в слайс по определенному индексу для сохранения порядка записи квадратов
			finalArray[i] = el * el
			wg.Done()
		}(i)
	}
	// ожидание завершения всех горутин. Иначе какой-то результат может быть утерян
	wg.Wait()
	var sum int
	for _, el := range finalArray {
		sum += el
	}
	fmt.Println(sum)
}
