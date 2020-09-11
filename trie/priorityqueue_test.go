package trie

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	items := []*Item{
		{value: "gutenberg", priority: 3},
		{value: "you", priority: 2},
		{value: "theory", priority: 1},
		{value: "the", priority: 4},
	}

	pq := make(PriorityQueue, len(items))

	for i, item := range items {
		pq[i] = item
		pq[i].index = i
	}

	heap.Init(&pq)

	item := &Item{
		value:    "homewrecker",
		priority: 25,
	}
	heap.Push(&pq, item)

	output := make([]string, 0, pq.Len())
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		output = append(output, item.value)
	}

	expected := []string{"homewrecker", "the", "gutenberg", "you", "theory"}

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("unexpected priority: got: %v, expected: %v", output, expected)
	}
}
