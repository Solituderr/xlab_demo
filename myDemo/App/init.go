package App

import (
	"github.com/labstack/echo"
)

var e *echo.Echo

func InitWebServer() {
	e = echo.New()
	e.HideBanner = true
	AddRoutes()
	//fmt.Println()
	//StartServer()
}

func StartServer() {
	e.Logger.Fatal(e.Start("127.0.0.1:1232"))
}
