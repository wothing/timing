package timing

import (
	"testing"
	"time"

	"github.com/arstd/log"
)

func TestTiming(t *testing.T) {
	Handler = func(t time.Time, item *Item) {
		log.Infof("log handler, timer %s", t)
		log.JSON(min)
	}

	when := time.Now().Add(15 * time.Second)
	Set(when, "===> 15")

	when = when.Add(-10 * time.Second)
	Set(when, "===>  5")

	when = when.Add(5 * time.Second)
	Set(when, "===> 10")

	time.Sleep(20 * time.Second)

	when = when.Add(15 * time.Second)
	Set(when, "===> 25")

	time.Sleep(20 * time.Second)
}

//
// func TestTimer(tt *testing.T) {
// 	t := time.NewTimer(100 * time.Second)
// 	reset := make(chan time.Duration)
// 	go func() {
// 		reset <- 5 * time.Second
// 		reset <- 3 * time.Second
// 		reset <- 2 * time.Second
// 	}()
// 	for {
// 		select {
// 		case d := <-reset:
// 			log.Debugf("reset %s", d)
// 			t.Reset(d)
// 		case now := <-t.C:
// 			log.Debug(now.Format("2006-01-02 15:04:05"))
//
// 			// if !t.Stop() {
// 			// 	now := <-t.C
// 			// 	log.Warnf("stop faild, timer %s", now)
// 			// }
// 			t.Reset(100 * time.Second)
// 		}
// 	}
//
// }
