package mytype

type Number int

func (n *Number) Double() {
	*n *= 2
}
