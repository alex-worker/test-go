package main

import (
	"flag"
	"fmt"
	"test-go/src/engine"
	"test-go/src/engine/def"
)

const fontName string = "CaslonBold"
const mapName string = "swamp.tmx"

// const mapName string = "laboratory3.tmx"
// const mapName string = "mycastle.tmx"

const ResourcePathDefault string = "data/"

var screenSize = def.Size{
	Width:  800,
	Height: 600,
}

func main() {

	var resDir string

	flag.StringVar(&resDir, "dir", ResourcePathDefault, "directory path")
	flag.Parse()

	loadInfo := def.LoadInfo{
		MapName:        mapName,
		FontName:       fontName,
		ResourceFolder: resDir,
		ScreenSize:     screenSize,
	}

	fmt.Println("Hello!")
	myEngine := engine.Create(loadInfo)
	myEngine.Run()
}
