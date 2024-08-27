package main

import (
	"fmt"
	"log"
	"unicode"
)

/*

/// DESCRIPTION

Julius Caesar protected his confidential information by encrypting it using a cipher.
Caesar's cipher shifts each letter by a number of letters. For example, in the case of
a rotation by 3, w, x, y and z would map to z, a, b and c.

Original alphabet:      abcdefghijklmnopqrstuvwxyz
Alphabet rotated +3:    defghijklmnopqrstuvwxyzabc

/// REQUIREMENTS.

1. If the shift takes you past the end of the alphabet, just rotate back to the front of the alphabet.
2. If the shift number is a negative number, log an error.
3. The cipher only encrypts letters; all other characters remain unencrypted.
4. Build a function CaesarCipher(s string, n int) (string, error) to handle your implementation.
5. Create unit tests to validate previous requirements.

/// EXAMPLE.

s := "abc-defghi *jk* %lyz"
n := 3

The alphabet is rotated by 3, matching the mapping above.

The encrypted string is: "def-ghijkl *mn* obc"

/// BONUS

Implement a concurrent mechanism to handle the execution of Caesar's Cipher over multiple strings

*/

func main() {
	fmt.Println("Welcome to the Code Challenge")

	s := "abc-defghi *jk* %lyz"
	n := 3

	x, _ := CaesarCipher(s, n)

	fmt.Println(x)
}

func CaesarCipher(s string, n int) (string, error) {

	if n < 0 {
		log.Printf("Shift number is a negative.", n)
		return "", nil
	}

	alphabetMap := CreateMap()
	c := []string{}

	for _, v := range s {
		if unicode.IsLetter(v) {
			fmt.Printf("Char at %v, %v \n", string(v), alphabetMap[string(v)])
			c = append(c, string(v))
		} else {

		}

	}

	fmt.Println(c)

	return "", nil

}

func CreateMap() map[string]int {

	alphabetMap := make(map[string]int)
	for i := 0; i < 26; i++ {
		letter := string('a' + i)
		alphabetMap[letter] = i + 1
	}

	return alphabetMap
}
