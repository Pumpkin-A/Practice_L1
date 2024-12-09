package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	ch := make(chan bool)

	// Запускаем горутину, которая будет отправлять сообщение в канал через заданное время
	go func() {
		time.AfterFunc(duration, func() {
			ch <- true
		})
	}()

	<-ch
}

func main() {
	fmt.Println("sleep for 5 seconds")
	mySleep(5 * time.Second)
	fmt.Println("sleep for 5 seconds is over")
}
