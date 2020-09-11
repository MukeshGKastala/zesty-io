package trie

import (
	"container/heap"
)

type node struct {
	word      string
	children  map[rune]*node
	frequency int
}

// Added for testing
func newNode() *node {
	return &node{}
}

// Perform pre-order traversal of Node and insert
// (word, frequency) pair into max-heap
func (n *node) preorder(pq *PriorityQueue) {
	// Frequency greater than 0 correlates to end-of-word
	if n.frequency > 0 {
		item := &Item{
			value:    n.word,
			priority: n.frequency,
		}
		heap.Push(pq, item)
	}

	for _, c := range n.children {
		c.preorder(pq)
	}
}

// Find first k-maximum occurring words from given Node
func (n *node) findKfrequentWords(k uint) []string {
	words := make([]string, 0, k)

	// Max-heap
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	n.preorder(&pq)

	for k > 0 && pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		words = append(words, item.value)
		k--
	}

	return words
}

func (n *node) insert(letters []rune) {
	current := n

	for _, r := range letters {
		if _, ok := current.children[r]; !ok {
			if current.children == nil {
				current.children = make(map[rune]*node)
			}
			current.children[r] = &node{}
		}
		current = current.children[r]
	}

	if current.word == "" {
		current.word = string(letters)
	}
	current.frequency++
}

func (n *node) getPrefixNode(letters []rune) *node {
	if len(letters) == 0 {
		return n
	}

	// Remove leading letter
	child, letters := letters[0], letters[1:]

	if _, ok := n.children[child]; !ok {
		return nil
	}

	return n.children[child].getPrefixNode(letters)
}
