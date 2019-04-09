package main

import (
    // "fmt"
//     "github.com/veandco/go-sdl2/sdl"
    "./game"
)


func main(){
    game.LoadLevelFromFile("game/maps/level1.map")
}

// import (
//     "fmt"
//     "github.com/veandco/go-sdl2/sdl"
//     "./myengine"
// )

// // this is a comment

// type person struct{
//     name string
//     age int
// }

// type Scene struct {}
// func (*Scene) Type() string { return "myGame" }
// func (*Scene) Preload() int { return 600 }

// var myScene Scene

// // func main() {

//     // if err := sdl.Init(sdl.SDL_INIT_EVERYTHING); err != nil {
// 	// 	panic(err)
// 	// }
// 	// defer sdl.Quit()
    
//     func main() {

//         if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
//             panic(err)
//         }
//         defer sdl.Quit()
    
//         window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
//             800, 600, sdl.WINDOW_SHOWN )
//         if err != nil {
//             panic(err)
//         }
//         defer window.Destroy()
    
//         surface, err := window.GetSurface()
//         if err != nil {
//             panic(err)
//         }
//         surface.FillRect(nil, 0)
    
//         rect := sdl.Rect{0, 0, 200, 200}
//         surface.FillRect(&rect, 0xffff0000)
//         window.UpdateSurface()
    
//         running := true

//         for running {
//             for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
//                 switch event.(type) {
//                 case *sdl.QuitEvent:
//                     println("Quit")
//                     running = false
//                     break
//                 }
//             }
//         }
//     // }
    
//     var myEntity = myengine.NewBasic()

//     var tom = person{ name: "lol", age: 40 }
//     fmt.Println( myScene.Preload() )
//     var ent_id uint64 = myEntity.ID()
//     fmt.Println( myEntity.ID )
//     fmt.Println( ent_id )

//     fmt.Println( tom.name )
//     fmt.Println( tom.age )
    
//     fmt.Println("Hello-World")
// }