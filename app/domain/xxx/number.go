package xxx

type number int

func NewNumber(value int) (number, error) {
	return number(value), nil
}

func (n number) Value() int {
	return int(n)
}
