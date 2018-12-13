package times

import (
	"time"
)

func TimeParse(str string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, str)

	return t
}
