package fuga

type FugaList struct {
	items []*Fuga
}

func (l *FugaList) Items() []*Fuga {
	return l.items
}

func NewList(items []*Fuga) *FugaList {
	res := new(FugaList)
	res.items = items
	return res
}
func (l *FugaList) Len() int {
	return len(l.items)
}
