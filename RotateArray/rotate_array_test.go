package main

import (
	"testing"
)

func TestRotateArray(t *testing.T) {
	// Arrange
	test := []int{1, 2, 3, 4, 5}
	k := 2
	expected := []int{4, 5, 1, 2, 3}

	// Act
	output := rotateArray(test, k)

	// Assert
	assert(expected, output, t)
}

func assert(expected, output []int, t *testing.T) {

	if len(expected) != len(output) {
		t.Fatalf(`Arrays length is not equal: expected: %v, output %v:`,expected, output)
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != output[i] {
			t.Fatalf(`Expected array is not equal to received rotated array: expected: %v, output: %v`, expected, output)
		}
	}
}
