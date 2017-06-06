package slab

type Pool interface {
	Alloc(int) []type
	Free([]byte)
}

type NoPool struct{}

func (p *NoPool) Alloc(size int) []byte {
	return make([]byte,size)
}

func (p *NoPool) Free(_ []byte) {}
