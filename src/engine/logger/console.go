package logger

import "fmt"

type ILogger interface {
	Log(a ...any)
}

type ConsoleLogger struct{}

func (m *ConsoleLogger) Log(a ...any) {
	fmt.Println(a)
}
