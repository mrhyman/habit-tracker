package domain

import (
	"sync"
	"time"
)

var (
	onceTimeNowFn sync.Once
	timeNowFn     = func() time.Time {
		return time.Now().UTC()
	}
	testNowUtc = time.Now().Truncate(time.Microsecond).UTC()
)

func setup() {
	onceTimeNowFn.Do(func() {
		timeNowFn = func() time.Time {
			return testNowUtc
		}
	})
}
