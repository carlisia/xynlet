package main

import "testing"

func TestPalindromes(t *testing.T) {
	for _, test := range testCases {
		if test.base != "decimal" {
			continue // Not yet implemented
		}
		observed := IsPalindrome(test.decimal)
		if observed != test.expected {
			t.Errorf("IsPalindrome of %d = %t, want %t (%s)",
				test.decimal, observed, test.expected, test.base)
		}
	}
}
