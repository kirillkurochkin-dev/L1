package main

import (
	"fmt"
	"strings"
)

func ReverseWords(sentence string) string {
	// Разбиваем строку на слова.
	words := strings.Fields(sentence)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		// Меняем местами слова симметрично относительно середины списка.
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun — sun dog snow"
	reversed := ReverseWords(input)
	fmt.Println("Перевернутая строка:", reversed)
}
