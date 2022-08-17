package main

import (
	"chatroom/config"
	"chatroom/internal/dto"
	"github.com/gaochuwuhan/goutils/initialize"
)

func main(){
	initialize.InitService()
	config.Initialize()
	dto.GetMgoCli()
	dto.InsertExample()

}

