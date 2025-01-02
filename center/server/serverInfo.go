package server

import (
	"encoding/json"
	"fmt"
	"gogogo/center/redis"
)

type Info struct {
	ServerName string
	Ip         string
	Port       int
	Status     int
}

// Init 将模拟服务器信息填进去
func Init() {
	info1 := Info{
		ServerName: "王庆01",
		Ip:         "127.0.0.1",
		Port:       8081,
		Status:     1,
	}

	info2 := Info{
		ServerName: "王庆02",
		Ip:         "127.0.0.1",
		Port:       8081,
		Status:     1,
	}

	serverList := make([]Info, 0, 10)
	serverList = append(serverList, info1, info2)
	sliceJSON, _ := json.Marshal(serverList)
	redis.SetValue("serverList", sliceJSON)
}

// List 查找服务器列表
func List() *[]Info {
	fmt.Println("获取服务器列表")
	serverList := redis.GetByKey("serverList")
	var mySliceFromRedis []Info
	err := json.Unmarshal([]byte(serverList), &mySliceFromRedis)
	if err != nil {
		return nil
	}
	return &mySliceFromRedis
}
