// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a и b,
// значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// в данной задаче использованы типы bigInt и bigFloat, так как даже при исходных данных
	// 2^50 и 2^62, как указано в примере ниже, числа хоть и поместятся в int64, так как
	//его вместимость до 2^63-1, но результаты не всех операций поместятся в этот тип.
	// так, например, результат операции произведения при данных числах привел бы к ошибке.
	//
	// вариант с использованием нетипизированной константы так же не подходит, так как
	// При передаче к функциям нетипизированные константы должны быть конвертированы в типизированные переменные,
	// то есть при выводе того же произведения (передачи в функцию ptintln) возникала бы ошибка

	a := big.NewInt(1)
	b := big.NewInt(1)

	a.Lsh(a, 50)
	b.Lsh(b, 62)

	a.SetString(a.String(), 10)
	b.SetString(b.String(), 10)

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	product := new(big.Int).Mul(a, b)
	quotient := new(big.Float).Quo(new(big.Float).SetInt(a), new(big.Float).SetInt(b))

	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Printf("Сумма: %s\n", sum.String())
	fmt.Printf("Разность: %s\n", diff.String())
	fmt.Printf("Произведение: %s\n", product.String())
	fmt.Printf("Частное: %s\n", quotient.String())
}
