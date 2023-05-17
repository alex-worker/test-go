package Map

import (
	. "test-go/src/defines"
)

// Layer слой
type Layer struct {
	Data []Cell
	Name string
	W    uint64
	H    uint64
}
