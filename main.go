package main

import (
	"gogogo/center/redis"
	"gogogo/center/server"
)

func main() {
	redis.InitRedis()
	server.Init()
	server.List()
	//net.StartServer()
	//time.Sleep(3000)
	//net.StartClient()
}
