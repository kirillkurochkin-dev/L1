package main

import (
	"fmt"
	"sync"
)

func main() {
	// Задаем исходный массив
	arr := []int{2, 4, 6, 8, 10}

	// Создаем канал !с буфером!
	ch := make(chan int, len(arr))

	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		// Это нам нужно для синхронизации
		wg.Add(1)
		go func(num int) {
			// Это нам тоже обязательно нужно для синхронизации
			defer wg.Done()
			// Вычисляем квадрат числа и отправляем его в канал
			ch <- num * num
		}(arr[i]) // Замыкание
	}

	// ТОже элемент синхронизации, просто ждем горутины
	wg.Wait()
	// Закрываем канал
	close(ch)

	for square := range ch {
		fmt.Println(square)
	}
}
