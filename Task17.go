package main

import "fmt"

// binarySearch принимает отсортированный массив arr и значение, дальше алгоритм
func binarySearch(arr []int, target int) int {
	low := 0             // Нижняя граница поиска
	high := len(arr) - 1 // Верхняя граница поиска

	for low <= high {
		mid := (low + high) / 2 // Вычисляем середину массива
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 11
	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден в позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
