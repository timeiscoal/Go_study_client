package ini_reader

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// iniInfo 별칭 생성
type iniInfo map[string]string

// ini 파일 읽기
func IniReader() iniInfo {

	fmt.Println("Reading")

	// 경로에 위치한 ini 폴더 읽기
	cfg, err := ini.Load("./static/laon.ini")
	if err != nil {
		fmt.Println(err)
	}

	// map 초기화
	m := make(iniInfo)

	// 읽은 정보들을 map에 key,value로 저장
	m["dbHost"] = cfg.Section("DB").Key("Host").String()
	m["dbUser"] = cfg.Section("DB").Key("User").String()
	m["dbPassword"] = cfg.Section("DB").Key("Password").String()
	m["dbDatabase"] = cfg.Section("DB").Key("Database").String()
	m["dbPort"] = cfg.Section("DB").Key("Port").String()
	m["netMethod"] = cfg.Section("Network").Key("Method").String()
	m["netHost"] = cfg.Section("Network").Key("Host").String()
	m["netPort"] = cfg.Section("Network").Key("Port").String()
	fmt.Println(m)
	return m

	// slice : 읽어야할 데이터가 만약 많다면 슬라이스가 조금 더 유리 하지 않을까?

	// dbHost := cfg.Section("DB").Key("Host").String()
	// dbUser := cfg.Section("DB").Key("User").String()
	// dbPassword := cfg.Section("DB").Key("Password").String()
	// dbDatabase := cfg.Section("DB").Key("Database").String()
	// dbPort := cfg.Section("DB").Key("Port").String()

	// netMethod := cfg.Section("Network").Key("Method").String()
	// netHost := cfg.Section("Network").Key("Host").String()
	// netPort := cfg.Section("Network").Key("Port").String()

	// result := []string{dbHost, dbUser, dbPassword, dbDatabase, dbPort, netMethod, netHost, netPort}
	// fmt.Println(result)

}
