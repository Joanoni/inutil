package inutil

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketManager struct {
	Sessions map[string]string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebsocketHandler(c *Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// panic(err)
		log.Printf("%s, error while Upgrading websocket connection\n", err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for {
		// Read message from client
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			// panic(err)
			log.Printf("%s, error while reading message\n", err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			break
		}

		// Echo message back to client
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			// panic(err)
			log.Printf("%s, error while writing message\n", err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			break
		}
	}
}
