package friend

import (
	"fmt"
	"gogogo/center/mysql"
	"strings"
)

var db = mysql.DB

func queryFriendList(userId string) []*Info {
	sqlStr := "select * from friend wher userId = " + userId
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	defer rows.Close()
	result := make([]*Info, 0)
	for rows.Next() {
		var info Info
		err := rows.Scan(&info.UserId, &info.FriendId, &info.CreateTime, &info.Remark, &info.FriendShip)
		if err != nil {
			continue
		}
		result = append(result, &info)
	}
	return result
}

func queryApplyList(userId string) []*ApplyInfo {
	sqlStr := "select * from friend_apply wher targetId = " + userId
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	defer rows.Close()
	result := make([]*ApplyInfo, 0)
	for rows.Next() {
		var info ApplyInfo
		err := rows.Scan(&info.SendId, &info.TargetId, &info.CreateTime, &info.ExpireTime)
		if err != nil {
			continue
		}
		result = append(result, &info)
	}
	return result
}

// saveFriendInfo 跟下面一个，只是模拟批量，实际生产不推荐
func saveFriendInfo(list ...Info) {
	for _, info := range list {
		sqlStr := "insert into friend value (?,?,?,?,?)"
		_, err := db.Exec(sqlStr, info.UserId, info.FriendId, info.CreateTime, info.Remark, info.FriendShip)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			continue
		}
	}
}

func saveApply(list ...ApplyInfo) {
	for _, info := range list {
		sqlStr := "insert into friend_apply value (?,?,?,?)"
		_, err := db.Exec(sqlStr, info.SendId, info.TargetId, info.CreateTime, info.ExpireTime)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			continue
		}
	}
}

func deleteFriendById(id string, ids []string) {
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "?"
	}
	IDs := make([]interface{}, len(ids))
	for i, idStr := range ids {
		IDs[i] = idStr
	}
	sqlStr := fmt.Sprintf("DELETE FROM friend WHERE userId =%s and friendId IN (%s)", id, strings.Join(placeholders, ","))
	_, err := db.Exec(sqlStr, IDs...)
	if err != nil {
		return
	}
}

func deleteApply(id string, ids []string) {
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "?"
	}
	IDs := make([]interface{}, len(ids))
	for i, idStr := range ids {
		IDs[i] = idStr
	}
	sqlStr := fmt.Sprintf("DELETE FROM friend_apply WHERE targetId =%s and SendId IN (%s)", id, strings.Join(placeholders, ","))
	_, err := db.Exec(sqlStr, IDs...)
	if err != nil {
		return
	}
}
