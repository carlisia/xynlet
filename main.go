package main

import (
	"errors"
	"math"
)

const defaultStartBase = 2
const defaultEndBase = 20
const defaultEndNum = 1000

// Input is the number being checked for a palindrome in the smallest base
type Input int

func main() {
	var num Input
	for num = 0; num <= defaultEndNum; num++ {
		num.smallesPalindBase(defaultEndBase)
	}
}

func (number Input) smallesPalindBase(maxNumBase int) (int, error) {
	if maxNumBase < 2 {
		return -1, errors.New("Base less than 2 not evaluated")
	}
	for base := defaultStartBase; base <= maxNumBase; base++ {
		size := number.representationSize(base)

		baseRepresentation := make([]int, size)
		number.decompose(baseRepresentation, base)

		if isPalindrome(baseRepresentation) {
			return base, nil
		}
	}

	return 0, nil
}

func (number Input) representationSize(base int) int {
	if number > 0 {
		x := logB(float64(number), base)
		size := int(math.Ceil(x))

		if int(x) == size {
			size++
		}
		return size
	}

	return 1
}

func (number Input) decompose(baseRepresentation []int, base int) {
	current := int(number)

	for i := range baseRepresentation {
		if current < base {
			baseRepresentation[i] = current
			break
		}
		baseRepresentation[i] = current % base
		current = current / base
	}
}

func isPalindrome(baseRepresentation []int) bool {
	size := len(baseRepresentation)

	if size == 1 {
		return true
	}

	midPoint := int(math.Floor(float64(size) / float64(2.0)))

	for i := 0; i < midPoint; i++ {
		if baseRepresentation[i] != baseRepresentation[size-i-1] {
			return false
		}
	}

	return true
}

func logB(x float64, base int) float64 {
	return math.Log(x) / math.Log(float64(base))
}
