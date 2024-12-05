package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	var (
		wg sync.WaitGroup
	)

	var result atomic.Int32
	//использование атомика обеспечивает эксклюзивный доступ без состояния гонки, при этом позволяет не использовать мьютекс
	for i := range len(array) {
		wg.Add(1)
		go func(el int) {
			result.Add(int32(el * el))
			wg.Done()
		}(array[i])
	}
	// ожидание завершения всех горутин. Иначе какой-то результат может быть утерян
	wg.Wait()

	fmt.Println(result.Load())
}
