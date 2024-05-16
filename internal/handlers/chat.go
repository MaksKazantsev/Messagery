package handlers

import (
	"encoding/json"
	"github.com/MaksKazantsev/Chattery/internal/models"
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Chat struct {
	mu sync.RWMutex

	conns map[string]*websocket.Conn
	users map[string]*Member
}

type Member struct {
	ID string
}

func NewMemeber() *Member {
	return &Member{
		ID: uuid.New().String(),
	}
}
func NewChat() *Chat {
	return &Chat{
		conns: make(map[string]*websocket.Conn),
		users: make(map[string]*Member),
	}
}

func (ch *Chat) Join(c *websocket.Conn) {
	member := NewMemeber()
	ch.users[member.ID] = member

	ch.mu.Lock()
	ch.conns[member.ID] = c
	ch.mu.Unlock()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			return
		}
		var message models.Message
		if err = json.Unmarshal(msg, &message); err != nil {
			return
		}
		message.SendAt = time.Now()
		message.SenderID = member.ID

		ch.mu.RLock()
		ch.conns[member.ID].WriteJSON(message)
		for _, v := range ch.users {
			ch.conns[v.ID].WriteJSON(message)
		}
		ch.mu.RUnlock()
	}

}
