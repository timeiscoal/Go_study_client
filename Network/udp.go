package Network

import (
	"fmt"
	"log"
	"net"
	"time"
)

func UdpHandler(info []string, dbInfo []string, datas map[string]int) {

	// db 정보를 읽고 변수에 초기화
	dbHost := dbInfo[0]
	dbUser := dbInfo[1]
	dbPassword := dbInfo[2]
	dbDatabase := dbInfo[3]
	dbPort := dbInfo[4]

	// 슬라이스에 db 데이터를 저장한다
	dbSlice := []string{dbHost, dbUser, dbPassword, dbDatabase, dbPort}

	// UDP server host, port 초기화
	host := info[0]
	port := info[1]

	// UDP port , host
	address := fmt.Sprintf("%s:%s", host, port)

	fmt.Println("udp access")
	// ResolveUDPAddr() 을 이용하여 *UDPAddr 구조체를 얻는다.
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}

	// carSlice 선언
	var carSlice []byte
	// value 값들을 변수에 초기화
	accessSequence := datas["accessSequence"]
	speed := datas["speed"]
	detectiline := datas["detectiline"]
	direction := datas["direction"]
	carInfo := datas["category"]

	carSlice = append(carSlice, byte(accessSequence), byte(speed), byte(detectiline), byte(direction), byte(carInfo))

	// DialUDP 는 반환된 소켓(혹은 connection)이 지정된 주소로만 전송/수신이 가능하도록 강제하는 역할.
	conn, err := net.DialUDP("udp", nil, addr)

	fmt.Println(carSlice)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i <= len(dbSlice)-1; i++ {
		// 서버로 보낼 데이터베이스 정보
		data := dbSlice[i]
		msg := []byte(data)
		_, err = conn.Write(msg)
		if err != nil {
			log.Fatal(err)
		}
		if i == len(dbSlice)-1 {
			msg := "END"
			time.Sleep(1 * time.Second)
			_, err = conn.Write([]byte(msg))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Hello")
			break
		}
	}
	for i := 0; i <= len(carSlice)-1; i++ {
		data := carSlice[i]
		var msg []byte
		msg = append(msg, data)
		_, err = conn.Write(msg)
		if err != nil {
			log.Fatal(err)
		}

	}

	buf := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("-----------", addr)
	fmt.Println(string(buf[:n]))
}
