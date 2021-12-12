package proto

type Proto struct{}

var singleton = newProto()

func NewProto() *Proto {
	return singleton
}

func newProto() *Proto {
	con := new(Proto)
	return con
}
