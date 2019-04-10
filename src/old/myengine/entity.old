package myengine

import (
	"sync/atomic"
)

var (
	idInc uint64
)	

type BasicEntity struct {
	// Entity ID.
	id       uint64
	parent   *BasicEntity
	children []BasicEntity
}

func NewBasic() BasicEntity {
	return BasicEntity{id: atomic.AddUint64(&idInc, 1)}
}

func (e BasicEntity) ID() uint64 {
	return e.id
}