package trie

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// ErrUnexpectedCharacters means that we encountered non-word entities
var ErrUnexpectedCharacters = errors.New("non-word entities are ignored")

func alphabetic(r rune) bool {
	return (('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z'))
}

// TODO(MukeshKastala): Don't error on words followed by "valid" punctuation (e.g. "Gutenberg,")
func toLower(word string) (string, error) {
	// Builder pattern was taken from `strings.ToLower()`
	var b strings.Builder
	b.Grow(len(word))

	for _, r := range word {
		if !alphabetic(r) {
			return "", fmt.Errorf("word `%s`: %w", word, ErrUnexpectedCharacters)
		}
		b.WriteByte(byte(unicode.ToLower(r)))
	}

	return b.String(), nil
}
