package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"gogogo/center/account"
	"gogogo/center/server"
	"net/http"
)

type User struct {
	Name  string `form:"name"`
	Age   int    `form:"age"`
	Email string `form:"email"`
}

func Test() {
	r := gin.Default()
	{
		r.GET("/account/info", accountInfo)
		r.GET("/server/list", serverList)
	}

	{
		r.POST("/account/create", createAccount)
		r.POST("/account/updatePwd", updatePwd)
		r.POST("/account/addForm", userAddForm)
	}
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

func accountInfo(context *gin.Context) {
	account.Info(context)
}

func createAccount(context *gin.Context) {
	account.Create(context)
}

func updatePwd(context *gin.Context) {
	account.UpdatePwd(context)
}

func serverList(context *gin.Context) {
	list := server.List()
	context.JSON(200, list)
}

// userAddJson 模拟json数据解析
func userAddJson(context *gin.Context) {
	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		// 如果绑定失败，返回400错误
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 输出绑定成功的user数据
	fmt.Printf("Received user: %+v\n", user)
	context.String(200, "收到userAddJson的请求", 3)
}

// userAddJson 模拟表单数据解析
func userAddForm(context *gin.Context) {
	var user User
	if err := context.Bind(&user); err != nil {
		// 如果绑定失败，返回400错误
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 输出绑定成功的user数据
	fmt.Printf("Received user: %+v\n", user)
	context.String(200, "收到userAddForm的请求", 3)
	//直接响应结构体，自动转换成json
	context.JSON(200, user)
	//响应自定义封装json
	context.JSON(200, gin.H{"message": "成功", "status": 200})
	//传protobuf格式数据
	reps := []int64{int64(1), int64(2)}
	label := "label"
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	context.ProtoBuf(200, data)
}
