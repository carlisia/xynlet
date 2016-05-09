package main

var testCases = []struct {
	number   Input // input as a decimal number
	expected int   // smallest base in which the input is a palindrome
	base     string
}{
	{0, 2, "binary"},
	{1, 2, "binary"},
	{2, 3, "three"},
	{3, 2, "binary"},
	{4, 3, "three"},
	{5, 2, "binary"},
	{6, 5, "five"},
	{7, 2, "binary"},
	{8, 3, "three"},
	{9, 2, "binary"},
	{10, 3, "three"},
	{11, 10, "decimal"},
	{12, 5, "five"},
	{13, 3, "three"},
	{14, 6, "six"},
	{15, 2, "binary"},
	{16, 3, "three"},
	{17, 2, "binary"},
	{18, 5, "five"},
	{19, 18, "eighteen"},
}

var testCasesHigherBases = []struct {
	number   Input // input as a decimal number
	expected int   // smallest base in which the input is a palindrome
	base     string
}{
	{6, 5, "five"},
	{11, 10, "decimal"},
	{12, 5, "five"},
	{14, 6, "six"},
	{18, 5, "five"},
	{19, 18, "eighteen"},
}
