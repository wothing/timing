package timing

import (
	"container/heap"
	"fmt"
	"sync"
	"time"

	"github.com/pborman/uuid"
)

type HandlerFunc func(items ...*Item)

var (
	// PersistFunc 落地定时项，以便系统重启后，恢复定时项
	PersistFunc HandlerFunc = func(items ...*Item) {}

	// DeleteFunc 时间到了，从落地库中移除定时项
	DeleteFunc HandlerFunc = func(items ...*Item) {}

	// RemindFunc 时间到了，提醒，处理定时项
	RemindFunc HandlerFunc = func(items ...*Item) {
		fmt.Printf("default remind: %#v\n", items)
	}
)

var (
	stage  = make(chan *Item)
	inited = make(chan struct{})
)

var once sync.Once

// Init 初始化定时器，并启动
func Init(items ...*Item) {
	once.Do(func() {
		q := make(Queue, len(items), 1024)
		for i, item := range items {
			// item.Id = uuid.New()
			// PersistFunc(item) // 初始化最大可能是从落地库中恢复定时项，无需重复落地
			q[i] = item
		}
		go start(q)
	})
}

// Add 用来添加定时项并设置ID，必须先 Init，否则 Add 会一直等待 Init
func Add(items ...*Item) {
	<-inited

	now := uint32(time.Now().Unix())
	for _, item := range items {
		if item.Timed < now {
			RemindFunc(item)
			continue
		}
		if item.Id == "" {
			item.Id = uuid.New()
		}
		stage <- item
	}
}

func start(q Queue) {
	heap.Init(&q)

	var min *Item
	var timer = time.NewTimer(24 * time.Hour)

	if len(q) > 0 {
		min = heap.Pop(&q).(*Item)
		timer.Reset(time.Unix(int64(min.Timed), 0).Sub(time.Now()))
	}

	close(inited)
	for {
		select {
		case item := <-stage:
			PersistFunc(item)

			if min == nil {
				// do nothing
			} else if item.Timed < min.Timed {
				heap.Push(&q, min)
			} else {
				heap.Push(&q, item)
				break
			}

			min = item

			timer.Reset(time.Unix(int64(min.Timed), 0).Sub(time.Now()))

		case <-timer.C:

			if min != nil {
				timeup := []*Item{min}

				for {
					if q.Len() == 0 {
						min = nil
						// fmt.Println("no item in heap")
						break
					}

					temp := heap.Pop(&q).(*Item)
					if temp.Timed != min.Timed {
						min = temp
						break
					}
					timeup = append(timeup, temp)
				}

				go func(timeup []*Item) {
					RemindFunc(timeup...)
					DeleteFunc(timeup...)
				}(timeup)
			}

			if min == nil {
				timer.Reset(24 * time.Hour)
				break
			}

			timer.Reset(time.Unix(int64(min.Timed), 0).Sub(time.Now()))
		}
	}
}
