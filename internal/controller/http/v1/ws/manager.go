package ws

import (
	"errors"
	"fmt"
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

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

type Manager struct {
	clients ClientsList
	mu      sync.RWMutex

	handlers map[string]EventHandler
}

func NewManager() *Manager {
	m := &Manager{
		clients:  make(ClientsList),
		handlers: make(map[string]EventHandler),
	}
	m.setupEventHandlers()
	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessageHandler
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	// Check if Handler is present in Map
	if handler, ok := m.handlers[event.Type]; ok {
		// Launch Handler
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}

func SendMessageHandler(event Event, c *Client) error {
	for client := range c.manager.clients {
		client.egress <- event
	}
	fmt.Println(event)
	return nil
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
