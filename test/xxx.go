package test

import (
	"xxx/app/domain/xxx"
)

func NewXxx(name string, number int) xxx.Props {
	tName, _ := xxx.NewName(name)
	tNumber, _ := xxx.NewNumber(number)
	return xxx.Create(tName, tNumber).Props()
}
