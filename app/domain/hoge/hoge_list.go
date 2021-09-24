package hoge

type HogeList struct {
	items []*Hoge
}

func (l *HogeList) Items() []*Hoge {
	return l.items
}

func NewList(items []*Hoge) *HogeList {
	res := new(HogeList)
	res.items = items
	return res
}

func (l *HogeList) Len() int {
	return len(l.items)
}
