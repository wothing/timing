package timing

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	// items := map[string]uint32{
	// 	"banana": 3, "apple": 2, "pear": 4,
	// }

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(Queue, 0)
	// i := 0
	// for label, when := range items {
	// 	pq[i] = &Item{
	// 		Label: label,
	// 		When:  when,
	// 		index: i,
	// 	}
	// 	i++
	// }
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		Label: "orange",
		When:  13,
	}
	heap.Push(&pq, item)

	// Insert a new item and then modify its priority.
	item = &Item{
		Label: "banana",
		When:  9,
	}
	heap.Push(&pq, item)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%d %s ", item.When, item.Label)
	}
}
