package main

import (
    "fmt"
    // "github.com/veandco/go-sdl2/sdl"
    // "./game"
    "./world"
    "./ui"
)

var scene *world.Scene

func main(){

    fmt.Println("Hello!")
    scene = world.LoadScene("game/maps/laboratory3.tmx")

    // fmt.Println( scene.Map )

    ui.Init()

    for true {
        ui.Draw( scene )
        if !ui.Update( ) {
			break
		}
    }

    ui.Destroy()

}
