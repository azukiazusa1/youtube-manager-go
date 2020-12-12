package routes

import (
	"github.com/azukiazusa1/youtube-manager-go/web/api"
	"github.com/labstack/echo"
)

func Ini(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
	}
}
