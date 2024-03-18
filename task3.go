package main

import (
	"fmt"
	"sync"
)

func main() {
	// Мьютекс - длясинхронизации каких-либо данных
	mu := sync.Mutex{}
	// Для синхронизации, чтобы прога не завершилась раньше чем нужно
	wg := sync.WaitGroup{}

	var sum int

	arr := []int{2, 4, 6, 8, 10}

	for _, i := range arr {

		wg.Add(1)
		go func(x int) {
			// Блокируем доступ
			mu.Lock()

			sum += x * x
			// Разблокируем доступ
			mu.Unlock()

			defer wg.Done()
		}(i)
	}
	// Ждем, пока все горутины завершатся
	wg.Wait()

	fmt.Println(sum)
}
