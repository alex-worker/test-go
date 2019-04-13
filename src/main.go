package main

import (
    "flag"
    "fmt"
    "./engine"
    "./engine/def"
)

const fontName string = "CaslonBold"
const mapName string = "swamp.tmx"
// const mapName string = "laboratory3.tmx"
const resPath string = "data/"

var screenSize = def.Rect{
    Width: 800,
    Height: 600,
}

func main(){

    resDir := resPath

    flag.StringVar(&resDir, "dir", resPath, "a string var")
    flag.Parse()

    loadInfo := def.LoadInfo{ 
        MapName: mapName,
        FontName: fontName,
        ResourceFolder: resDir,
        ScreenSize: screenSize,
    }
    fmt.Println("Hello!")
    engine.Init(loadInfo)
    engine.Run()
}
