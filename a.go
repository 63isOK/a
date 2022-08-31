package a

type Obj struct {
	A, B, C int
}

func (o Obj) Sum() int {
	return o.A + o.B
}
