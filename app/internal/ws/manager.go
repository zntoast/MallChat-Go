package ws

import (
	"context"
	"mallchat-go/app/internal/model"
	"sync"
	"time"

	wsTypes "mallchat-go/app/internal/types/ws"

	"github.com/zeromicro/go-zero/core/logx"
)

type Manager struct {
	clients    map[int64]*wsTypes.Client
	broadcast  chan *wsTypes.Message
	register   chan *wsTypes.Client
	unregister chan *wsTypes.Client
	mutex      sync.RWMutex
	retryQueue chan *RetryMessage
	models     *model.Models
}

type RetryMessage struct {
	Message    *wsTypes.Message
	RetryCount int
	LastRetry  time.Time
	TargetUser int64
}

func NewManager(models *model.Models) wsTypes.Manager {
	m := &Manager{
		clients:    make(map[int64]*wsTypes.Client),
		broadcast:  make(chan *wsTypes.Message),
		register:   make(chan *wsTypes.Client),
		unregister: make(chan *wsTypes.Client),
		retryQueue: make(chan *RetryMessage),
		models:     models,
	}
	go m.Start()
	return m
}

func (m *Manager) GetRegisterChan() chan<- *wsTypes.Client {
	return m.register
}

func (m *Manager) GetUnregisterChan() chan<- *wsTypes.Client {
	return m.unregister
}

func (m *Manager) Start() {
	for {
		select {
		case client := <-m.register:
			m.mutex.Lock()
			m.clients[client.UserId] = client
			m.mutex.Unlock()
			logx.Infof("用户[%d]已连接", client.UserId)

		case client := <-m.unregister:
			m.mutex.Lock()
			if _, ok := m.clients[client.UserId]; ok {
				delete(m.clients, client.UserId)
				close(client.Send)
			}
			m.mutex.Unlock()
			logx.Infof("用户[%d]已断开", client.UserId)

		case message := <-m.broadcast:
			m.mutex.RLock()
			for userId, client := range m.clients {
				select {
				case client.Send <- []byte(message.Content):
				default:
					close(client.Send)
					delete(m.clients, userId)
				}
			}
			m.mutex.RUnlock()

		case <-context.Background().Done():
			return
		}
	}
}

func (m *Manager) SendToUser(userId int64, message *wsTypes.Message) {
	m.mutex.RLock()
	if client, ok := m.clients[userId]; ok {
		select {
		case client.Send <- []byte(message.Content):
			// 发送成功
		default:
			// 发送失败,加入重试队列
			m.retryQueue <- &RetryMessage{
				Message:    message,
				RetryCount: 0,
				LastRetry:  time.Now(),
				TargetUser: userId,
			}
		}
	} else {
		// 用户离线,存储离线消息
		m.storeOfflineMessage(userId, message)
	}
	m.mutex.RUnlock()
}

func (m *Manager) retryWorker() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case retry := <-m.retryQueue:
			if retry.RetryCount >= 3 {
				// 重试次数达到上限,存储为离线消息
				m.storeOfflineMessage(retry.TargetUser, retry.Message)
				continue
			}

			// 重试发送
			m.mutex.RLock()
			if client, ok := m.clients[retry.TargetUser]; ok {
				select {
				case client.Send <- []byte(retry.Message.Content):
					// 发送成功
				default:
					// 发送失败,重新加入队列
					retry.RetryCount++
					retry.LastRetry = time.Now()
					m.retryQueue <- retry
				}
			}
			m.mutex.RUnlock()

		case <-ticker.C:
			// 定期检查重试队列
		}
	}
}

func (m *Manager) Broadcast(message *wsTypes.Message) {
	m.broadcast <- message
}

func (m *Manager) BroadcastToGroup(groupId int64, message *wsTypes.Message) {
	// 获取群成员
	members, err := m.models.GroupModel.GetMembers(groupId)
	if err != nil {
		logx.Errorf("获取群成员失败: %v", err)
		return
	}

	// 给每个在线成员发送消息
	for _, userId := range members {
		m.SendToUser(userId, message)
	}
}

func (m *Manager) storeOfflineMessage(userId int64, message *wsTypes.Message) {
	offlineMsg := &model.OfflineMessage{
		UserId:     userId,
		SenderId:   message.SenderId,
		Content:    message.Content,
		Type:       int64(message.Type),
		CreateTime: message.Timestamp,
	}

	if _, err := m.models.OfflineMessageModel.Insert(context.Background(), offlineMsg); err != nil {
		logx.Errorf("存储离线消息失败: %v", err)
	}
}
