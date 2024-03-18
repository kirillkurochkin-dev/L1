package main

import "fmt"

func main() {
	// Для нас важны только ключи, значения не имеют значения.
	set := make(map[string]struct{})

	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, s := range sequence {
		// Добавляем элемент s в множество set
		// Поскольку map использует ключи для хранения уникальных значений, дубликаты будут автоматически игнорироваться
		// Кстати, в Java примерно так и работает Set там она идет из коробки
		set[s] = struct{}{}
	}

	fmt.Println("Уникальные элементы:")
	for key := range set {
		fmt.Println(key)
	}
}
