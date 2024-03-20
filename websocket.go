package inutil

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	connections map[string]*WebSocketConnection
}

type WebSocketConnection struct {
	*websocket.Conn
	Context *Context
}

type WebSocketHandlerInput struct {
	NewConnection func(*WebSocketConnection)
	ReadFunc      func(WebSocketReadMessage)
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebsocketHandler(input WebSocketHandlerInput) HandlerFunc {
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
				logInternal("New message arrived")
				input.ReadFunc(WebSocketReadMessage{
					Type:    messageType,
					Data:    data,
					Error:   &err,
					Context: c,
				})
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
