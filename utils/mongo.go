package utils

import (
	"context"
	"github.com/gaochuwuhan/goutils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoEngine(uri string) (*mongo.Client,error){
	clientOptions:=options.Client().ApplyURI(uri)
	// 连接到mongoDB
	mgoCli,err:=mongo.Connect(context.TODO(),clientOptions)
	if err != nil {
		panic("mongodb connection failed")
	}
	// 检查连接
	err=mgoCli.Ping(context.TODO(),nil)
	if err != nil {
		logger.Log.Sugar().Errorf("mongo ping 失败,msg:%s",err.Error())
		return nil,err
	}
	return mgoCli,nil

}


