package main

import "testing"

func TestSmallestBases(t *testing.T) {
	for _, test := range testCases {
		observed, _ := test.number.smallesPalindBase(defaultEndBase)
		if observed != test.expected {
			t.Errorf("smallesPalindBase of %d returned %d, want %d (%s)",
				test.number, observed, test.expected, test.base)
		}
	}
}

func TestSmallestBasesWithHigherBase(t *testing.T) {
	notHighEnoughEndBase := 4
	for _, test := range testCasesHigherBases {
		observed, _ := test.number.smallesPalindBase(notHighEnoughEndBase)
		if observed == test.expected {
			t.Errorf("smallesPalindBase of %d returned %d, want 0 (%s)",
				test.number, observed, test.base)
		}
	}
}

func TestSmallestBasesWithEndBaseTooSmall(t *testing.T) {
	smallEndBase := 1
	for _, test := range testCases {
		observed, error := test.number.smallesPalindBase(smallEndBase)
		if error == nil {
			t.Errorf("smallesPalindBase of %d returned %d, want -1 and error (%s)",
				test.number, observed, test.base)
		}
	}
}
