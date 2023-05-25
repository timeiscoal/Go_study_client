package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// udp addr 생성
	remoteAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6000")

	if err != nil {
		log.Fatal(err)
	}

	// 생성한 remoteAddr로 dial ,ip주소, 포트번호
	conn, err := net.DialUDP("udp", nil, remoteAddr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conn)

	// 서버에게 데이터 송신
	// 얼마나 보내는지 체크
	count := 0
	for {
		msg := []uint8{1, 2, 3, 4, 5}
		_, err = conn.Write(msg)
		if err != nil {
			log.Fatal(err)
		}
		count++
		if count > 10 {
			break
		}
	}

	// 서버에게 받은 데이터
	buf := make([]byte, 30)
	n, addr, err := conn.ReadFromUDP(buf)

	fmt.Println(addr)
	fmt.Println(string(buf[:n]))
}
