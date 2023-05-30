package main

import (
	"fmt"
	"runtime"
	"test-go/src/core/Engine"
)

func main() {
	fmt.Println("Hello!")
	eng, err := Engine.GetEngine("./data")
	if err != nil {
		panic(err)
	}

	runtime.LockOSThread() // примораживаем текущую горутину к текущему треду

	eng.Run()
}
