package main

import (
	"fmt"
)

// Функция intersection возвращает пересечение двух множеств set1 и set2
func intersection(set1, set2 []int) []int {
	setMap := make(map[int]struct{})
	result := []int{}

	// Заполняем карту уникальными значениями из множества set1
	for _, v := range set1 {
		setMap[v] = struct{}{}
	}

	// Проверяем каждый элемент множества set2 на наличие в карте
	// Если элемент найден, добавляем его в результат
	for _, v := range set2 {
		if _, ok := setMap[v]; ok {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}
	fmt.Println(intersection(set1, set2))
}
