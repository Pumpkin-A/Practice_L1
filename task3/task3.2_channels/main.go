package main

import (
	"fmt"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	//инициализация канала
	ch := make(chan int)

	for _, el := range array {
		go func() {
			ch <- el * el
		}()
	}
	var sum int
	for range len(array) {
		sum += <-ch
	}
	//ожидаем завершение работы горутин с помощью канала
	close(ch)
	fmt.Println(sum)
}
