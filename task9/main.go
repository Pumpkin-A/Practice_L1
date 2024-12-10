// Разработать конвейер чисел. Даны два канала:
//  в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//  после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	buf := [20]int{}
	for i := range 20 {
		buf[i] = i
	}

	go func() {
		for _, v := range buf {
			ch1 <- v
		}
		close(ch1)
	}()

	go func() {
		for {
			v, ok := <-ch1
			if !ok {
				close(ch2)
				return
			}
			ch2 <- v * 2
		}
	}()

	for v := range ch2 {
		fmt.Println(v)
	}
	// wait group можно не использовать, так как при данной реализации к моменту закрытия канала 2
	// все горутины завершат свое выполнение
	fmt.Println("all goroutines are done")
}
