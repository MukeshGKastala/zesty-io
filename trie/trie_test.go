package trie

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		input string
		err   error
	}{
		{input: "venkata", err: nil},
		{input: "AJAI", err: nil},
		{input: "Mukesh", err: nil},
		{input: "", err: ErrEmptyEntity},
		{input: "1234", err: ErrUnexpectedCharacters},
		{input: "abc5d", err: ErrUnexpectedCharacters},
		{input: "[abcd]", err: ErrUnexpectedCharacters},
	}

	prefixTree := New()

	for _, c := range cases {
		err := prefixTree.Insert(c.input)

		if !errors.Is(err, c.err) {
			t.Errorf("input: %v, got: %v, expected: %v", c.input, err, c.err)
		}
	}
}

func TestSubTrie(t *testing.T) {

	words := []string{
		"gutenberg",
		"gutenberg",
		"GUTENBERG",
		"you",
	}

	prefixTree := New()

	for _, word := range words {
		prefixTree.Insert(word)
	}

	cases := []struct {
		input       string
		expectedErr error
	}{
		{input: "gutenberg", expectedErr: nil},
		{input: "GUTENBERG", expectedErr: nil},
		// Partial word
		{input: "yo", expectedErr: nil},
		// Special characters
		{input: "ignore&", expectedErr: ErrUnexpectedCharacters},
		// Nonexistent word
		{input: "zzz", expectedErr: ErrNonexistentPrefix},
		{input: "yoou", expectedErr: ErrNonexistentPrefix},
		{input: "", expectedErr: ErrEmptyEntity},
	}

	for _, c := range cases {
		// Output is tested in `node_test.go`
		_, err := prefixTree.getSubTrie(c.input)

		if !errors.Is(err, c.expectedErr) {
			t.Errorf("unexpected error: input: %v, got: %v, expected: %v", c.input, err, c.expectedErr)
		}
	}
}

func TestAutocomplete(t *testing.T) {

	words := []string{
		"the",
		"theory",
		"theory",
		"Theatric",
		"theatric",
		"theatric",
	}

	prefixTree := New()

	for _, word := range words {
		prefixTree.Insert(word)
	}

	cases := []struct {
		input         string
		k             uint
		expectedWords []string
	}{
		{input: "th", k: 2, expectedWords: []string{"theatric", "theory"}},
		{input: "the", k: 3, expectedWords: []string{"theatric", "theory", "the"}},
		{input: "", k: 5, expectedWords: []string{}},
	}

	for _, c := range cases {
		output := prefixTree.Autocomplete(c.input, c.k)

		if !reflect.DeepEqual(output, c.expectedWords) {
			t.Errorf("unexpected autocomplete: input: `%v`, got: %v, expected: %v", c.input, output, c.expectedWords)
		}
	}
}

func TestAutocompleteShakespeare(t *testing.T) {
	f, err := os.Open("./../data/shakespeare-sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	prefixTree := New()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		// TODO(MukeshKastala): Output failed inserts to file for inspection
		prefixTree.Insert(scanner.Text())
	}

	// Test cases are dependent on the sample of Shakespeare's work
	cases := []struct {
		input         string
		k             uint
		expectedWords []string
	}{
		{input: "th", k: 5, expectedWords: []string{"the", "thy", "thou", "that", "thee"}},
	}

	for _, c := range cases {
		output := prefixTree.Autocomplete(c.input, c.k)

		if !reflect.DeepEqual(output, c.expectedWords) {
			t.Errorf("unexpected autocomplete: got: %v, expected: %v", output, c.expectedWords)
		}
	}
}
