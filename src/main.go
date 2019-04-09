package main

import (
    "fmt"
    // "github.com/veandco/go-sdl2/sdl"
    // "./game"
    "./world"
    "./ui"
)

// var level *game.Level
var scene world.Scene

func main(){

    fmt.Println("Hello!")
    // level = game.LoadLevelFromFile("game/maps/level1.map")
    scene = world.LoadScene("game/maps/laboratory3.tmx")    

    // fmt.Println( myMap[:1] )

    ui.Init()

    for true {
        ui.Draw( scene )
        if !ui.Update() {
			break
		}
    }

    ui.Destroy()

}
