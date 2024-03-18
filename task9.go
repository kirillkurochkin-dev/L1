package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch1 := make(chan int)
	ch2 := make(chan int)

	// Запуск первой горутины для обработки данных из ch1 и отправки результата в ch2
	go func() {
		wg.Add(1)
		defer wg.Done()
		for num := range ch1 {
			// Отправка результата в ch2 (квадрат числа)
			ch2 <- num * num
		}
		close(ch2)
	}()

	// Запуск второй горутины для чтения данных из ch2 и вывода их на экран
	go func() {
		wg.Add(1)
		defer wg.Done()
		for num := range ch2 {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}()

	data := []int{1, 2, 3, 4, 5, 6}

	for _, value := range data {
		ch1 <- value
	}
	close(ch1)
	wg.Wait()
}
