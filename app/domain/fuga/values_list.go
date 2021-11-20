package fuga

type ValuesList []Values

func NewValuesList(list []Values) ValuesList {
	vl := ValuesList(list)
	return vl
}
