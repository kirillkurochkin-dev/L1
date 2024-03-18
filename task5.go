package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	var N int

	// Запрашиваем у пользователя длительность времени в секундах
	fmt.Print("Enter time duration: ")
	_, err := fmt.Scanln(&N)
	if err != nil {
		return
	}

	// Создаем таймер с указанной длительностью
	timer := time.NewTimer(time.Duration(N) * time.Second)

	// Сигнальный канал
	done := make(chan struct{})

	// Запускаем горутину, которая отправляет значения в канал каждую секунду
	go func(ch chan<- int) {
		defer close(ch)
		for i := 0; ; i++ {
			select {
			// Если получен сигнал о завершении работы, выходим из горутины
			case <-done:
				return
			// Отправляем значение в канал и выводим его на экран
			case ch <- i:
				fmt.Println("Send value", i)
				time.Sleep(1 * time.Second)
			}
		}
	}(ch)

	for {
		select {
		// Если пришло значение из канала, выводим его на экран
		case data := <-ch:
			fmt.Println("Get value", data)
		// Если произошел таймаут, выводим сообщение и завершаем работу горутины
		case <-timer.C:
			fmt.Println("Timeout")
			close(done)
			return
		}
	}
}
