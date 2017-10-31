package main

import (
	h "echo-hero/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", h.Hello)
	e.GET("/test", h.Test)
	e.Logger.Fatal(e.Start(":8080"))
}
