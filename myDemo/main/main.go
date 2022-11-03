package main

import (
	"github.com/sirupsen/logrus"
	"myDemo/App"
	"myDemo/model"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	model.InitSql()
	App.InitWebServer()
	App.StartServer()
}
