// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//  Приведите корректный пример реализации.

package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
	if size <= 0 {
		return ""
	}
	return strings.Repeat("A", size)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	// justString = v[:100]

	//При использовании старого варианта (создания среза) сборщик мусора не может очистить большую строку,
	// на нее указывает созданный срез. В исправленном варианте возвращает новую копию среза.
	// Это гарантирует, что копия s будет добавлена в новую область памяти, избавляясь от большой строки.
	justString = strings.Clone(v[:100])
	return justString
}

func main() {
	fmt.Println(someFunc())
}
