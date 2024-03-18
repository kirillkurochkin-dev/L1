package main

import (
	"fmt"
)

func main() {
	arr := []int{-1, -1, 10, 5, -99, -27, 11}
	fmt.Println("Unsorted array:", arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}

// сортирует массив arr в порядке возрастания
func quicksort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quicksort(arr, low, pivot-1)
		quicksort(arr, pivot+1, high)
	}
}

// выполняет разделение массива arr на две части относительно опорного элемента и возвращает индекс опорного элемента после разделения
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
