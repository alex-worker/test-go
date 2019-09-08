package main

import (
	"./engine"
	"./engine/def"
	"flag"
	"fmt"
)

const fontName string = "CaslonBold"
const mapName string = "swamp.tmx"

// const mapName string = "laboratory3.tmx"
// const mapName string = "mycastle.tmx"
const resPath string = "data/"

var screenSize = def.Size{
	Width:  800,
	Height: 600,
}

func main() {

	resDir := resPath

	flag.StringVar(&resDir, "dir", resPath, "directory path")
	flag.Parse()

	loadInfo := def.LoadInfo{
		MapName:        mapName,
		FontName:       fontName,
		ResourceFolder: resDir,
		ScreenSize:     screenSize,
	}
	fmt.Println("Hello!")
	engine.Init(loadInfo)
	engine.Run()
}
