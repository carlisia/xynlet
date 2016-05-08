package main

import "strconv"

// IsPalindrome checks if a given  number is a palindrome in any base
func IsPalindrome(i int) bool {
	s := strconv.Itoa(i)
	return s == reverse(s)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) // Convert back to UTF-8
}
