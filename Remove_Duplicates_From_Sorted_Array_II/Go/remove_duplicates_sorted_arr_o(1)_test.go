package main

import (
	"testing"
)

func TestRemoveDuplicatesSortedArrayLineal(t *testing.T) {
	// Arrange
	test := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	//output := []int{0,0,1,1,2,3,3,0,0}
	k := 7

	// Act
	ouput := countDuplicated(test)

	// Assert
	assertK(k, ouput, t)
}

func countDuplicated(arr []int) int {
	len := len(arr)
	count := 0

	for i := 0; i < len; i++ {
		if i < len-1 && arr[i] == arr[i+1] {
			continue
		}
		arr[count] = arr[i]
		count++
	}

	return count
}

func assertK(e int, o int, t *testing.T) {

	if e != o {
		t.Fatalf(`Count does not match.`)
	}

}
