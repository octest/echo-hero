package main

import (
	h "echo-hero/handler"
	"time"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		//zap.String("url", url),
		//zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", h.Hello)
	e.GET("/test", h.Test)
	e.GET("/test2", h.Test2)
	e.Logger.Fatal(e.Start(":8080"))
}
