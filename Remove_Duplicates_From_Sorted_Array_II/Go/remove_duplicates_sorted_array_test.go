package main

import (
	"testing"
)

func Test_Remove_Duplicates_Sorted_Array(t *testing.T) {
	// Arrange
	test := []int{1, 2, 3, 4, 4, 5, 6, 7, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Act
	output := removeDuplicated(test)

	// Assert
	assert(expected, output, t)
}

func assert(e []int, o []int, t *testing.T) {

	t.Logf("Expected: %v \n", e)
	t.Logf("Output:   %v \n", o)

	var expLen, outLen = len(e), len(o)

	if expLen != outLen {
		t.Fatalf(`Expected array length does not match, expected: %v, output: %v `, expLen, outLen)
	}

	for i := 0; i < expLen; i++ {
		if e[i] != o[i] {
			t.Fatalf(`Expected array element: %v does not match with output: %v`, e[i], o[i])
		}
	}
}
