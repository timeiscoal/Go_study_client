package Network

import (
	"fmt"
	"net"
)

func CheckServer(m string) {
	fmt.Println("Check Server")

	client, err := net.Dial("tcp", "127.0.0.1:2023")

	defer client.Close()

	method := fmt.Sprintf("%s", m)

	_, err = client.Write([]byte(method))
	if err != nil {
		fmt.Println(err)

	}

}
