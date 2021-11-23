package xxx

import (
	tm "time"
)

type time tm.Time

func NewTime(value tm.Time) time {
	return time(value)
}

func (t time) Value() tm.Time {
	return tm.Time(t)
}

func TimeNow() time {
	t := tm.Now().UTC()
	return time(t)
}
