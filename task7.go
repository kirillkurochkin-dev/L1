package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Безопасня map
type SafeMap struct {
	m   map[string]int
	mux sync.Mutex
}

// Add добавляет значение в map по ключу
func (sm *SafeMap) Add(key string, value int) {
	sm.mux.Lock()
	defer sm.mux.Unlock()
	sm.m[key] = value
}

// Get возвращает значение из map по ключу
func (sm *SafeMap) Get(key string) int {
	sm.mux.Lock()
	defer sm.mux.Unlock()
	return sm.m[key]
}

func main() {
	wg := sync.WaitGroup{}
	// Создаем экземпляр SafeMap для хранения данных
	sm := SafeMap{m: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			sm.Add(strconv.Itoa(i), 1)
		}()
	}
	wg.Wait()
	fmt.Println(sm.m)
}
