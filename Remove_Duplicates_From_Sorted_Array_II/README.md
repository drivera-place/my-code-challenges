# Remove duplicates from sorted Array II

Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears **at most twice**.

The relative order of the elements should be kept the same.

You must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. *It does not matter what you leave beyond the first k elements.*

**Return k after placing the final result in the first k slots of nums.**

**Do not allocate extra space for another array.**

**You must do this by modifying the input array in-place with O(1) extra memory.**

Example of scenarios:

```Go
// Arrange
test := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
expectedArray := []int{0, 0, 1, 1, 2, 3, 3}
expectedK := 7

// Act
output := countDuplicated(test)

// Assert
assertK(expectedK, output, expectedArray, test, t)
```

## Test

Golang

```bash
cd Go/
go test -v
```
