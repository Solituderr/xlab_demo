package MiddleWare

import (
	"github.com/labstack/echo"
	"myDemo/model"
)

var user model.User

// 中间件
//处理登录 没实现

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//默认用户
		user.Name = "user1"
		user.Passwd = "user1"
		user.Id = 1
		//fmt.Println(user)
		return next(c)
	}
}

// AddHeader 加返回头的东西
func AddHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "todoDemo/1.0")
		return next(c)
	}
}
