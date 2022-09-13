package service

import (
	"chatroom/internal/pb"
)

type Message struct{}

func (m *Message) SaveMessage(msg *pb.Message) {
	//db:=dao.GetDB()
	////user coll中查找要发送的用户，不存在则直接return
	//coll:=db.Database("chatroom").Collection("user")

}
