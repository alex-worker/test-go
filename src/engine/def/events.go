package def

// GameEvent действия
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

	// EventPressRestart рестарт
	EventPressRestart

	// EventQuit завершение
	EventQuit
)
