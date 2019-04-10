package main

import (
    "fmt"
    "./engine"
)

const mapName string = "data/laboratory3.tmx"
const tileName string = "data/tiles_many.png"

func main(){

    loadInfo := engine.LoadInfo{ MapName: mapName, TileName: tileName }
    fmt.Println("Hello!")
    engine.Init(loadInfo)
    engine.Run()

}
