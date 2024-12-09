// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWordsInString(str string) string {
	words := strings.Split(str, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку для переворота:")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	reversed := reverseWordsInString(input)

	fmt.Println("Перевёрнутая строка:")
	fmt.Println(reversed)
}
