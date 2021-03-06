package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ndidplatform/ndid/api/identity"
	"github.com/ndidplatform/ndid/api/rp"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/identity/:ns/:id", identity.GetIdentifier)
	e.POST("/identity", identity.CreateIdentity)

	e.POST("/rp/requests/:ns/:id", rp.CreateRequest)

	e.Logger.Fatal(e.Start(":8000"))
}
