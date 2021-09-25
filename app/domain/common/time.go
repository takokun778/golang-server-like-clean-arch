package common

import (
	"time"
)

type Time time.Time

func (t Time) Value() time.Time {
	return time.Time(t)
}

func Now() Time {
	time := time.Now().UTC()
	return Time(time)
}
