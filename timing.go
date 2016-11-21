package timing

import (
	"container/heap"
	"fmt"
	"time"
)

type HandlerFunc func(when time.Time, item *Item)

var Handler = func(when time.Time, item *Item) {
	fmt.Printf("default handler, timer %s, item %#v\n", when, item)
}

var stage = make(chan *Item)

var inited = make(chan struct{})

// Init 初始化定时器，并启动
func Init(q Queue) {
	if q == nil {
		q = make(Queue, 0, 1024)
	}
	go start(q)
}

// Add 用来设置定时项，必须先 Init，否则 Add 会一直等待 Init
func Add(items ...*Item) {
	<-inited

	for _, item := range items {
		// process
		stage <- item
	}
}

func start(q Queue) {
	var (
		min   *Item
		timer *time.Timer
	)
	heap.Init(&q)

	if len(q) > 0 {
		min = heap.Pop(&q).(*Item)
		timer = time.NewTimer(time.Unix(int64(min.When), 0).Sub(time.Now()))
	} else {
		timer = time.NewTimer(24 * time.Hour)
	}
	close(inited)

	for {
		select {
		case item := <-stage:
			if min == nil {
				// do nothing
			} else if item.When < min.When {
				heap.Push(&q, min)
			} else if item.When > min.When {
				heap.Push(&q, item)
				break
			}

			min = item

			until := time.Unix(int64(min.When), 0)
			dur := until.Sub(time.Now())
			timer.Reset(dur)

		case now := <-timer.C:
			go Handler(now, min)

			if q.Len() == 0 {
				min = nil
				// fmt.Println("no item in heap")
				timer.Reset(24 * time.Hour)
				break
			}

			h := heap.Pop(&q)
			if h == nil {
				min = nil
				// fmt.Println("pop nil from heap")
				timer.Reset(24 * time.Hour)
				break
			}

			min = h.(*Item)

			until := time.Unix(int64(min.When), 0)
			dur := until.Sub(time.Now())
			timer.Reset(dur)
		}
	}
}
