// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func intersection(set1, set2 map[int]struct{}) map[int]struct{} {
	setIntersection := make(map[int]struct{})
	for k := range set1 {
		if _, ok := set2[k]; ok {
			setIntersection[k] = struct{}{}
		}
	}
	return setIntersection
}

func main() {
	var set1 = make(map[int]struct{})
	var set2 = make(map[int]struct{})
	set1Keys := []int{1, 4, 5, 6}
	set2Keys := []int{4, 6, 7, 8, 45}
	for _, v := range set1Keys {
		set1[v] = struct{}{}
	}
	for _, v := range set2Keys {
		set2[v] = struct{}{}
	}

	setIntersection := intersection(set1, set2)
	for k := range setIntersection {
		fmt.Println(k)
	}
}
