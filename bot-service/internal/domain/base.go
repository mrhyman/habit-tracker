package domain

import "time"

var (
	timeNowFn = func() time.Time {
		return time.Now().UTC()
	}
)
