package main

import (
    "fmt"
    "./myengine"
)

// this is a comment

// type person struct{
    // name string
    // age int
// }

// var myEntity = myengine.BasicEntity{ Id: 0, Parent: nil, Children: nil }

// type Scene struct {}
// func (*Scene) Type() string { return "myGame" }
// func (*Scene) Preload() int { return 600 }
// 
// var myScene Scene

func main() {
    
    var myEntity = myengine.NewBasic()

    // var tom = person{ name: "lol", age: 40 }
    // fmt.Println( myScene.Preload() )
    var ent_id uint64 = myEntity.ID()
    // fmt.Println( myEntity.ID )
    fmt.Println( ent_id )

    // fmt.Println( tom.name )
    // fmt.Println( tom.age )
    
    fmt.Println("Hello-World")
}