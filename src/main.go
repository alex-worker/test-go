package main

import (
    "fmt"
    "./engine"
    "./engine/def"
)

const mapName string = "mycastle.tmx"
// const mapName string = "laboratory3.tmx"
const resPath string = "data/"

var screenSize = def.Rect{ 
    Width: 800,
    Height: 600,
}

func main(){
    loadInfo := def.LoadInfo{ 
        MapName: mapName,
        ResourceFolder: resPath,
        ScreenSize: screenSize,
    }
    fmt.Println("Hello!")
    engine.Init(loadInfo)
    engine.Run()
}
