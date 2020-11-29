package p2p

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

var sockets []*websocket.Conn

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

	// TODO implement web sockets
}
