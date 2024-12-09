// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

//вариант решения задачи без использования каналов и горутин

func sleep(timeFinish time.Time) {
	fmt.Println("sleep for 5 seconds")
	for {
		if time.Now().After(timeFinish) {
			return
		}
	}
}

func main() {
	timeFirst := time.Now()
	timeSleep := time.Duration(5 * time.Second)

	sleep(timeFirst.Add(timeSleep))
	fmt.Println("sleep for 5 seconds is over")
}
