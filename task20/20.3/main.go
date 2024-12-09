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
	var sb strings.Builder
	words := strings.Split(str, " ")

	// если не просто менять порядок слов в слайсе строк, то можно использовать конкатенацию строк.
	// Использование builder для конкатенции строк наиболее эффективно, особенно при большом объеме строк,
	// так как при его использовании не выделяется память под каждую новую строку, как это происходит
	// при использовании "+"
	for i := len(words) - 1; i >= 0; i-- {
		sb.WriteString(words[i])
		if i > 0 {
			sb.WriteString(" ")
		}
	}

	return sb.String()
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
