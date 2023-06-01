package Network

import (
	"fmt"
	"golang/randomData"

	"github.com/go-resty/resty/v2"
)

type DbInfo struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

type Data struct {
	randomData.CarInfo
	DbInfo
}

func RestHandler(info []string, dbInfo []string, data map[string]int) {

	var dB Data
	// db 정보를 읽고 변수에 초기화
	dB.Host = dbInfo[0]
	dB.User = dbInfo[1]
	dB.Password = dbInfo[2]
	dB.Database = dbInfo[3]
	dB.Port = dbInfo[4]

	fmt.Println(dB)

	// 슬라이스에 db 데이터를 저장한다
	//dbSlice := []string{dbHost, dbUser, dbPassword, dbDatabase, dbPort}

	// rest server host, port 초기화
	host := info[0]
	port := info[1]

	// rest Api 주소
	restApi := "post"
	address := fmt.Sprintf("http://%s:%s/%s", host, port, restApi)

	// 자동차 정보 타입 선언 및 초기화

	var c Data
	c.AccessSequence = data["accessSequence"]
	c.Category = data["category"]
	c.DetectLine = data["detectiline"]
	c.Direction = data["direction"]
	c.Speed = data["speed"]

	// post 요청
	client := resty.New()

	// map을 활용하여 데이터를 전송
	_, err := client.R().SetBody(map[string]interface{}{
		"accessSequence": c.AccessSequence,
		"category":       c.Category,
		"detectline":     c.DetectLine,
		"direction":      c.Direction,
		"speed":          c.Speed,
		"Host":           dB.Host,
		"User":           dB.User,
		"Password":       dB.Password,
		"Database":       dB.Database,
		"Port":           dB.Port}).
		Post(address)
	if err != nil {
		panic(err)
	}

}
