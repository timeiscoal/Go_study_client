package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// 클라 TCP
func main() {

	client, err := net.Dial("tcp", "127.0.0.1:8000") // tcp 프로토콜 서버 연결
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close() // main 함수가 끝나기전에 TCP 연결 닫기

	go func(c net.Conn) {
		data := make([]byte, 4096)

		for {
			n, err := c.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(data[:n]))
		}
	}(client)

	go func(c net.Conn) {
		i := 0
		for {
			s := "Hello" + strconv.Itoa(i)

			_, err := c.Write([]byte(s))
			if err != nil {
				fmt.Println(err)
				return
			}
			i++
			time.Sleep(1 * time.Second)
		}
	}(client)
	fmt.Scanln()
}

// https://pyrasis.com/book/GoForTheReallyImpatient/Unit56/02

// https://kamang-it.tistory.com/entry/golanggohttpWebSocketwebsocket%EC%82%AC%EC%9A%A9%ED%95%B4%EC%84%9C-server-client-%EB%B8%8C%EB%9D%BC%EC%9A%B0%EC%A0%80%EB%A5%BC-%ED%86%B5%ED%95%98%EC%97%AC-%ED%86%B5%EC%8B%A0%ED%95%98%EA%B8%B0-%EB%8B%A8%EC%9D%BC-%ED%86%B5%EC%8B%A0

//https://angehende-ingenieur.tistory.com/219?category=1067950

//https://www.bearpooh.com/122

// https://dejavuqa.tistory.com/16
//https://psychoria.tistory.com/783
//https://kicarussays.tistory.com/53
//http://www.gisdeveloper.co.kr/?p=2456
//https://brownbears.tistory.com/186
//https://dksshddl.tistory.com/entry/Go-web-programming-postgresSQL%EC%97%B0%EA%B2%B0-CRUD%EC%82%AC%EC%9A%A9%ED%95%B4%EB%B3%B4%EA%B8%B0
//https://tzara.tistory.com/168
//https://velog.io/@chappi/golang-%EB%B0%B1%EC%97%94%EB%93%9C-%EA%B0%9C%EB%B0%9C%EC%9D%84-%ED%95%B4%EB%B3%B4%EC%9E%90
