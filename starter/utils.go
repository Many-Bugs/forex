package starter

import "log"

func Assert(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
