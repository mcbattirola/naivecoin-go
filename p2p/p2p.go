package p2p

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

var sockets []*websocket.Conn

type P2PMessage struct {
	Data string `json:"data"`
}

func ConnectToPeers(newPeer string) error {
	url := url.URL{
		Scheme: "ws",
		Host:   newPeer,
		Path:   "/",
	}

	connection, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		fmt.Println("dial:", err)
		return err
	}

	initConnection(connection)

	return nil
}

func GetSockets() []*websocket.Conn {
	return sockets
}

func initConnection(connection *websocket.Conn) {
	sockets = append(sockets, connection)

	connection.SetCloseHandler(func(code int, text string) error {
		closeConnection(connection)
		return nil
	})

	go func() {
		for {
			_, message, err := connection.ReadMessage()
			if err != nil {
				closeConnection(connection)
				break
			}

			var newMessage P2PMessage
			if err := json.Unmarshal(message, newMessage); err != nil {
				closeConnection(connection)
				break
			}

			fmt.Println("msg received: ", newMessage.Data)

		}

	}()
}

func closeConnection(socketToClose *websocket.Conn) {
	var newSockets []*websocket.Conn

	for _, socket := range sockets {
		if socket != socketToClose {
			newSockets = append(newSockets, socket)
		}
	}

	socketToClose.Close()

	sockets = newSockets
}
