package main

func main() {

}

func productExclusion(a []int) []int {
	size := len(a)
	var product []int = make([]int, size)

	scalar := 1

	for i := 0; i < size; i++ {

		for j := 0; j < size; j++ {

			if j == i {
				continue
			}

			scalar *= a[j]

		}
		product[i] = scalar
		scalar = 1

	}

	return product
}
