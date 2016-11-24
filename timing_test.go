package timing

import (
	"testing"
	"time"

	"github.com/arstd/log"
	"github.com/pborman/uuid"
)

func TestTiming(t *testing.T) {
	RemindFunc = func(items ...*Item) {
		log.JSON("remind", items)
	}

	when := uint32(time.Now().Add(time.Second).Unix())

	// load data
	loaded := []*Item{
		{Id: uuid.New(), Timed: when + 2, Event: "test", Param: "2"},
		{Id: uuid.New(), Timed: when + 10, Event: "test", Param: "10"},
		{Id: uuid.New(), Timed: when + 2, Event: "test", Param: "2"},
		{Id: uuid.New(), Timed: when + 4, Event: "test", Param: "4"},
	}
	Init(loaded...)

	loaded = []*Item{
		{Id: uuid.New(), Timed: when + 3, Event: "test", Param: "3"},
		{Id: uuid.New(), Timed: when + 5, Event: "test", Param: "5"},
	}
	Init(loaded...) // ingored

	Add(&Item{Timed: when + 2, Event: "test", Param: "2"})
	Add(&Item{Timed: when + 5, Event: "test", Param: "5"})
	Add(&Item{Timed: when + 9, Event: "test", Param: "9"})

	time.Sleep(10 * time.Second)

	Add(&Item{Timed: when + 14, Event: "test", Param: "14"})
	Add(&Item{Timed: when + 12, Event: "test", Param: "12"})
	Add(&Item{Timed: when + 14, Event: "test", Param: "14"})

	time.Sleep(5 * time.Second)
}
