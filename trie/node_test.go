package trie

import (
	"reflect"
	"testing"
)

func TestNode(t *testing.T) {
	inputs := [][]rune{
		[]rune{'g', 'u', 't', 'e', 'n', 'b', 'e', 'r', 'g'},
		[]rune{'y', 'o', 'u'},
	}

	root := newNode()

	for _, letters := range inputs {
		root.insert(letters)
	}

	cases := []struct {
		input        []rune
		expectedNode *node
	}{
		// Nonexistent word
		{input: []rune{'z', 'z', 'z'}, expectedNode: nil},
		// Empty
		{input: []rune{}, expectedNode: root},
	}

	for _, c := range cases {
		output := root.getPrefixNode(c.input)

		if output != c.expectedNode {
			t.Errorf("unexpected node: input: %v, got: %v, expected: %v", string(c.input), output, c.expectedNode)
		}

	}
}

func TestNodeMembers(t *testing.T) {

	inputs := [][]rune{
		[]rune{'g', 'u', 't', 'e', 'n', 'b', 'e', 'r', 'g'},
		[]rune{'g', 'u', 't', 'e', 'n', 'b', 'e', 'r', 'g'},
		[]rune{'y', 'o', 'u'},
	}

	root := newNode()

	for _, letters := range inputs {
		root.insert(letters)
	}

	cases := []struct {
		input             []rune
		expectedWord      string
		expectedFrequency int
	}{
		{
			input:             []rune{'g', 'u', 't', 'e', 'n', 'b', 'e', 'r', 'g'},
			expectedWord:      "gutenberg",
			expectedFrequency: 2,
		},
		// Partial word
		{
			input:             []rune{'y', 'o'},
			expectedWord:      "",
			expectedFrequency: 0,
		},
	}

	for _, c := range cases {
		output := root.getPrefixNode(c.input)

		if output.frequency != c.expectedFrequency {
			t.Errorf("unexpected frequency: input: %v, got: %v, expected: %v", c.input, output.frequency, c.expectedFrequency)
		}

		if output.word != c.expectedWord {
			t.Errorf("unexpected value: input: %v, got: %v, expected: %v", c.input, string(output.word), string(c.expectedFrequency))
		}
	}
}

func TestFindKfrequentWords(t *testing.T) {

	inputs := [][]rune{
		[]rune{'g', 'u', 't', 'e', 'n', 'b', 'e', 'r', 'g'},
		[]rune{'g', 'u', 't', 'e', 'n', 'b', 'e', 'r', 'g'},
		[]rune{'g', 'u', 't', 'e', 'n', 'b', 'e', 'r', 'g'},
		[]rune{'y', 'o', 'u'},
		[]rune{'y', 'o', 'u'},
		[]rune{'t', 'h', 'e', 'o', 'r', 'y'},
		[]rune{'t', 'h', 'e', 'o', 'r', 'y'},
		[]rune{'t', 'h', 'e', 'o', 'r', 'y'},
		[]rune{'t', 'h', 'e', 'o', 'r', 'y'},
		[]rune{'t', 'h', 'e'},
	}

	root := newNode()

	for _, letters := range inputs {
		root.insert(letters)
	}

	// Order matters and is dependent on word frequency
	cases := []struct {
		input         []rune
		expectedWords []string
	}{
		{
			input:         []rune{},
			expectedWords: []string{"theory", "gutenberg", "you", "the"},
		},
		{
			input:         []rune{'t', 'h'},
			expectedWords: []string{"theory", "the"},
		},
		{
			input:         []rune{'y'},
			expectedWords: []string{"you"},
		},
		{
			input:         []rune{'y', 'o', 'u'},
			expectedWords: []string{"you"},
		},
	}

	for _, c := range cases {
		output := root.getPrefixNode(c.input).findKfrequentWords(5)

		if !reflect.DeepEqual(output, c.expectedWords) {
			t.Errorf("unexpected words: input: %v, got: %v, expected: %v", string(c.input), output, c.expectedWords)
		}
	}
}
