package friend

import (
	"gogogo/center/msg"
	"gogogo/server/user"
	"gogogo/util"
	"time"
)

type Info struct {
	UserId     string
	FriendId   string
	CreateTime int64
	Remark     string //备注
	FriendShip int32  //亲密度
	Name       string
	Head       string
	Level      int32
}

type ApplyInfo struct {
	SendId     string
	TargetId   string
	CreateTime int64
	ExpireTime int64
	Name       string
	Head       string
	Level      int32
}

var friendCache map[string][]*Info
var applyCache map[string][]*ApplyInfo

// friendList 好友列表
func friendList(userId string) msg.Response {
	list := queryFriendList(userId)
	for _, info := range list {
		cacheUser := user.GetCacheUser(info.FriendId)
		info.Head = cacheUser.Head
		info.Level = cacheUser.Level
		info.Name = cacheUser.NickName
	}
	friendCache[userId] = list
	return msg.Success(list)
}

// ApplyList 申请列表
func ApplyList(userId string) msg.Response {
	list := queryApplyList(userId)
	for _, info := range list {
		cacheUser := user.GetCacheUser(info.SendId)
		info.Head = cacheUser.Head
		info.Level = cacheUser.Level
		info.Name = cacheUser.NickName
	}
	applyCache[userId] = list
	return msg.Success(list)
}

// sendApply 发送好友申请，过期时间7天
func sendApply(userId, targetId string) msg.Response {
	//如果对方也给我发送了，则直接走添加好友流程
	infos := applyCache[userId]
	for _, i := range infos {
		if i.SendId == targetId {
			addFriend(userId, targetId)
			return msg.Success(nil)
		}
	}
	now := time.Now().Unix()
	apply := ApplyInfo{
		SendId:     userId,
		TargetId:   targetId,
		CreateTime: now,
		ExpireTime: now + 60*60*24*7,
	}
	saveApply(apply)
	return msg.Success(nil)
}

func addFriend(userId string, friendId string) msg.Response {
	info := &Info{
		UserId:     userId,
		FriendId:   friendId,
		CreateTime: time.Now().Unix(),
	}
	info2 := &Info{
		UserId:     friendId,
		FriendId:   userId,
		CreateTime: time.Now().Unix(),
	}
	saveFriendInfo(*info, *info2)
	infos := friendCache[userId]
	infos = append(infos, info)
	return msg.Success(nil)
}

// 删除好友
func deleteFriend(userId string, ids []string) msg.Response {
	deleteFriendById(userId, ids)
	//清楚缓存
	infos := friendCache[userId]
	var newInfos []*Info
	for _, v := range infos {
		if !util.SliceContains(ids, v.FriendId) {
			newInfos = append(newInfos, v)
		}
	}
	friendCache[userId] = newInfos
	return msg.Success(nil)
}

// dealApply 处理申请列表
func dealApply(userId string, result int, ids []string) msg.Response {
	//1是同意，0是拒绝
	if result == 1 {
		for _, i := range ids {
			addFriend(userId, i)
		}
	}
	//删除申请列表
	deleteApply(userId, ids)
	//清楚缓存
	return msg.Success(ids)
}
