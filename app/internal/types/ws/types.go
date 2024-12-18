package ws

import "github.com/gorilla/websocket"

type Message struct {
	Type      int32  `json:"type"`
	SenderId  int64  `json:"senderId"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

type Client struct {
	UserId int64
	Conn   *websocket.Conn
	Send   chan []byte
}

type Manager interface {
	Start()
	SendToUser(userId int64, message *Message)
	Broadcast(message *Message)
	BroadcastToGroup(groupId int64, message *Message)
	GetRegisterChan() chan<- *Client
	GetUnregisterChan() chan<- *Client
}
