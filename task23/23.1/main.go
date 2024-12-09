// Удалить i-ый элемент из слайса.

package main

import "fmt"

// Этот вариант удаления элемента слайса наиболее оптимален, так как при его использовании
// создается новый буфер и не происходит обращения к буферу старого слайса, как это было бы при использовании срезов.

func deleteElem(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("the position of the element out of range slice")
		return slice
	}
	var newSlice []int
	for j := range slice {
		if i != j {
			newSlice = append(newSlice, slice[j])
		}
	}
	return newSlice
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := deleteElem(slice, 3)
	fmt.Println(newSlice)
}
