package main

import (
	"chatroom/config"
	"chatroom/internal/dao"
	"github.com/gaochuwuhan/goutils/initialize"
)

func main(){
	initialize.InitService()
	config.Initialize()
	dao.GetMgoCli()
	//dao.InsertExample()
	dao.FindOneExample()
	//dao.InitMongoData()

}

