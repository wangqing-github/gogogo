package net

import (
	"github.com/magiclvzs/antnet"
	pb "gogogo/pb/msg"
)

var ExtNetHandler = &antnet.DefMsgHandler{}
var PbParser = &antnet.Parser{Type: antnet.ParserTypePB}

func AntNet() {
	handlerInit()
	antnet.StartServer("ws://:5000", antnet.MsgTypeMsg, ExtNetHandler, PbParser)
	antnet.WaitForSystemExit(nil)
}

func handlerInit() {
	//查找用户信息
	addHandler(pb.Pb_CMD_USER, pb.Pb_ACT_USER_INFO, &pb.Pb_RqstUserInfo{}, &pb.Pb_RespUserInfo{}, handlerUserInfo)
	//创建用户
	addHandler(pb.Pb_CMD_USER, pb.Pb_ACT_USER_CREATE, &pb.Pb_RqstCreateUser{}, &pb.Pb_RespCreateUser{}, handlerCreateUser)
}

func addHandler(cmd pb.Pb_CMD, act pb.Pb_ACT, c2s interface{}, s2c interface{}, fun antnet.HandlerFunc) {
	PbParser.Register(uint8(cmd), uint8(act), c2s, s2c)
	ExtNetHandler.Register(uint8(cmd), uint8(act), fun)
}

func handlerUserInfo(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	//执行查找用户信息的一些业务逻辑
	return true
}

func handlerCreateUser(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	//执行创建用户的一些业务逻辑
	return true
}
