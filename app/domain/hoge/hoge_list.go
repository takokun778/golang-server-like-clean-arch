package hoge

type HogeList []*Hoge

func NewList(list []*Hoge) *HogeList {
	hl := HogeList(list)
	return &hl
}

func (l *HogeList) Len() int {
	return len(*l)
}
