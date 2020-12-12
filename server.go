package main

import (
	"github.com/azukiazusa1/youtube-manager-go/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
