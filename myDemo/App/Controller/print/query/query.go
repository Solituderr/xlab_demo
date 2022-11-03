package query

import (
	"fmt"
	"github.com/labstack/echo"
	"myDemo/App/Response"
	"net/http"
)

// DelGet handler 接受query 学习时用的，最后案例没用
func DelGet(c echo.Context) error {
	name := c.QueryParam("name")
	fmt.Println(name)
	return Response.SendRes(c, http.StatusOK, "query is ok", "123")
}
