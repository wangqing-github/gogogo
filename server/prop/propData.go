package prop

import (
	"fmt"
	"gogogo/center/mysql"
)

var db = mysql.DB

func queryPropList(userId string) []*Prop {
	sqlStr := fmt.Sprintf("select * from prop wher userId = %s", userId)
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	defer rows.Close()
	result := make([]*Prop, 0)
	for rows.Next() {
		var prop Prop
		err := rows.Scan(&prop.UserId, &prop.Id, &prop.Num, &prop.PropType, &prop.ExpireTime)
		if err != nil {
			continue
		}
		result = append(result, &prop)
	}
	return result
}

func saveProp(prop *Prop) {
	sqlStr := "insert into prop value (?,?,?,?,?)"
	_, err := db.Exec(sqlStr, prop.UserId, prop.Id, prop.Num, prop.PropType, prop.ExpireTime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
	}
}

func updateProp(prop *Prop) {
	sqlStr := "Update prop set num=? where userId =? and id = ?"
	_, err := db.Exec(sqlStr, prop.Num, prop.UserId, prop.Id)
	if err != nil {
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
}
