package main

func main() {

}

func rotateArray(arr []int, k int) []int {
	if k < 0 || len(arr) == 0 {
		return arr
	}

	r := len(arr) - k%len(arr)
	arr = append(arr[r:], arr[:r]...)

	return arr
}

/*
func rotateArray(arr []int, k int) []int {
	k = k % len(arr)

	result := append(arr[len(arr)-k:], arr[:len(arr)-k]...)
	for i := 0; i < len(arr); i++ {
		arr[i] = result[i]
	}

	return arr
}
*/
