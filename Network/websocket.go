package Network

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func process(c *websocket.Conn, info []string, dbInfo []string, datas map[string]int) {

	// db 정보를 읽고 변수에 초기화
	dbHost := dbInfo[0]
	dbUser := dbInfo[1]
	dbPassword := dbInfo[2]
	dbDatabase := dbInfo[3]
	dbPort := dbInfo[4]

	// 슬라이스에 db 데이터를 저장한다
	dbSlice := []string{dbHost, dbUser, dbPassword, dbDatabase, dbPort}

	// value 값들을 변수에 초기화
	accessSequence := datas["accessSequence"]
	speed := datas["speed"]
	detectiline := datas["detectiline"]
	direction := datas["direction"]
	carInfo := datas["category"]

	// car slice

	carSlice := []byte{byte(accessSequence), byte(speed), byte(detectiline), byte(direction), byte(carInfo)}

	for i := 0; i <= len(dbSlice)-1; i++ {
		dbData := dbSlice[i]

		data := dbData

		err := c.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("Error", err)
			return
		}

		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(message))

		if i == len(dbSlice)-1 {
			err := c.WriteMessage(websocket.TextMessage, []byte("END"))
			if err != nil {
				log.Println("Error", err)
				return
			}
		}

	}

	for i := 0; i <= len(carSlice)-1; i++ {
		err := c.WriteMessage(websocket.TextMessage, []byte(carSlice))
		if err != nil {
			log.Println("Error", err)
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

func WebsocketHandler(info []string, dbInfo []string, datas map[string]int) {

	// Websocket server host, port 초기화
	host := info[0]
	port := info[1]

	// Websocket port , host
	address := fmt.Sprintf("ws://%s:%s/", host, port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := address
	log.Println(u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go catchSig(interrupt, c)

	process(c, info, dbInfo, datas)
}
