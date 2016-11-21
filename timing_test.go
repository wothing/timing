package timing

import (
	"testing"
	"time"

	"github.com/arstd/log"
)

func TestTiming(t *testing.T) {
	Handler = func(t time.Time, item *Item) {
		log.Infof("log handler, timer %s, %+v", t, item)
	}

	q := make(Queue, 0)
	when := time.Now().Add(2 * time.Second)
	q = append(q, &Item{When: uint32(when.Unix()), Label: "label_2"})
	when = when.Add(5 * time.Second)
	q = append(q, &Item{When: uint32(when.Unix()), Label: "label_7"})
	when = when.Add(-3 * time.Second)
	q = append(q, &Item{When: uint32(when.Unix()), Label: "label_4"})
	when = when.Add(-3 * time.Second)
	q = append(q, &Item{When: uint32(when.Unix()), Label: "label_1"})
	Init(q)

	when = when.Add(8 * time.Second)
	Add(&Item{When: uint32(when.Unix()), Label: "label_9"})
	when = when.Add(-6 * time.Second)
	Add(&Item{When: uint32(when.Unix()), Label: "label_3"})
	when = when.Add(5 * time.Second)
	Add(&Item{When: uint32(when.Unix()), Label: "label_8"})
	when = when.Add(-3 * time.Second)
	Add(&Item{When: uint32(when.Unix()), Label: "label_5"})

	time.Sleep(10 * time.Second)

	when = when.Add(7 * time.Second)
	Add(&Item{When: uint32(when.Unix()), Label: "label_12"})
	when = when.Add(-1 * time.Second)
	Add(&Item{When: uint32(when.Unix()), Label: "label_11"})
	when = when.Add(3 * time.Second)
	Add(&Item{When: uint32(when.Unix()), Label: "label_14"})

	time.Sleep(5 * time.Second)
}
