package services

import (
	"github.com/sirupsen/logrus"
	"myDemo/model"
)

//增加todo

func TodoAdd(todo *model.Todo) {
	var tmp model.Todo
	tmp = *todo
	err := model.DB.Debug().Create(&tmp).Error
	if err != nil {
		logrus.Error("sql adding error")
	}
}

//

func TodoUpdate(todo *model.Todo) {
	var tmp model.Todo
	tmp = *todo
	err := model.DB.Debug().Updates(&tmp).Error
	if err != nil {
		logrus.Error("sql adding error")
	}
}

//根据id删除todo

func TodoDelete(id uint) {
	tmp := new(model.Todo)
	tmp.Id = id
	err := model.DB.Debug().Delete(&tmp).Error
	if err != nil {
		logrus.Error("sql adding error")
	}
}

//搜索todo

func TodoSelect(id uint) *model.Todo {
	tmp := new(model.Todo)
	tmp.Id = id
	err := model.DB.Debug().Find(&tmp).Error
	if err != nil {
		logrus.Error("sql adding error")
	}
	return tmp
}
