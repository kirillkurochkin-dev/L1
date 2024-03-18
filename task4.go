package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// читает значения из канала ch и выводит их в stdout
func read(ch <-chan string, id int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for {
		mu.Lock()
		fmt.Printf("Worker with id %d get value from chanel: %s\n", id, <-ch)
		mu.Unlock()
		time.Sleep(time.Second)
	}
}

func main() {
	var count int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	// Создаем канал для сигналов
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Print("Enter workers count: ")
	_, err := fmt.Scanln(&count)
	if err != nil {
		return
	}

	// Создаем канал для передачи данных между горутинами
	ch := make(chan string, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go read(ch, i, &wg, &mu)
	}

	// Ожидаем сигнала завершения работы
	for i := 0; ; i++ {
		select {
		case <-sigs:
			fmt.Println("Done..")
			close(sigs)
			close(ch)
			return
		default:
			ch <- fmt.Sprintf("%d", i)
		}

	}

	wg.Wait()
}
