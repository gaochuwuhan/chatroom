package dto

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type TestDoc struct{
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Level int `bson:"level"`
}

func Initcoll() *mongo.Collection{

	//2.选择数据库 my_db
	db := mgoCli.Database("chatroom") //即使client里的mongo uri有db，这里也要选择一下database
	//3.选择表 my_collection
	collection := db.Collection("example")
	return collection

}

func InsertExample(){
	coll:=Initcoll()
	t:=&TestDoc{
		Id: "1",
		Name: "testInsert",
		Level: 0,
	}
	objId, err := coll.InsertOne(context.TODO(), t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("insert example,get objId:",objId)
}
