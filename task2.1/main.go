//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из
// массива (2,4,6,8,10) и выведет их квадраты в stdout.

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
	fmt.Println(finalArray)
}
