package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitSql() {
	connectDatabase()
	fmt.Println("connected successfully")
	if !DB.Migrator().HasTable(&User{}) {
		err := DB.Migrator().CreateTable(&User{})
		if err != nil {
			logrus.Error("create table error")
		}
	}
	if !DB.Migrator().HasTable(&Todo{}) {
		err := DB.Migrator().CreateTable(&Todo{})
		if err != nil {
			logrus.Error("create table error")
		}
	}
	//fmt.Println("okkkkkkkk!")
}

func connectDatabase() {
	//path, _ := os.Getwd()
	//viper.SetConfigFile("./config.yml")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./model")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("viper read error")
	}
	loginInfo := viper.GetStringMapString("User")
	fmt.Println(loginInfo["username"])
	dbArgs := loginInfo["username"] + ":" + loginInfo["passwd"] +
		"@tcp(localhost:3306)/" + loginInfo["dbname"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	//dns := "root:caijunwu123@tcp(localhost:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dbArgs)
	var err error
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Error("mysql connected error")
	}

}
