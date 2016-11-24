package timing

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	items := map[string]uint32{
		"banana": 3, "apple": 2, "pear": 4, "pear2": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(Queue, len(items))
	i := 0
	for param, when := range items {
		pq[i] = &Item{
			Param: param,
			Timed: when,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		Event: "orange",
		Param: "xatarstiarsd",
		Timed: 13,
	}
	heap.Push(&pq, item)

	// Insert a new item and then modify its priority.
	item = &Item{
		Event: "banana",
		Param: "xatarstiarsd",
		Timed: 9,
	}
	heap.Push(&pq, item)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%+v\n", item)
	}
	// x := heap.Pop(&pq) // panic
}
