package router

import (
	"chatroom/internal/wsManager"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var ws = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func NewVisitorWs(c *gin.Context){
	user:=c.Query("user")
	if user==""{
		return
	}
	conn,err:=ws.Upgrade(c.Writer,c.Request,nil)
	if err!=nil{
		return
	}
	client := &wsManager.Client{
		Name:user,
		Conn:conn,
		Send: make(chan []byte),
	}
	wsManager.ClManager.Register <- client
	go client.Read()
	go client.Write()


}