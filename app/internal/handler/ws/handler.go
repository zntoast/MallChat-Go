package ws

import (
	"encoding/json"
	"net/http"
	"time"

	"mallchat-go/app/internal/common"
	"mallchat-go/app/internal/svc"
	wsTypes "mallchat-go/app/internal/types/ws"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	pingPeriod = 30 * time.Second
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源
	},
}

func WebSocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 验证用户身份
		userId, ok := common.GetUserIDFromContext(r.Context())
		if !ok {
			http.Error(w, "未授权", http.StatusUnauthorized)
			return
		}

		// 升级HTTP连接为WebSocket连接
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Errorf("WebSocket升级失败: %v", err)
			return
		}

		client := &wsTypes.Client{
			UserId: userId,
			Conn:   conn,
			Send:   make(chan []byte, 256),
		}

		svcCtx.WS.GetRegisterChan() <- client

		go handleClientMessages(client, svcCtx.WS)
	}
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
