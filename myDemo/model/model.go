package model

type User struct {
	Id     uint   `json:"id" form:"id" gorm:"primaryKey"`
	Name   string `json:"name" form:"name"`
	Passwd string `json:"passwd" form:"passwd"`
}

type Todo struct {
	Id      uint   `json:"id" form:"id" gorm:"primaryKey"`
	UserId  uint   `json:"userId" form:"userId"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}
