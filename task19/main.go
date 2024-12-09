// Разработать программу, которая переворачивает подаваемую на ход строку
//  (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку для переворота:")

	input, _ := reader.ReadString('\n')
	// Убираем символ новой строки.
	input = strings.TrimSpace(input)

	reversed := reverseString(input)

	fmt.Println("Перевёрнутая строка:")
	fmt.Println(reversed)
}
