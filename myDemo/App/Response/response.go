package Response

import (
	"github.com/labstack/echo"
	"net/http"
)

type Res struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func SendRes(c echo.Context, code int, message string, data ...interface{}) error {
	return c.JSON(http.StatusOK, Res{
		code,
		message,
		data,
	})
}
