package account

import (
	"encoding/json"
	"fmt"
	"gogogo/center/redis"
	"log"
	"net"
)

var accountMap map[string]*Account

type Account struct {
	id       string
	name     string
	ip       string
	serverId int64
	coon     net.Conn
}

// Info 模拟通过id查询角色信息
func Info(id string) Account {
	value := redis.GetByKey("account" + id)
	if len(value) == 0 {
		return Account{id: "-1"}
	}
	var user Account
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {
		log.Fatalf("无法反序列化User: %v", err)
	}
	accountMap[user.id] = &user
	return user
}

func Create() {
	user := Account{
		id:   "111",
		name: "王庆",
		ip:   "成都",
	}
	userJson, _ := json.Marshal(user)
	redis.SetValue("account"+user.id, userJson)
	fmt.Println("创建角色成功")
}
