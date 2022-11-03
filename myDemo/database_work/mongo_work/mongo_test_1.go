package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name string `bson:"name"`
	Id   uint   `bson:"id"`
}

func main() {
	//创建客户端服务
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	//连接mongodb
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected successfully")

	//获取数据库和表
	colStu := client.Database("my_data").Collection("Students")
	//单数据插入
	//stu1 := Student{"cjw", 21051407}
	//_, err = colStu.InsertOne(context.TODO(), stu1)

	//多数据插入
	//name := []string{"cjw", "cj", "c"}
	//id := []int{1, 2, 3}
	//var allData []interface{}
	//for i := 0; i < len(name); i++ {
	//	data := bson.D{{"name", name[i]}, {"id", id[i]}}
	//	allData = append(allData, data)
	//}
	//fmt.Println(allData)
	//
	//_, err = colStu.InsertMany(context.TODO(), allData)
	//if err != nil {
	//	panic(err)
	//}

	//查找

	//cur, _ := colStu.Find(context.TODO(), bson.D{{"name", "cjw"}})
	//for cur.Next(context.TODO()) {
	//	var elem Student
	//	err = cur.Decode(&elem)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("%v\n", elem)
	//}
	//err = cur.Close(context.TODO())

	//修改

	//filter := bson.D{{"name", "cj"}}
	////update := bson.D{{"$set", bson.D{{"id", "123456"}}}}
	//update1 := bson.D{{"$inc", bson.D{{"id", 20}}}}
	////_, err = colStu.UpdateMany(context.TODO(), filter, update)
	//_, err = colStu.UpdateMany(context.TODO(), filter, update1)
	//删除

	filter := bson.D{{"name", "cj"}}
	_, err = colStu.DeleteMany(context.TODO(), filter)

}
