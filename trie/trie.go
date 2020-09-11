package trie

import (
	"errors"
	"fmt"
	"log"
)

// ErrNonexistentPrefix means that the prexif isn't a start to any words in the corpus
var ErrNonexistentPrefix = errors.New("word isn't present in the corpus")

// ErrEmptyEntity means we encountered an empty word/prefix
var ErrEmptyEntity = errors.New("empty entity isn't allowed")

// Trie tracks the root
type Trie struct {
	root *node
}

// New creates an empty Trie
func New() *Trie {
	return &Trie{
		root: &node{},
	}
}

func (t *Trie) getSubTrie(prefix string) (*Trie, error) {
	if prefix == "" {
		return nil, fmt.Errorf("sub-trie: %w", ErrEmptyEntity)
	}

	lower, err := toLower(prefix)
	if err != nil {
		return nil, fmt.Errorf("sub-trie: %w", err)
	}

	node := t.root.getPrefixNode([]rune(lower))
	if node == nil {
		return nil, fmt.Errorf("sub-trie prefix `%s`: %w", prefix, ErrNonexistentPrefix)
	}

	subTrie := &Trie{
		root: node,
	}

	return subTrie, nil
}

// Insert adds valid words to the Trie
func (t *Trie) Insert(word string) (err error) {
	if word == "" {
		err = fmt.Errorf("insert: %w", ErrEmptyEntity)
		return
	}

	lower, err := toLower(word)
	if err != nil {
		err = fmt.Errorf("insert: %w", err)
		return
	}

	t.root.insert([]rune(lower))
	return
}

// Autocomplete gets k frequent prefix matches
func (t *Trie) Autocomplete(prefix string, k uint) []string {
	if prefix == "" {
		return []string{}
	}

	subTrie, err := t.getSubTrie(prefix)
	if err != nil {
		log.Println("autocomplete:", err)
		return []string{}
	}

	return subTrie.root.findKfrequentWords(k)
}
