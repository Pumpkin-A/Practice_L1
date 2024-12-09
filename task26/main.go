// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func areAllUnique(str string) bool {
	// Приводим строку к нижнему регистру для игнорирования регистра
	str = strings.ToLower(str)
	runes := []rune(str)

	seen := make(map[rune]bool)
	for _, char := range runes {
		if seen[char] {
			return false
		}
		seen[char] = true
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
