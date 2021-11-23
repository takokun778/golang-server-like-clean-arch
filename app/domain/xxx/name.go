package xxx

type name string

func NewName(value string) (name, error) {
	// TODO Validation
	return name(value), nil
}

func (n name) Value() string {
	return string(n)
}
