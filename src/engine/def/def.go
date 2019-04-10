package def

type Cell uint32

type Rect struct {
	Width uint32
	Height uint32
}

type LoadInfo struct {
	MapName string
	TileName string
	ScreenSize Rect
}

