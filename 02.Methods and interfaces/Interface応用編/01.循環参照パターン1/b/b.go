package b

type B struct {
	A2 A2
}

func (b *B) Get() string {
	return b.A2.Get()
}
