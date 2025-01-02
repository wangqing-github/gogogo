package account

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gogogo/center/redis"
	"gogogo/util"
	"log"
	"net/http"
)

var accountMap map[int64]*Account

type Account struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Ip       string
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	PassWord string `json:"passWord"`
	DeviceId string `json:"deviceId"`
	ServerId int64  `json:"serverId"`
}

type updatePwd struct {
	UserId int64  `json:"userId"`
	Token  string `json:"token"`
	NewPwd string `json:"newPwd"`
}

// Info 模拟通过id查询角色信息
func Info(context *gin.Context) {
	userId := context.DefaultQuery("userId", "0")
	value := redis.GetByKey("account" + userId)
	if len(value) == 0 {
		context.String(200, "未查找到用户%d", value)
		return
	}
	var user Account
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {
		log.Fatalf("无法反序列化User: %v", err)
	}
	accountMap[user.Id] = &user
	context.JSON(200, user)
}

func Create(context *gin.Context) {
	var newAccount Account
	if err := context.ShouldBindJSON(&newAccount); err != nil {
		// 如果绑定失败，返回400错误
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newAccount.Ip = context.ClientIP()
	newAccount.Id = util.NewSnowflakeId()
	userJson, _ := json.Marshal(newAccount)
	redis.SetValue(fmt.Sprintf("%s%d", "account", newAccount.Id), userJson)
	fmt.Println("创建角色成功")
	context.JSON(200, newAccount)

}

func UpdatePwd(context *gin.Context) {
	var requestBody updatePwd
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// todo 短信校验，验证token
	if len(requestBody.Token) == 0 || len(requestBody.NewPwd) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "token或者密码不能为空"})
		return
	}
	account := accountMap[requestBody.UserId]
	account.PassWord = util.Md5Encode(requestBody.NewPwd)
	userJson, _ := json.Marshal(account)
	redis.SetValue(fmt.Sprintf("%s%d", "account", account.Id), userJson)
	context.JSON(200, "success")
}
