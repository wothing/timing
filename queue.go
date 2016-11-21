package timing

type Item struct {
	Id    string
	Timed uint32
	Event string
	Param string
}

type Queue []*Item

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].Timed < q[j].Timed
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	item := x.(*Item)
	*q = append(*q, item)
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	// item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}
