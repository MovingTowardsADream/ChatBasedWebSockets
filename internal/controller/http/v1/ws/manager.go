package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var websoketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Manager struct {
	clients ClientsList
	mu      sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{clients: make(ClientsList)}
}

func (m *Manager) ServeWs(w http.ResponseWriter, r *http.Request) {

	// Open new connection
	conn, err := websoketUpgrader.Upgrade(w, r, nil)
	//defer conn.Close()

	if err != nil {
		return
	}

	// Create and add new client
	client := NewClient(conn, m)
	m.addNewClient(client)

	// Read and write messages for client
	go client.readMessages()
	go client.writeMessages()
}

func (m *Manager) addNewClient(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients[client] = true
}

func (m *Manager) deleteClient(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}
