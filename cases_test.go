package main

var testCases = []struct {
	number     Input // input as a decimal number
	maxNumBase int   // any base greater than the lower limit (2) to be used as an upper bound
	expected   int   // smallest base in which the input is a palindrome
}{
	{0, 3, 2},
	{1, 3, 2},
	{2, 4, 3},
	{3, 3, 2},
	{4, 4, 3},
	{5, 3, 2},
	{6, 6, 5},
	{7, 3, 2},
	{8, 4, 3},
	{9, 3, 2},
	{10, 3, 3},
	{11, 10, 10},
	{12, 5, 5},
	{13, 3, 3},
	{14, 6, 6},
	{15, 2, 2},
	{16, 3, 3},
	{17, 2, 2},
	{18, 5, 5},
	{19, 18, 18},
}

var testCasesBaseTooLow = []struct {
	number               Input // input as a decimal number
	notHighEnoughEndBase int   // any base greater than the lower limit (2) but not high enough to find a match for the given input
	expected             int   // smallest base in which the input is a palindrome
}{
	{6, 4, 5},
	{11, 9, 10},
	{12, 4, 5},
	{14, 5, 6},
	{18, 3, 5},
	{19, 10, 18},
}

var testCasesRepresentationSize = []struct {
	number   Input // given input as a decimal number
	base     int   // given base to calculate from
	expected int   // number of digits needed to represent the input number in positional notation for the given base
}{
	{0, 2, 1},
	{1, 2, 1},
	{2, 3, 1},
	{3, 2, 2},
	{4, 3, 2},
	{5, 2, 3},
	{6, 5, 2},
	{7, 2, 3},
	{8, 3, 2},
	{9, 2, 4},
	{10, 3, 3},
	{11, 10, 2},
	{12, 5, 2},
	{13, 3, 3},
	{14, 6, 2},
	{15, 2, 4},
	{16, 3, 3},
	{17, 2, 5},
	{18, 5, 2},
	{19, 18, 2},
}

var testCasesDecomposition = []struct {
	number   Input // given input as a decimal number
	base     int   // given base to calculate from
	expected []int //
}{
	{0, 2, []int{0}},
	{1, 2, []int{1}},
	{2, 3, []int{2}},
	{3, 2, []int{1, 1}},
	{4, 3, []int{1, 1}},
	{5, 2, []int{1, 0, 1}},
	{6, 5, []int{1, 1}},
	{7, 2, []int{1, 1, 1}},
	{8, 3, []int{2, 2}},
	{9, 2, []int{1, 0, 0, 1}},
	{10, 3, []int{1, 0, 1}},
	{11, 10, []int{1, 1}},
	{12, 5, []int{2, 2}},
	{13, 3, []int{1, 1, 1}},
	{14, 6, []int{2, 2}},
	{15, 2, []int{1, 1, 1, 1}},
	{16, 3, []int{1, 2, 1}},
	{17, 2, []int{1, 0, 0, 0, 1}},
	{18, 5, []int{3, 3}},
	{19, 18, []int{1, 1}},
}

var testCasesIsPalindrome = []struct {
	number   Input // input number
	base     int   // base for which to check if input is palindrome
	expected bool
}{
	{2, 3, true},
	{3, 2, true},
	{4, 3, true},
	{5, 2, true},
	{5, 3, false},
	{14, 2, false},
}

var testCasesLogB = []struct {
	number   Input
	base     int
	expected float64
}{
	{2, 3, 0.6309303},
	{3, 2, 1.584963},
	{4, 3, 1.261860},
	{5, 2, 2.321928},
}
