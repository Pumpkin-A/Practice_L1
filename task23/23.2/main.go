package main

import "fmt"

// Этот вариант удаления элемента слайса валиден, но в таком случае буфер изначального слайса не удаляется, то есть
// новый слайс корректен, элемента в нем не будет, однако где-то в памяти он все еще будет храниться, сборщик
// мусора не сможет его удалить, так как на буфер указывает срез.

func deleteElem(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("the position of the element out of range slice")
		return slice
	}

	var newSlice []int
	newSlice = append(newSlice, slice[:i]...)
	newSlice = append(newSlice, slice[i+1:]...)
	return newSlice
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice = deleteElem(slice, 2)
	fmt.Println(slice)
}
