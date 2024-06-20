package main

import (
	"fmt"
	"new_app/src/engine"
	. "new_app/src/utils"
)

func main() {
	fmt.Println("Hello, World!")
	dataPath := "./data"
	e, err := engine.New(dataPath)
	PanicAssert(err)
	err = e.Run()
	PanicAssert(err)
}
