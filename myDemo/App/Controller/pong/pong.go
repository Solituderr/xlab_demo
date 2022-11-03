package pong

import (
	"github.com/labstack/echo"
	"myDemo/App/Response"
	"net/http"
)

// GetRes handler 检测是否存活 pong
func GetRes(c echo.Context) error {
	return Response.SendRes(c, http.StatusOK, "get is ok", "pong!")
}
