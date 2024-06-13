package main

func main() {

}

func countDuplicated(arr []int) int {
	length := len(arr)

	if length < 3 {
		return length
	}

	l := 2

	for r := 2; r < length; r++ {
		if arr[l-2] != arr[r] {
			arr[l] = arr[r]
			l++
		}
	}

	return l
}

func removeDuplicated(arr []int) []int {
	var output []int = []int{}
	len := len(arr)

	for i := 0; i < len; i++ {
		if i < len-1 && arr[i] == arr[i+1] {
			continue
		}
		output = append(output, arr[i])
	}

	return output
}
