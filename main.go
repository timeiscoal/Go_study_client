package main

import (
	"fmt"
	"golang/Network"
	"golang/database"
	"golang/randomData"
	"strconv"
	"time"

	"golang/ini_reader"
)

// client
func main() {

	// 0. random data

	carData := randomData.DataHandler()

	// 1. Read ini : ini에 읽은 데이터들을 ini에 초기화
	ini := ini_reader.IniReader()

	// DB Info
	dbHost := ini["dbHost"]
	dbUser := ini["dbUser"]
	dbPassword := ini["dbPassword"]
	dbDatabase := ini["dbDatabase"]
	dbPort := ini["dbPort"]

	dbInfo := []string{dbHost, dbUser, dbPassword, dbDatabase, dbPort}

	// Network Info
	netMethod := ini["netMethod"]
	netHost := ini["netHost"]
	netPort := ini["netPort"]

	netInfo := []string{netHost, netPort}

	// 어떤 통신으로 진행 할지 결정해서 서버에게 우선적으로 알려줌
	Network.CheckServer(netMethod)
	// 서버가 열릴 때 까지 잠시 대기.
	time.Sleep(1 * time.Second)

	// 2. network connection
	switch n, _ := strconv.Atoi(netMethod); n {
	case 0:

		fmt.Println("none")
	case 1:
		fmt.Println(netMethod)
		Network.TcpHandler(netInfo, dbInfo, carData)
	case 2:
		Network.UdpHandler(netInfo, dbInfo, carData)
	case 3:
		Network.WebsocketHandler(netInfo, dbInfo, carData)
	case 4:
		Network.RestHandler(netInfo, dbInfo, carData)

	}

	// 3. send message

	// 4. database connection

	database.PostgresHandler(dbInfo)

	// 5. create random data

}
