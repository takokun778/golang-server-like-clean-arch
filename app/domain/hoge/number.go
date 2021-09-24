package hoge

type Number int32

func (n Number) Value() int32 {
	return int32(n)
}

func NewNumber(value int32) Number {
	return Number(value)
}
