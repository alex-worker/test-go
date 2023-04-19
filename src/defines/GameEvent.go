package defines

type GameEvent uint8

const (

	// EventNo нет событий
	EventNo GameEvent = iota

	// EventPressUp вверх
	EventPressUp
	// EventPressDown вниз
	EventPressDown
	// EventPressLeft влево
	EventPressLeft
	// EventPressRight вправо
	EventPressRight
	// EventPressFire огонь
	EventPressFire

	// EventQuit завершение
	EventQuit
)
