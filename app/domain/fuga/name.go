package fuga

type Name string

func (n Name) Value() string {
	return string(n)
}

func NewName(value string) Name {
	return Name(value)
}
