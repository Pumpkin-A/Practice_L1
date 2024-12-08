// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func binarySearchAscend(arr []int, searchElem int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := int((left + right) / 2)

		if arr[mid] == searchElem {
			return mid
		} else if arr[mid] < searchElem {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}

func binarySearchDescend(arr []int, searchElem int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := int((left + right) / 2)

		if arr[mid] == searchElem {
			return mid
		} else if arr[mid] < searchElem {
			right = mid - 1

		} else {
			left = mid + 1
		}

	}
	return -1
}

func main() {
	arr1 := make([]int, 0)

	//массив обязательно должен быть отсортирован
	for i := range 100 {
		arr1 = append(arr1, i)
	}

	searchElemPos := binarySearchAscend(arr1, 68)
	if searchElemPos == -1 {
		fmt.Println("there is no such element in the array")
	} else {
		fmt.Println("position of element:", searchElemPos)
	}

	arr2 := []int{87, 65, 44, 32, 13, 5, 1}
	searchElemPos = binarySearchDescend(arr2, 13)
	if searchElemPos == -1 {
		fmt.Println("there is no such element in the array")
	} else {
		fmt.Println("position of element:", searchElemPos)
	}

}
