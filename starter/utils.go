package starter

import (
	"log"
	"time"
)

func Assert(err error) (exist bool) {
	exist = err != nil
	if exist {
		log.Fatalln(err)
	}
	return
}

func RecursionCall(f func() error, count, duration int, done bool) bool {
	var err error
	if !done {
		err = f()
		count--
	}
	if count > 0 && err == nil {
		return true
	} else if count == 0 && err != nil {
		return true
	} else {
		time.Sleep(time.Duration(duration) * time.Second)
	}
	return RecursionCall(f, count, duration, false)
}
