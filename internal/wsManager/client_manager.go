package wsManager

import (
	"fmt"
	"sync"
)

var (
	ClManager = NewClientManager()
)

// 连接管理
type ClientManager struct {
	Clients     map[string]*Client   // 全部的连接
	ClientsLock    *sync.RWMutex       // 读写锁
	Register    chan *Client       // 连接连接处理
	Unregister  chan *Client       // 断开连接处理程序
	Broadcast   chan []byte        // 广播 向全部成员发送数据
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
		ClientsLock: &sync.RWMutex{},
	}

	return
}

func(c *ClientManager) AddClient(client *Client){
	c.ClientsLock.Lock() //add客户端的时候即不能读也不能写
	defer c.ClientsLock.Unlock()
	c.Clients[client.Name] = client
}

//func(c *ClientManager) Start(){
//	//持续接收消息
//	for {
//		select {
//		case conn := <-c.Register:
//			// 建立连接事件
//			c.EventRegister(conn)
//
//		//case login := <-c.Login:
//		//	// 用户登录
//		//	c.EventLogin(login)
//
//		//case conn := <-c.Unregister:
//		//	// 断开连接事件
//		//	c.EventUnregister(conn)
//
//		case message := <-c.Broadcast:
//			// 广播事件
//			clients := c.GetClients()
//			for conn := range clients {
//				select {
//				case conn.Send <- message:
//				default:
//					close(conn.Send)
//				}
//			}
//		}
//	}
//}

func(c *ClientManager) EventRegister(client *Client){
	//将新链接的client放到Clients这个map中
	c.AddClient(client)
	fmt.Println("EventRegister 用户建立连接", client.Name)
}

