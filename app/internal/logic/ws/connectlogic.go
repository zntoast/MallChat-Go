package ws

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"mallchat-go/app/internal/common"
	"mallchat-go/app/internal/svc"
	wsTypes "mallchat-go/app/internal/types/ws"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type ConnectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewConnectLogic(svcCtx *svc.ServiceContext) *ConnectLogic {
	return &ConnectLogic{
		Logger: logx.WithContext(context.Background()),
		ctx:    context.Background(),
		svcCtx: svcCtx,
	}
}

func (l *ConnectLogic) Connect(w http.ResponseWriter, r *http.Request) {
	// 验证用户身份
	userId, ok := common.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "未授权", http.StatusUnauthorized)
		return
	}

	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		l.Error("WebSocket升级失败", logx.Field("error", err))
		return
	}

	// 创建客户端
	client := &wsTypes.Client{
		UserId: userId,
		Conn:   conn,
		Send:   make(chan []byte, 256),
	}

	// 注册客户端
	l.svcCtx.WS.GetRegisterChan() <- client

	// 启动读写goroutine
	go handleClientMessages(client, l.svcCtx.WS)
}

func handleClientMessages(client *wsTypes.Client, manager wsTypes.Manager) {
	// Start write pump
	go func() {
		ticker := time.NewTicker(pingPeriod)
		defer func() {
			ticker.Stop()
			client.Conn.Close()
		}()

		for {
			select {
			case message, ok := <-client.Send:
				if !ok {
					client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}
				client.Conn.WriteMessage(websocket.TextMessage, message)
			case <-ticker.C:
				if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
			}
		}
	}()

	// Read pump
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logx.Errorf("读取WebSocket消息错误: %v", err)
			}
			break
		}

		var msg wsTypes.Message
		if err := json.Unmarshal(message, &msg); err != nil {
			logx.Errorf("解析WebSocket消息错误: %v", err)
			continue
		}

		msg.SenderId = client.UserId
		msg.Timestamp = time.Now().Unix()
		manager.Broadcast(&msg)
	}

	manager.GetUnregisterChan() <- client
}

const pingPeriod = 30 * time.Second
