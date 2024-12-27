package user

import (
	"encoding/json"
	"gogogo/center/redis"
	"log"
	"time"
)

var userMap map[string]*User

type User struct {
	accountId  string
	serverId   int64
	id         string
	nickName   string
	level      int32
	gem        int64
	gold       int64
	exp        int32
	vipLevel   int32
	createTime int64
}

// 模拟玩家增加经验升级
func (user *User) addExp(num int32) {
	var maxLevel int32 = 100
	var oneLevelExp int32 = 10
	totalExp := user.exp + num
	upLevel := totalExp / oneLevelExp
	if upLevel+user.level > maxLevel {
		user.level = maxLevel
		user.exp = 0
	} else {
		user.level += upLevel
		rest := totalExp % oneLevelExp
		user.exp = rest
	}
}

func Create(name, accountId string, serverId int64) {
	user := User{id: "-1",
		accountId:  accountId,
		serverId:   serverId,
		nickName:   name,
		level:      1,
		createTime: time.Now().Unix(),
	}
	redis.SetValue("user:"+user.id, user)
	userMap[user.id] = &user
}

func Info(id string) User {
	value := redis.GetByKey("user" + id)
	if len(value) == 0 {
		return User{id: "-1"}
	}
	var user User
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {
		log.Fatalf("无法反序列化User: %v", err)
	}
	userMap[user.id] = &user
	return user
}

func ChangeName(id, name string) string {
	user := userMap[id]
	if user == nil {
		return "玩家不存在"
	}
	checkWord(name)
	user.nickName = name
	redis.SetValue("user:"+user.id, user)
	return "修改成功"
}

// 敏感词检测 todo
func checkWord(name string) {

}

// AddExp 增加经验 todo
func AddExp(id string, num int32) {
	user := userMap[id]
	if user == nil {
		return
	}
	user.addExp(num)
}
