package randomData

import (
	"math/rand"
	"time"
)

// 자동차 정보 구조체 선언
type CarInfo struct {
	AccessSequence int //접근로 시퀀스
	Speed          int //속도
	DetectLine     int //검지된 차선
	Direction      int //이동방향
	Category       int //차량 정보(종류)
	// Created_at     time.Time //데이터 생성시간
}

func DataHandler() map[string]int {

	data := make(map[string]int)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	// 접근로 시퀀스
	data["accessSequence"] = rand.Intn(81) + 20

	// 속도
	data["speed"] = rand.Intn(31) + 30

	/// 감지된 차선
	data["detectiline"] = rand.Intn(5) + 1

	// 이동방향
	data["direction"] = rand.Intn(4) + 1

	// 차량(종류)
	data["category"] = rand.Intn(6) + 2

	return data

}
