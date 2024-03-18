package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Способ 1: Возврат из функции
	go func() {
		fmt.Println("Горутина завершится после выполнения функции")
	}()

	// Способ 2: Использование сигнального канала
	ch := make(chan struct{})
	go func() {
		fmt.Println("Горутина ожидает сигнала для завершения")
		<-ch
	}()
	// Вызовем завершение горутины через канал
	ch <- struct{}{}

	// Способ 3: Использование контекста
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("Горутина ожидает отмены контекста")
		<-ctx.Done()
	}()
	// Вызовем отмену контекста для завершения горутины
	cancel()

	// Способ 4: runtime.Goexit()
	go func() {
		fmt.Println("Горутина завершится при использовании runtime.Goexit()")
		runtime.Goexit()
	}()

	// Способ 5: Примитивы синхронизации
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина завершится с помощью sync.WaitGroup")
	}()
	wg.Wait()

	// Способ 6: Системный вызов
	go func() {
		fmt.Println("Горутина завершится после спящего режима")
		time.Sleep(time.Second)
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Программа завершена")
}
