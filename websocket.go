package inutil

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

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

var upgrader websocket.Upgrader

func (swsi *StartWebSocketInput) startWebSocket() *WebSocketManager {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  swsi.ReadBufferSize,
		WriteBufferSize: swsi.WriteBufferSize,
	}
	if swsi.Path == "" {
		Log("No websocket path, using default /ws")
		swsi.Path = "/ws"
	}
	wsm := &WebSocketManager{
		connections:   map[string]*WebSocketConnection{},
		path:          swsi.Path,
		newConnection: swsi.NewConnection,
		readFunc:      swsi.ReadFunc,
	}
	return wsm
}

func WebsocketHandler() HandlerFunc {
	return func(c *Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			// panic(err)
			log.Printf("%s, error while Upgrading websocket connection\n", err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		go func() {
			for {
				// Read message from client
				logInternal("Waiting new messages from the websocket")
				messageType, data, err := conn.ReadMessage()
				logInternalF("New message arrived type:%v data:%v", messageType, string(data))
				if messageType == 12371823 {
					inutil.WebSocketManager.newConnection(&WebSocketConnection{
						Conn:    conn,
						Context: c,
					})
				} else {
					inutil.WebSocketManager.readFunc(WebSocketReadMessage{
						Type:    messageType,
						Data:    data,
						Error:   &err,
						Context: c,
					})
				}
			}
		}()
	}
}

func (wsm *WebSocketManager) SetNewConnection(input WebSocketSetNewConnectionInput) {
	wsm.connections[input.Key] = input.Connection
}

func (wsm *WebSocketManager) Send(input WebSocketSendInput) error {
	return wsm.connections[input.Key].Conn.WriteMessage(input.Type, input.Data)
}

const (
	DefaultReadBufferSize  = 1024
	DefaultWriteBufferSize = 1024
)
