// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора
// количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	var workersNumber int
	fmt.Print("Print number of workers: ")
	_, err := fmt.Scanf("%d", &workersNumber)
	if err != nil {
		fmt.Println("int expected")
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for i := range workersNumber {
		wg.Add(1)
		go func(ctx context.Context, i int) {
			defer wg.Done()
			for {
				value, ok := <-ch
				if !ok {
					return
				}
				fmt.Println(i, ":", value)
			}
		}(ctx, i)
	}

	var i int
	for {
		if ctx.Err() == context.Canceled {
			close(ch)
			wg.Wait()
			fmt.Println("все хорошо!!!!!!!!!!!!!!!!!!!!!!!")
			return
		}
		fmt.Println("отправляем:", i)
		ch <- i
		i++
	}

}
