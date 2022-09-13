package model

import (
	"time"
)

//功能：将srcIp+端口 的消息发送给 dstIp
type Message struct {
	ID          int32                 `bson:"_id"`
	CreatedAt   time.Time         	  `bson:"createAt"`
	UpdatedAt   time.Time             `bson:"updatedAt"`
	DeletedAt   time.Time  		   	  `bson:"deletedAt"`
	FromUserId  int32                 `bson:"fromUserId"`
	ToUserId    int32                 `bson:"toUserId"`
	Content     string                `bson:"content"`
	MessageType int16                 `bson:"messageType"`
	ContentType int16                 `bson:"contentType"`
	Pic         string                `bson:"pic"`
	Url         string                `bson:"url"`
}

const (
	HEAT_BEAT = "heatbeat"
	PONG      = "pong"

	// 消息类型，单聊或者群聊
	MESSAGE_TYPE_USER  = 1
	MESSAGE_TYPE_GROUP = 2

	// 消息内容类型
	TEXT         = 1
	FILE         = 2
	IMAGE        = 3
	AUDIO        = 4
	VIDEO        = 5
	AUDIO_ONLINE = 6
	VIDEO_ONLINE = 7

	// 消息队列类型
	GO_CHANNEL = "gochannel"
	KAFKA      = "kafka"
)
