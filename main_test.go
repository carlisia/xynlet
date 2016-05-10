package main

import (
	"math"
	"testing"
)

func TestSmallestBases(t *testing.T) {
	for _, test := range testCases {
		observed, _ := test.number.smallesPalindBase(test.maxNumBase)
		if observed != test.expected {
			t.Errorf("smallesPalindBase of %d returned %d, want %d",
				test.number, observed, test.expected)
		}
	}
}

func TestSmallestBasesWithEndBaseTooLow(t *testing.T) {
	for _, test := range testCasesBaseTooLow {
		observed, _ := test.number.smallesPalindBase(test.notHighEnoughEndBase)
		if observed == test.expected {
			t.Errorf("smallesPalindBase of %d returned %d, want 0",
				test.number, observed)
		}
	}
}

func TestSmallestBasesWithEndBaseBelowLimit(t *testing.T) {
	belowLimitEndBase := 1
	for _, test := range testCases {
		observed, error := test.number.smallesPalindBase(belowLimitEndBase)
		if error == nil {
			t.Errorf("smallesPalindBase of %d returned %d, want -1 and error",
				test.number, observed)
		}
	}
}

func TestRepresentationSize(t *testing.T) {
	for _, test := range testCasesRepresentationSize {
		observed := test.number.representationSize(test.base)
		if observed != test.expected {
			t.Errorf("representationSize of %d base %d returned %d, want %d",
				test.number, test.base, observed, test.expected)
		}
	}
}

func TestDecompose(t *testing.T) {
	for _, test := range testCasesDecomposition {
		observed := true
		size := test.number.representationSize(test.base)
		baseRepresentation := make([]int, size)
		test.number.decompose(baseRepresentation, test.base)

		for i := range baseRepresentation {
			if baseRepresentation[i] != test.expected[i] {
				observed = false
				break
			}
		}
		if observed != true {
			t.Errorf("decomposition was %v, want %v",
				baseRepresentation, test.expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range testCasesIsPalindrome {
		size := test.number.representationSize(test.base)
		baseRepresentation := make([]int, size)
		test.number.decompose(baseRepresentation, test.base)

		observed := isPalindrome(baseRepresentation)
		if observed != test.expected {
			t.Errorf("isPalindrome of %d base %d returned %t, want %t",
				test.number, test.base, observed, test.expected)
		}
	}
}

func TestLogB(t *testing.T) {
	const TOLERANCE = 0.000001

	for _, test := range testCasesLogB {
		observed := logB(float64(test.number), test.base)
		// Note: two float64 operands are considered equal when the absolute
		// value of their deltas are smaller than the tolerance.
		if diff := math.Abs(observed - test.expected); diff >= TOLERANCE {
			t.Errorf("logB of %d and %d returned %f, want %f",
				test.number, test.base, observed, test.expected)
		}
	}
}
