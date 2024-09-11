package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {
	s := "abc-defghi *jk* %lyz"
	n := 3

	x, _ := CaesarCipher(s, n)

	fmt.Println(x)
}

func CaesarCipher(s string, n int) (string, error) {

	if n < 0 {
		errorMessage := `shift number is a negative: %v`
		log.Printf(errorMessage, n)
		return "", fmt.Errorf(errorMessage, n)
	}

	alphabetMap := CreateMap()
	c := []rune{}

	for _, v := range s {
		if unicode.IsLetter(v) {
			index := alphabetMap[v] + n
			if index > len(alphabetMap) {
				index = index - len(alphabetMap)
			}
			c = append(c, getRune(alphabetMap, index))
		} else {
			c = append(c, v)
		}
	}

	return string(c), nil
}

func getRune(alphabetMap map[rune]int, index int) rune {
	var letter rune = 'a'

	for v, i := range alphabetMap {
		if i == index {
			letter = v
			break
		}
	}

	return letter
}

func CreateMap() map[rune]int {
	alphabetMap := make(map[rune]int)

	for i := 0; i < 26; i++ {
		letter := rune('a' + i)
		alphabetMap[letter] = i + 1
	}

	return alphabetMap
}
