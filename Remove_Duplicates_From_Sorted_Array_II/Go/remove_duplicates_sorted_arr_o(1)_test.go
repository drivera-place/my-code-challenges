package main

import (
	"testing"
)

func Test_Count_Duplicated_First_Case(t *testing.T) {
	// Arrange
	test := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	expectedArray := []int{0, 0, 1, 1, 2, 3, 3}
	expectedK := 7

	// Act
	t.Logf(`Initial array: %v`, test)
	output := countDuplicated(test)

	// Assert
	assertK(expectedK, output, expectedArray, test, t)
}

func Test_Count_Duplicated_Second_Case(t *testing.T) {
	// Arrange
	test := []int{1, 1, 1, 2, 2, 3}
	expectedArray := []int{1, 1, 2, 2, 3}
	expectedK := 5

	// Act
	t.Logf(`Initial array: %v`, test)
	output := countDuplicated(test)

	// Assert
	assertK(expectedK, output, expectedArray, test, t)
}

func Test_Count_Duplicated_Third_Case(t *testing.T) {
	// Arrange
	test := []int{1, 1, 2, 2}
	expectedArray := []int{1, 1, 2, 2}
	expectedK := 4

	// Act
	t.Logf(`Initial array: %v`, test)
	output := countDuplicated(test)

	// Assert
	assertK(expectedK, output, expectedArray, test, t)
}

func Test_Count_Duplicated_Flat_Array(t *testing.T) {
	// Arrange
	test := []int{1, 1, 1, 1}
	expectedArray := []int{1, 1}
	expectedK := 2

	// Act
	t.Logf(`Initial array: %v`, test)
	output := countDuplicated(test)

	// Assert
	assertK(expectedK, output, expectedArray, test, t)
}

func Test_Count_Duplicated_Single_Item(t *testing.T) {
	// Arrange
	test := []int{1}
	expectedArray := []int{1}
	expectedK := 1

	// Act
	t.Logf(`Initial array: %v`, test)
	output := countDuplicated(test)

	// Assert
	assertK(expectedK, output, expectedArray, test, t)
}

func Test_Count_Duplicated_Single_Empty_Array(t *testing.T) {
	// Arrange
	test := []int{}
	expectedArray := []int{}
	expectedK := 0

	// Act
	t.Logf(`Initial array: %v`, test)
	output := countDuplicated(test)

	// Assert
	assertK(expectedK, output, expectedArray, test, t)
}

func assertK(k int, o int, exArr []int, arr []int, t *testing.T) {

	if k != o {
		t.Fatalf(`K index does not match, expected: %v, got: %v`, k, o)
	}

	for i := 0; i < k; i++ {
		if arr[i] != exArr[i] {
			t.Fatalf(`Expected item array does not match with modified array, expected: expArr[%v] %v, got: arr[%v] %v`, i, exArr[i], i, arr[i])
		}
	}

	t.Logf(`K index expected: %v`, k)
	t.Logf(`K index got: %v`, o)
	t.Logf(`Expected array: %v`, exArr)
	t.Logf(`Modified array: %v`, arr)

}
