package main

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)

func main(){
	fmt.Println("go is ok")
	// Echo instance
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, lol! World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}