package entity

var (
	idInc uint64
)

type BasicEntity struct {
	// Entity ID.
	id       uint64
	parent   *BasicEntity
	children []BasicEntity
}
