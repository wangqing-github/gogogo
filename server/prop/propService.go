package prop

import (
	"gogogo/center/msg"
)

type Prop struct {
	UserId     string
	Id         int32
	Num        int64
	PropType   int32
	ExpireTime int32
}

type Config struct {
	Id         int32
	PropType   int32
	Name       string
	Desc       string
	ExpireTime int32
}

var propCache map[string]map[int32]*Prop
var configCache map[int32]*Config

// GetPropList
func GetPropList(userId string) msg.Response {
	propInfo := propCache[userId]
	if propInfo == nil {
		propInfo = make(map[int32]*Prop)
		list := queryPropList(userId)
		for _, v := range list {
			propInfo[v.Id] = v
		}
		propCache[userId] = propInfo
	}
	return msg.Success(propInfo)
}

func AddProp(userId string, id int32, num int64) msg.Response {
	config := configCache[id]
	if config == nil {
		return msg.Fail("没有找到道具配置")
	}
	prop := propCache[userId][id]
	if prop == nil {
		prop = &Prop{
			UserId:   userId,
			Id:       config.Id,
			PropType: config.PropType,
			Num:      num,
		}
		saveProp(prop)
	} else {
		prop.Num += num
		updateProp(prop)
	}
	return msg.Success(nil)
}
