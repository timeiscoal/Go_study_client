package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func process(c *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(" ,")
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}

		data = strings.TrimSpace(data)

		err = c.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println(err)
			return
		}

		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(message))
	}
}
