package main

import (
	"strings"
	"testing"
)

type assertErr func(error, *testing.T)

func Test_Simple_Rotation_By_Negative(t *testing.T) {
	// Arrange
	input := "abcdefghijklmnopqrstuvwxyz"
	expected := ""
	rotation := -3

	// Act
	t.Logf(`Input string:    %v`, input)
	t.Logf(`Rotation index:  %v`, rotation)
	output, err := CaesarCipher(input, rotation)

	// Assert
	assert(expected, output, err, assertError, t)
}
func Test_Simple_Rotation_By_3(t *testing.T) {
	// Arrange
	input := "abcdefghijklmnopqrstuvwxyz"
	expected := "defghijklmnopqrstuvwxyzabc"
	rotation := 3

	// Act
	t.Logf(`Input string:    %v`, input)
	t.Logf(`Rotation index:  %v`, rotation)
	output, err := CaesarCipher(input, rotation)

	// Assert
	assert(expected, output, err, assertNoError, t)
}

func Test_Simple_Rotation_By_1_Change_Only_Letters(t *testing.T) {
	// Arrange
	input := "abc-defghi *jk* %lyz"
	expected := "bcd-efghij *kl* %mza"
	rotation := 1

	// Act
	t.Logf(`Input string:    %v`, input)
	t.Logf(`Rotation index:  %v`, rotation)
	output, err := CaesarCipher(input, rotation)

	// Assert
	assert(expected, output, err, assertNoError, t)
}

func assert(expected, output string, err error, pred assertErr, t *testing.T) {
	t.Logf(`Expected string: %v`, expected)
	t.Logf(`Output string:   %v`, output)

	if strings.Compare(expected, output) != 0 {
		t.Error(`Strings does not match.`)
	}

	pred(err, t)
}

func assertError(err error, t *testing.T) {
	if err == nil {
		t.Error(`Expected error but is nil.`)
	}
	t.Logf(`Asserting error %v`, err)
}

func assertNoError(err error, t *testing.T) {
	if err != nil {
		t.Errorf(`No error expected but found error: %v`, err)
	}
	t.Logf(`Asserting no error: %v`, err)
}
