package fuga

type FugaList []Fuga

func NewList(list []Fuga) FugaList {
	fl := FugaList(list)
	return fl
}

func (fl FugaList) ValuesList() ValuesList {
	vl := make([]Values, 0)
	for _, v := range fl {
		vl = append(vl, v.Values())
	}
	return ValuesList(vl)
}

func (l FugaList) Len() int {
	return len(l)
}
