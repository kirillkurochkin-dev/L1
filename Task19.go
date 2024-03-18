package main

import "fmt"

// Функция reverse принимает строку line и возвращает её перевёрнутую версию
func reverse(line string) string {
	data := []rune(line)
	// Итерируемся по половине массива.
	for i := 0; i < len(data)/2; i++ {
		// Меняем местами символы симметрично относительно середины массива
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}
	return string(data)
}

func main() {
	var line string
	fmt.Scan(&line)
	fmt.Println(reverse(line))
}
