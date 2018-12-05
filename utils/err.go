package utils

import "log"

func AssertErr(err error) bool {
	if err != nil {
		log.Println(err)
	}
	return err != nil
}
