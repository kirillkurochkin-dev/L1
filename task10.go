package main

import "fmt"

func main() {
	// Создаем пустой словарь
	dict := make(map[int][]float64)
	list := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, f := range list {
		// Вычисляем ключ для словаря, округляя число f до ближайшего десятка
		// Добавляем число f в массив, соответствующий ключу key в словаре
		dict[int(f/10)*10] = append(dict[int(f/10)*10], f)
	}
	fmt.Println(dict)
}
