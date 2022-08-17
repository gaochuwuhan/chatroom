package dto

import (
	"chatroom/config"
	"chatroom/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

var mgoCli *mongo.Client


func GetMgoCli(){
	if mgoCli==nil {
		client,_:=utils.InitMongoEngine(config.VP.Mongo.URI)
		if client != nil{
			mgoCli=client
		}
	}
	if mgoCli==nil{
		panic("InitMongoEngine failed")
	}
}