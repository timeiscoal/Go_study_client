package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	k := "127.0.0.1"
	j := "8000"

	b := fmt.Sprintf("%s:%s", k, j)

	client, err := net.Dial("tcp", b) // TCP 프로토콜, 127.0.0.1:8000 서버에 연결
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

	for {
		i := []uint8{1, 2, 3, 4, 5}
		for {
			s := i
			fmt.Println(i)
			_, err := client.Write([]byte(s)) // 서버로 데이터를 보냄
			if err != nil {
				fmt.Println(err)
				return
			}

			time.Sleep(1 * time.Second)
		}
	}

}
