package main

import (
	"fmt"
	"test-go/src/core/Engine"
)

func main() {
	fmt.Println("Hello!")
	eng := Engine.GetEngine()
	eng.Run()
}
