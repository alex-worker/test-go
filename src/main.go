package main

import (
    "fmt"
    "./engine"
    "./engine/def"
)

const mapName string = "data/laboratory3.tmx"
const tileName string = "data/tiles_many.png"

const screenWidth uint32 = 640
const screenHeight uint32 = 480

func main(){
    loadInfo := def.LoadInfo{ MapName: mapName, TileName: tileName,
        ScreenSize: def.Rect{ Width: screenWidth, Height:screenHeight },
    }
    fmt.Println("Hello!")
    engine.Init(loadInfo)
    engine.Run()
}
