package main

import "testing"

func Test_Product_Exclusion(t *testing.T) {
	// Arrange
	input := []int{2, 3, 4, 5}
	expected := []int{60, 40, 30, 24}

	// Act
	output := productExclusion(input)

	// Assert
	assert(expected, output, t)
}

func Test_Product_Exclusion_Empty_Array(t *testing.T) {
	// Arrange
	input := []int{}
	expected := []int{}

	// Act
	output := productExclusion(input)

	// Assert
	assert(expected, output, t)
}

func Test_Product_Exclusion_Single_Item_Array(t *testing.T) {
	// Arrange
	input := []int{0}
	expected := []int{1}

	// Act
	output := productExclusion(input)

	// Assert
	assert(expected, output, t)
}

func Test_Product_Exclusion_Array_With_A_Zero(t *testing.T) {
	// Arrange
	input := []int{0, 2, 3, 5}
	expected := []int{30, 0, 0, 0}

	// Act
	output := productExclusion(input)

	// Assert
	assert(expected, output, t)
}

func assert(expected []int, output []int, t *testing.T) {

	t.Logf(`Expected array: %v`, expected)
	t.Logf(`Output array:   %v`, output)

	for i := 0; i < len(expected); i++ {
		if expected[i] != output[i] {
			t.Fatalf(`Item does not match, expected[%v]:%v, output[%v]:%v`, i, expected[i], i, output[i])
		}
	}
}
