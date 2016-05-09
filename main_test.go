package main

import "testing"

func TestSmallestBases(t *testing.T) {
	for _, test := range testCases {
		observed := test.number.smallesPalindBase(defaultEndBase)
		if observed != test.expected {
			t.Errorf("smallesPalindBase of %d returned %d, want %d (%s)",
				test.number, observed, test.expected, test.base)
		}
	}
}
