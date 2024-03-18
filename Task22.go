package main

import (
	"fmt"
	"math"
)

func main() {
	a := math.Pow(2, 20) + 1
	b := math.Pow(2, 20) + math.Pow(2, 20)

	// Это просто убедиться выполняю ли я правильно условие задачи
	if a > math.Pow(2, 20) && b > math.Pow(2, 20) {
		multiplication := a * b
		fmt.Printf("Умножение: %.0f\n", multiplication)

		division := a / b
		fmt.Printf("Деление: %.5f\n", division)

		addition := a + b
		fmt.Printf("Сложение: %.0f\n", addition)

		subtraction := a - b
		fmt.Printf("Вычитание: %.0f\n", subtraction)
	} else {
		fmt.Println("Значения переменных a и b должны быть больше двух в двадцатой степени")
	}
}
