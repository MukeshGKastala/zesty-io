package trie

import (
	"errors"
	"testing"
)

func TestAlphabetic(t *testing.T) {
	cases := []struct {
		input    rune
		expected bool
	}{
		// Lowercase
		{input: 'v', expected: true},
		{input: 'a', expected: true},
		{input: 'z', expected: true},
		// Uppercase
		{input: 'Q', expected: true},
		{input: 'A', expected: true},
		{input: 'Z', expected: true},
		// Numeric
		{input: '0', expected: false},
		{input: '3', expected: false},
		// Special characters
		{input: ' ', expected: false},
		{input: '[', expected: false},
		{input: '\n', expected: false},
		{input: '\t', expected: false},
		{input: '&', expected: false},
	}

	for _, c := range cases {
		output := alphabetic(c.input)

		if output != c.expected {
			t.Errorf("input: %v, got: %v, expected: %v", string(c.input), output, c.expected)
		}
	}
}

func TestToLower(t *testing.T) {
	cases := []struct {
		input       string
		expectedStr string
		expectedErr error
	}{
		// Lowercase
		{input: "venkata", expectedStr: "venkata", expectedErr: nil},
		// Uppercase
		{input: "AJAI", expectedStr: "ajai", expectedErr: nil},
		// Uppercase/Lowercase combo
		{input: "Mukesh", expectedStr: "mukesh", expectedErr: nil},
		// Numeric
		{input: "1234", expectedStr: "", expectedErr: ErrUnexpectedCharacters},
		// Alphanumeric
		{input: "abc5d", expectedStr: "", expectedErr: ErrUnexpectedCharacters},
		// Special characters
		{input: "[abcd]", expectedStr: "", expectedErr: ErrUnexpectedCharacters},
	}

	for _, c := range cases {
		output, err := toLower(c.input)
		if !errors.Is(err, c.expectedErr) {
			t.Errorf("unexpected error: input: %v, got: %v, expected: %v", c.input, err, c.expectedErr)
		}

		if output != c.expectedStr {
			t.Errorf("unexpected output: input: %v, got: %v, expected: %v", c.input, output, c.expectedStr)
		}
	}
}
