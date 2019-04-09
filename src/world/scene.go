package world

type Cell uint64

type Scene struct {
	Map *[]Cell
	Width uint32
	Height uint32
}

func LoadScene(filename string) Scene{

	mymap, w, h := loadMap(filename)
	return Scene{ Map: mymap, Width: w, Height: h }

	// return nil

}