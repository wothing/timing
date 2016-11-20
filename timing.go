package timing

import (
	"container/heap"
	"fmt"
	"time"

	"github.com/arstd/log"
)

type HandlerFunc func(when time.Time, item *Item)

var Handler = func(when time.Time, item *Item) {
	fmt.Printf("default handler, timer %s, item %#v\n", when, item)
}

var q Queue
var reset = make(chan *Item)
var min *Item // 最近的元素

func init() {
	q = make(Queue, 0)
	heap.Init(&q)

	go wait()
}

func wait() {
	timer := time.NewTimer(24 * time.Hour)
	for {
		select {
		case item := <-reset:
			log.Warnf("get item %s %d", item.Label, item.When)
			log.Warnf("min item %#v", min)
			if min == nil {
				// do nothing
			} else if item.When < min.When {
				log.Debugf("heap len %d", q.Len())
				heap.Push(&q, min)
				log.Debugf("min label %s in heap", min.Label)
				log.Debugf("heap len %d", q.Len())
			} else if item.When > min.When {
				heap.Push(&q, item)
				break
			}

			min = item
			until := time.Unix(int64(min.When), 0)
			dur := until.Sub(time.Now())
			timer.Reset(dur)
			log.Debugf("reset to %s, %s", dur, until)

		case now := <-timer.C:
			log.Infof("<<< timing %s", now)
			Handler(now, min)

			if q.Len() == 0 {
				min = nil
				log.Info("no item in heap")
				timer.Reset(24 * time.Hour)
				break
			}

			h := heap.Pop(&q)
			if h == nil {
				log.Error("pop nil from heap")
				timer.Reset(24 * time.Hour)
				break
			}

			min = h.(*Item)
			until := time.Unix(int64(min.When), 0)
			dur := until.Sub(time.Now())
			timer.Reset(dur)
			log.Debugf("reset to %s, %s", dur, until)
		}
	}
}

func Set(when time.Time, label string) {
	SetUnix(uint32(when.Unix()), label)
}

func SetUnix(when uint32, label string) {
	log.Debugf("set %s, %s", label, time.Unix(int64(when), 0))
	item := &Item{
		When:  when,
		Label: label,
	}
	reset <- item
}
