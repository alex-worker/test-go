package main

import (
	"fmt"
	"runtime"
	"test-go/src/core/Engine"
)

func main() {
	fmt.Println("Hello!")
	eng := Engine.GetEngine("./data")

	runtime.LockOSThread() // примораживаем текущую горутину к текущему треду

	eng.Run()
}
