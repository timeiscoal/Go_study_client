package Network

import (
	"fmt"
	"net"
	"time"
)

func TcpHandler(info []string, dbInfo []string, datas map[string]int) {

	// db 정보를 읽고 변수에 초기화
	dbHost := dbInfo[0]
	dbUser := dbInfo[1]
	dbPassword := dbInfo[2]
	dbDatabase := dbInfo[3]
	dbPort := dbInfo[4]

	// 슬라이스에 db 데이터를 저장한다
	dbSlice := []string{dbHost, dbUser, dbPassword, dbDatabase, dbPort}

	// tcp server host, port 초기화
	host := info[0]
	port := info[1]

	// tcp port , host
	address := fmt.Sprintf("%s:%s", host, port)

	// Dial : 프로토콜, ip주소,포트번호를 설정하여 서버에 연결. TCP프로토콜 ip : 127.0.0.1 : 2023
	client, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	// main 함수가 끝나기 직전에 TCP 연결 닫기
	defer client.Close()

	// value 값들을 변수에 초기화
	accessSequence := datas["accessSequence"]
	speed := datas["speed"]
	detectiline := datas["detectiline"]
	direction := datas["direction"]
	carInfo := datas["category"]

	for i := 0; i <= len(dbSlice)-1; i++ {
		//ini에서 읽은 데이터를 서버로 전송한다.
		data := dbSlice[i]
		_, err := client.Write([]byte(data))
		if err != nil {
			fmt.Println(err)
			return
		}
		// data send time
		time.Sleep(1 * time.Second)
		// 데이터를 전부 다 보내면 , END 메세지를 서버에 보낸다.
		if i == len(dbSlice)-1 {
			message := "END"
			_, err := client.Write([]byte(message))
			if err != nil {
				fmt.Println(err)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}
	count := 0

	for {
		// 검지된 차량 데이터를 서버로 전송
		data := []byte{byte(accessSequence), byte(speed), byte(detectiline), byte(direction), byte(carInfo)}

		fmt.Println("전송 데이터", data)
		_, err := client.Write([]byte(data))
		if err != nil {
			fmt.Println(err)
			return
		}
		// data send time
		time.Sleep(1 * time.Millisecond)

		count++

		if count >= 1 {
			break
		}
	}

}
