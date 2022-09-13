package dao

import (
	"chatroom/internal/model"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func InitMongoData(){
	InitUsers()
}

//只执行一次，或者插入的时候检查邮箱名是否存在不存在则插入
func InitUsers(){
	coll := mgoCli.Database("chatroom").Collection("user")
	users:=make([]interface{},0)
	users = append(users,model.User{
		Id: primitive.NewObjectID(),
		Uuid: uuid.New().String(),
		Username: "u_1",
		Email: "u_1@ws.com",
		CreateAt: time.Now(),

	})

	users = append(users,model.User{
		Id: primitive.NewObjectID(),
		Uuid: uuid.New().String(),
		Username: "u_2",
		Email: "u_2@ws.com",
		CreateAt: time.Now(),

	})
	ctx:=context.Background()
	result, err := coll.InsertMany(ctx, users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("new user ids :",result)



}
