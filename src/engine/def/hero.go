package def

// Direction направление
type Direction uint8
const (
	// DirUp вверх
	DirUp Direction = iota
	// DirLeft влево
	DirLeft
	// DirRight вправо
	DirRight
	// DirDown вниз
	DirDown
)

// Action действия
type Action uint8
const (
	// DoStand встать
	DoStand Direction = iota
	// DoGet взять
	DoGet
	// DoDrop бросить
	DoDrop
)

// Hero герои и прочие npc
type Hero struct {
	X uint32
	Y uint32
	Dir Direction
	Sprite Cell
}
