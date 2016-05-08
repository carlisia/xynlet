package main

var testCases = []struct {
	decimal      int  // input as a decimal number
	smallestBase int  // smallest base in which the input is a palindrome
	expected     bool // is palindrome?
	base         string
}{
	{0, 2, true, "binary"},
	{1, 2, true, "binary"},
	{2, 3, true, "three"},
	{3, 2, true, "binary"},
	{4, 3, true, "three"},
	{5, 2, true, "binary"},
	{6, 5, true, "five"},
	{7, 2, true, "binary"},
	{8, 3, true, "three"},
	{9, 2, true, "binary"},
	{10, 3, true, "three"},
	{11, 10, true, "decimal"},
	{12, 5, true, "five"},
	{13, 3, true, "three"},
	{14, 6, true, "six"},
	{15, 2, true, "binary"},
	{16, 3, true, "three"},
	{17, 2, true, "binary"},
	{18, 5, true, "five"},
	{19, 18, true, "eighteen"},
}
