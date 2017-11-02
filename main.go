package main

import (
	h "echo-hero/handler"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/labstack/echo"
)

type Config struct {
	DB      string `json: "db"`
	User    string `json: "user"`
	Pass    string `json: "pass"`
	Port    string `json: "port"`
	Debug   bool   `json: "debug"`
	Logfile string `json: "logfile"`
}

func main() {
	config := Config{Port: "8080", Debug: true, Logfile: "logs.log"}

	// let it panic
	raw, _ := ioutil.ReadFile("./config.json")
	json.Unmarshal(raw, &config)

	e := echo.New()

	if !config.Debug {
		f, err := os.OpenFile("logs/"+config.Logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			e.Logger.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		e.Logger.SetOutput(f)
		//e.Logger.SetLevel(log.ERROR)
		e.Logger.SetPrefix("Echo-Hero")
		e.HideBanner = true
	}

	e.Static("/static", "static")
	e.GET("/", h.Hello)
	e.GET("/test", h.Test)

	e.GET("/test2", h.Test2)

	e.Logger.Fatal(e.Start(":" + config.Port))
	//e.Start(":8080")
}
