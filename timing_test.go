package timing

import (
	"testing"
	"time"

	"github.com/arstd/log"
)

func TestTiming(t *testing.T) {
	RemindFunc = func(item ...*Item) {
		log.Infof("log handler: %+v", item[0])
	}

	q := make(Queue, 0)
	when := time.Now().Add(2 * time.Second)
	q = append(q, &Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_2"})
	when = when.Add(5 * time.Second)
	q = append(q, &Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_7"})
	when = when.Add(-3 * time.Second)
	q = append(q, &Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_4"})
	when = when.Add(-3 * time.Second)
	q = append(q, &Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_1"})

	Init(q...)

	when = when.Add(8 * time.Second)
	Add(&Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_9"})
	when = when.Add(-6 * time.Second)
	Add(&Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_3"})
	when = when.Add(5 * time.Second)
	Add(&Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_8"})
	when = when.Add(-3 * time.Second)
	Add(&Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_5"})

	time.Sleep(10 * time.Second)

	when = when.Add(7 * time.Second)
	Add(&Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_12"})
	when = when.Add(-1 * time.Second)
	Add(&Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_11"})
	when = when.Add(3 * time.Second)
	Add(&Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_14"})

	when = when.Add(-30 * time.Second)
	Add(&Item{Timed: uint32(when.Unix()), Event: "test", Param: "label_-30"})

	time.Sleep(5 * time.Second)
}
