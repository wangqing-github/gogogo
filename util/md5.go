package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(value string) string {

	// 创建一个新的md5.Hash对象
	hasher := md5.New()

	// 将字符串转换为字节切片并写入hasher
	hasher.Write([]byte(value))

	// 计算哈希值的字节切片
	hashBytes := hasher.Sum(nil)

	// 将字节切片转换为十六进制字符串
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
