package main

import (
	"fmt"
	"math/rand"
)

// setBit устанавливает бит в числе на указанной позиции
func setBit(num int64, pos uint, bitValue bool) int64 {
	if bitValue {
		return num | (1 << pos) // Установка бита в 1
	} else {
		return num &^ (1 << pos) // Установка бита в 0
	}
}

func main() {
	// Генерация случайного числа типа int64
	num := rand.Int63()

	// Генерация случайной позиции бита в диапазоне от 0 до 63
	pos := uint(rand.Intn(64))

	// Генерация случайного значения бита (0 или 1)
	bitValue := rand.Intn(2) == 1

	fmt.Printf("Сгенерированное число: %d\n", num)
	fmt.Printf("Сгенерированная позиция бита: %d\n", pos)
	fmt.Printf("Сгенерированное значение бита: %v\n", bitValue)

	// Установка бита в числе и вывод результата
	result := setBit(num, pos, bitValue)
	fmt.Printf("Число после установки бита: %d\n", result)
}
