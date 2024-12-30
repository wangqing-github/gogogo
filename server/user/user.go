package user

import (
	"encoding/json"
	"gogogo/center/redis"
	"log"
	"time"
)

var userMap map[string]*User

type User struct {
	AccountId  string
	ServerId   int64
	Id         string
	NickName   string
	Head       string
	Level      int32
	Gem        int64
	Gold       int64
	Exp        int32
	VipLevel   int32
	CreateTime int64
}

// 模拟玩家增加经验升级
func (user *User) addExp(num int32) {
	var maxLevel int32 = 100
	var oneLevelExp int32 = 10
	totalExp := user.Exp + num
	upLevel := totalExp / oneLevelExp
	if upLevel+user.Level > maxLevel {
		user.Level = maxLevel
		user.Exp = 0
	} else {
		user.Level += upLevel
		rest := totalExp % oneLevelExp
		user.Exp = rest
	}
}

func Create(name, AccountId string, ServerId int64) {
	user := User{
		Id:         "-1",
		AccountId:  AccountId,
		ServerId:   ServerId,
		NickName:   name,
		Level:      1,
		CreateTime: time.Now().Unix(),
	}
	redis.SetValue("user:"+user.Id, user)
	userMap[user.Id] = &user
}

func Info(id string) User {
	value := redis.GetByKey("user" + id)
	if len(value) == 0 {
		return User{Id: "-1"}
	}
	var user User
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {
		log.Fatalf("无法反序列化User: %v", err)
	}
	userMap[user.Id] = &user
	return user
}

func ChangeName(id, name string) string {
	user := userMap[id]
	if user == nil {
		return "玩家不存在"
	}
	checkWord(name)
	user.NickName = name
	redis.SetValue("user:"+user.Id, user)
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

func GetCacheUser(userId string) *User {
	return userMap[userId]
}
