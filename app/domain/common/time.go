package common

import (
	"time"
)

type Time time.Time

func (t Time) Value() time.Time {
	return time.Time(t)
}

func NewTime(value time.Time) Time {
	return Time(value)
}

func Now() Time {
	time := time.Now().UTC()
	return NewTime(time)
}
