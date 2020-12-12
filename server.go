package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.Now()
	e.Logger.Fatal(e.start(":8080"))
}
