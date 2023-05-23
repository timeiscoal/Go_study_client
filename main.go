package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-resty/resty/v2"
)

// var carInFo = []string{"", "", "car", "miniTruck", "truck", "miniBus", "bus", "Motorcycle"}
// var detectLine = []int{0, 1, 2, 3, 4, 5}
// var direction = []string{"", "straight", "left", "right", "u_turn"}

func testData() {

	// 해당 데이터들은 int8 ,
	// 데이터 범위는 코드 단에서 검증해서 범위 내에서 입력된 값만 받을 수 있게 분기문 작성
	// 접근로 시퀀스 (1~100)
	// 검지된 차선(1~5)
	// 속도 (20~60)

	// DB에 적용되는 시간을 기록.
	// 생성 시간(DB에 insert시간)

	// 테이블을 각각 따로 생성
	car := map[int]string{
		2: "소형차",
		3: "소형 트럭",
		4: "대형 트럭",
		5: "소형 버스",
		6: "대형 버스",
		7: "오토바이",
	}

	direction := map[int]string{
		1: "직진",
		2: "좌회전",
		3: "우회전",
		4: "유턴",
	}

	for key, value := range car {
		fmt.Println(key, value)
	}
	for key, value := range direction {
		fmt.Println(key, value)
	}

}

// 데이터 베이스 테이블
type carCountingInfo struct {
	AccessSequence int8 //접근로 시퀀스
	Speed          int8 //속도
	DetectLine     int8 //검지된 차선
	Direction      int8 //이동방향
	CarInfo        int8 //차량 정보(종류)
	// Created_at     time.Time //데이터 생성시간
}

func main() {
	testdata := createData()
	createData()
	testData()
	carQuery()
	carPost(testdata)
}

func createData() [5]int {
	rand.Seed(time.Now().UnixNano())
	// 접근로 시퀀스 범위 설정
	num := rand.Intn(100)
	// 속도 범위 (속도 20미만은 탐지되지 않음)
	// if speedNum := rand.Intn(81); speedNum < 20 {
	// 	speedNum += 20
	// 	println("속도 : ", speedNum)
	// } else {
	// 	println("속도 : ", speedNum)
	// }

	speedNum := rand.Intn(81)

	// 감지된 차선
	detectiline := rand.Intn(5)
	detectiline += 1

	// map 활용

	// 이동방향
	direction := rand.Intn(4)
	direction += 1

	// 차량(종류)
	carInfo := rand.Intn(6)
	carInfo += 2

	// DB 적용
	//데이터 베이스 적용 시간
	t := time.Now()
	fmt.Println(t.Format(time.DateTime))

	arr := [...]int{num, speedNum, detectiline, direction, carInfo}

	return arr

}

func carQuery() {
	client := resty.New()
	response, err := client.R().EnableTrace().Get("http://localhost:8080/")
	fmt.Println(response.StatusCode(), err)

}

func carPost(arr [5]int) {
	client := resty.New()
	response, err := client.R().SetBody(carCountingInfo{
		CarInfo:        int8(arr[4]),
		DetectLine:     int8(arr[2]),
		Speed:          int8(arr[1]),
		AccessSequence: int8(arr[0]),
		Direction:      int8(arr[3]),
	}).Post("http://localhost:8080/")

	fmt.Println(response.StatusCode(), err)
}
