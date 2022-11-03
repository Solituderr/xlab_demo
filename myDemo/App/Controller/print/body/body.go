package body

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"io"
	"myDemo/App/Response"
	"myDemo/model"
	"myDemo/services"
	"net/http"
	"strconv"
)

//全是handlers

// DelPost handler 接受json,post请求  测试能否正常接受 暂时没啥用
func DelPost(c echo.Context) error {
	var todo model.Todo
	defer c.Request().Body.Close()
	data, err := io.ReadAll(c.Request().Body)
	if err == nil {
		logrus.Error("read error!")
	}
	//bind只能用于表单形式
	//err := c.Bind(&user)
	//if err != nil {
	//	panic(err)
	//}
	if err = json.Unmarshal(data, &todo); err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": todo.UserId,
			"id":      todo.Id,
			"title":   todo.Title,
			"content": todo.Content,
		}).Error("json unmarshal error!")
		return c.String(http.StatusInternalServerError, "json unmarshal error!")
	}
	fmt.Println(todo)
	return Response.SendRes(c, http.StatusOK, "post is ok", todo)
}

// CreateTodo 创建todo
func CreateTodo(c echo.Context) error {
	todo := new(model.Todo)
	defer c.Request().Body.Close()
	data, err := io.ReadAll(c.Request().Body)
	if err == nil {
		logrus.Error("read error!")
	}
	if err := json.Unmarshal(data, &todo); err != nil {
		logrus.Error("json unmarshal error!")
	}
	services.TodoAdd(todo)
	return Response.SendRes(c, http.StatusOK, "post is ok", "add is ok")
}

// UpdateTodo 传json删除
func UpdateTodo(c echo.Context) error {
	todo := new(model.Todo)
	defer c.Request().Body.Close()
	data, err := io.ReadAll(c.Request().Body)
	if err == nil {
		logrus.Error("read error!")
	}
	if err := json.Unmarshal(data, &todo); err != nil {
		logrus.Error("json unmarshal error!")
	}
	services.TodoUpdate(todo)
	return Response.SendRes(c, http.StatusOK, "post is ok", "update is ok")
}

// DeleteTodo get请求传id删除
func DeleteTodo(c echo.Context) error {
	data := c.QueryParam("id")
	i, _ := strconv.Atoi(data)
	num := uint(i)
	services.TodoDelete(num)
	return Response.SendRes(c, http.StatusOK, "post is ok", "delete is ok")
}

// SelectTodo get请求传id搜索
func SelectTodo(c echo.Context) error {
	data := c.QueryParam("id")
	i, _ := strconv.Atoi(data)
	num := uint(i)
	todo1 := services.TodoSelect(num)
	return Response.SendRes(c, http.StatusOK, "post is ok", todo1)
}
