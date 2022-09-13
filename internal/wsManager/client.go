package wsManager

import (
	"chatroom/internal/model"
	"chatroom/internal/pb"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Name string `json:"name"`
	Conn *websocket.Conn
	Send chan []byte

}


func(c *Client) Read(){

	//持续读取ws conn接收到的数据
	for {
		_,msg,err:=c.Conn.ReadMessage()
		if err!=nil{
			log.Fatalf("读取客户端数据 %s错误:%s", c.Name, err)
			return
		}
		c.ProcessData(msg)

	}
}

func(c *Client) ProcessData(msg []byte){
	//将二进制data解析为结构体
	mess:=new(pb.Message)
	proto.Unmarshal(msg,mess)
	if mess.Type != model.HEAT_BEAT {
		ClManager.Broadcast<-msg
	}
}

func(c *Client) Write(){
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send { // 这里之前肯定要有消息如send成员的操作
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}
