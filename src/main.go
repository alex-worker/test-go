package main

import (
    "fmt"
    "./engine"
    "./engine/def"
)

// const mapName string = "data/laboratory3.tmx"
const mapName string = "data/mycastle.tmx"
// const tileName string = "data/tiles_many.png"
// const tileName string = "data/tiles.png"

var screenSize = def.Rect{ 
    Width: 800,
    Height: 600,
}

func main(){
    loadInfo := def.LoadInfo{ 
        MapName: mapName,
        // TileName: tileName,
        ScreenSize: screenSize,
    }
    fmt.Println("Hello!")
    engine.Init(loadInfo)
    engine.Run()
}
