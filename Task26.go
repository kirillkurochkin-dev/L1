package main

import (
	"fmt"
	"strings"
)

// isUnique проверяет, содержит ли строка только уникальные символы.
func isUnique(str string) bool {
	str = strings.ToLower(str)
	charMap := make(map[rune]bool)

	for _, char := range str {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}

	return true
}

func main() {
	testStrings := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, str := range testStrings {
		fmt.Printf("%s — %v\n", str, isUnique(str))
	}
}
