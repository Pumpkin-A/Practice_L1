// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import "fmt"

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]

	for k := low; k <= i-1; k++ {
		fmt.Print(arr[k], " ")
	}
	fmt.Println()
	fmt.Println("pivot:", pivot, "pos", i)
	for k := i + 1; k <= high; k++ {
		fmt.Print(arr[k], " ")
	}
	fmt.Println()
	fmt.Println()

	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 1, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array: ", arr)
}
