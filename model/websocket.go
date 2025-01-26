package model

import "github.com/gorilla/websocket"

type StartWebSocketInput struct {
	ReadBufferSize  int
	WriteBufferSize int
	Path            string
	NewConnection   func(*WebSocketConnection)
	ReadFunc        func(WebSocketReadMessage)
}

type WebSocketManager struct {
	path          string
	connections   map[string]*WebSocketConnection
	newConnection func(*WebSocketConnection)
	readFunc      func(WebSocketReadMessage)
}

type WebSocketConnection struct {
	*websocket.Conn
	Context *Context
}

type WebSocketReadMessage struct {
	Type    int
	Data    []byte
	Error   *error
	Context *Context
}

type WebSocketSendInput struct {
	Key  string
	Type int
	Data []byte
}

type WebSocketSetNewConnectionInput struct {
	Key        string
	Connection *WebSocketConnection
}
