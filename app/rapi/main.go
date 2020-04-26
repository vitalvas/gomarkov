package main

import (
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/swaggo/echo-swagger"
	"github.com/vitalvas/gomarkov/app/rapi/app"
	_ "github.com/vitalvas/gomarkov/app/rapi/docs"
)

// @title Some API
func main() {
	order := flag.Int("order", 3, "Chain order to use")
	flag.Parse()

	m := app.NewModel(*order)

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/api/model", m.GetModel)
	e.POST("/api/model/learn", m.PostModel)
	e.POST("/api/model/learn/batch", m.PostModelLearnBatch)

	e.GET("/api/model/dump", m.GetModelDump)

	e.GET("/api/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
