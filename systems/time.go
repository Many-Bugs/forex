package systems

import "time"

func NowInUNIX() time.Time {
	return time.Unix(time.Now().Unix(), 0)
}
