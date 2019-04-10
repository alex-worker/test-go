package world

type Cell uint32

type Scene struct {
	Map *[][]Cell
	Width uint32
	Height uint32
}

func LoadScene(filename string) *Scene{

	mymap, w, h := loadTmxMap(filename)
	scene := Scene{ Map: mymap, Width: w, Height: h }
	return &scene;

}