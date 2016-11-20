package timing

import "container/heap"

type Item struct {
	index int
	When  uint32
	Label string
}

type Queue []*Item

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].When < q[j].When
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *Queue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*q)
	*q = append(*q, item)
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

func (q *Queue) update(item *Item, time uint32, label string) {
	item.Label = label
	item.When = time
	heap.Fix(q, item.index)
}
