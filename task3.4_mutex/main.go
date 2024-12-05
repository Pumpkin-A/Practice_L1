package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	var (
		wg sync.WaitGroup
		m  sync.Mutex
	)

	var result int
	for i := range len(array) {
		wg.Add(1)
		go func(el int) {
			//блокируем мьютекс, чтобы только одна горутина в один момент времени могла изменять переменную суммы
			m.Lock()
			result += el * el
			m.Unlock()
			wg.Done()
		}(array[i])
	}
	// ожидание завершения всех горутин. Иначе какой-то результат может быть утерян
	wg.Wait()

	fmt.Println(result)
}
