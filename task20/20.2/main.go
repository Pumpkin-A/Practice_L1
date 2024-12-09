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
	var reversedString string
	words := strings.Split(str, " ")

	// если не просто менять порядок слов в слайсе строк, то можно использовать конкатенацию строк.
	// Использование "+" для этой задачи корректно, но не эффективно при большом объемы строк
	for i := len(words) - 1; i >= 0; i-- {
		reversedString += words[i]
		if i > 0 {
			reversedString += " "
		}
	}

	return reversedString
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
