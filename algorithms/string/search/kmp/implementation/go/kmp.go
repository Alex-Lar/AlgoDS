package main

import (
	"fmt"
)

// Сложность составления таблицы префиксов - O(m)
func computePrefixTable(pattern string) []int {
	prefixTable := make([]int, len(pattern))
	j := 0
	i := 1

	for i < len(pattern) {
		if pattern[j] == pattern[i] {
			prefixTable[i] = j + 1
			i += 1
			j += 1
		} else {
			if j == 0 {
				prefixTable[i] = 0
				i += 1
			} else {
				j = prefixTable[j-1]
			}
		}
	}

	return prefixTable
}

// Сложность KMP-алгоритма - O(n+m)
func KMP(pattern string, text string) {
	prefixTable := computePrefixTable(pattern)

	patternLength := len(pattern)
	text_length := len(text)

	i := 0
	j := 0
	for i < text_length {
		if text[i] == pattern[j] {
			i += 1
			j += 1
			if j == patternLength {
				fmt.Println("Pattern is found")
				break
			}
		} else {
			if j > 0 {
				j = prefixTable[j-1]
			} else {
				i += 1
			}
		}
	}

	if i == text_length {
		fmt.Println("Pattern is not found")
	}
}

func main() {
	pattern := "ababac"
	text := "ababaqde ababacde"

	KMP(pattern, text)
}
