package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func process(c *websocket.Conn) {
	reader := []byte{1, 2, 3, 4, 5}
	for {
		fmt.Printf(" ,")
		data := reader

		err := c.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("test", err)
			return
		}

		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(message)
	}
}

func catchSig(ch chan os.Signal, c *websocket.Conn) {

	<-ch
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println(err)
	}
	return
}

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := "ws://localhost:8000/"
	log.Println(u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go catchSig(interrupt, c)

	process(c)
}
