package App

import (
	"myDemo/App/Controller/pong"
	"myDemo/App/Controller/print/body"
	"myDemo/App/MiddleWare"
)

// AddRoutes handlers
func AddRoutes() {
	api := e.Group("api")
	api.Use(MiddleWare.Auth)
	api.Use(MiddleWare.AddHeader)
	api.GET("/pong", pong.GetRes)
	//api.GET("/getUser", query.DelGet)
	//api.POST("/postUser", body.DelPost)
	api.POST("/service/add", body.CreateTodo)
	api.POST("/service/update", body.UpdateTodo)
	api.GET("/service/delete", body.DeleteTodo)
	api.GET("/service/select", body.SelectTodo)
}
