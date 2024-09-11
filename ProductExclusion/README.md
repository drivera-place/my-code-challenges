# Product Exclusion

Given an array of integers, return an array of integers such that the value of index i of the output is the product of everything in the input except the input value at index i. An empty list should return `[]`. A list of length 1 should return`[1]` no matter what the input element is.

Example:

```Go
input := []int{2,3,4,5}

expectedOutput := []int{60,40,30,24}
```

## Test

Golang

```bash
cd Go/
go test -v
```
