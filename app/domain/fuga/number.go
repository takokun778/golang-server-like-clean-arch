package fuga

type Number int32

func (n Number) Value() int32 {
	return int32(n)
}
