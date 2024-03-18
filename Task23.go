package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	index := 2

	// Прямого удаления в Go нет, поэтому мы буквально выдавливаем указанное значение
	if index < len(slice) {
		slice = append(slice[:index], slice[index+1:]...)
	} else {
		fmt.Println("Индекс вне диапазона среза")
	}

	fmt.Println(slice)
}
