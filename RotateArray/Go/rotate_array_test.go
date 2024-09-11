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
	t.Logf(`Input array: %v`, test)
	output := rotateArray(test, k)

	// Assert
	assert(expected, output, t)
}

func TestCasetwo(t *testing.T) {
	// Arrange
	input := []int{-1, -100, 3, 99}
	k := 2
	expected := []int{3, 99, -1, -100}

	// Act
	t.Logf(`Input array: %v`, input)
	output := rotateArray(input, k)

	// Assert
	assert(expected, output, t)
}

func assert(e, o []int, t *testing.T) {

	t.Logf(`Expected array: %v`, e)
	t.Logf(`Output array: %v`, o)

	if len(e) != len(o) {
		t.Fatalf(`Arrays length is not equal: expected: %v, output %v:`, e, o)
	}

	for i := 0; i < len(e); i++ {
		if e[i] != o[i] {
			t.Fatalf(`Expected array is not equal to received rotated array: expected: %v, output: %v`, e, o)
		}
	}
}
