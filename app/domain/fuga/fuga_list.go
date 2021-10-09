package fuga

type FugaList []*Fuga

func NewList(list []*Fuga) *FugaList {
	fl := FugaList(list)
	return &fl
}

func (l *FugaList) Len() int {
	return len(*l)
}
