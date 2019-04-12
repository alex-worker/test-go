package main

import (
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
    loadInfo := def.LoadInfo{ 
        MapName: mapName,
        FontName: fontName,
        ResourceFolder: resPath,
        ScreenSize: screenSize,
    }
    fmt.Println("Hello!")
    engine.Init(loadInfo)
    engine.Run()
}
