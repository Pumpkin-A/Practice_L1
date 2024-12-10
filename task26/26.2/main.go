package main

import (
	"fmt"
	"sort"
	"strings"
)

// Если порядок символов не важен, в решении задачи можно использовать сортировку
// Если сортировать строку, то повторяющиеся символы окажутся рядом, и их можно легко обнаружить, сравнивая соседние символы.
// из минусов: требует сортировки (O(n log n)
func areAllUnique(str string) bool {
	str = strings.ToLower(str)

	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Проверяем соседние символы
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			return false
		}
	}

	return true
}

func main() {
	stringsToCheck := []string{"abcd", "abCdefAf", "aabcd", "Ррпва", "Правоыф"}

	for _, str := range stringsToCheck {
		result := areAllUnique(str)
		fmt.Println(str, result)
	}
}
