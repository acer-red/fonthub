package system

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(data []byte) string {

	// 创建一个SHA-256哈希对象
	hash := sha256.New()

	// 写入输入数据
	hash.Write(data)

	// 计算哈希值
	hashValue := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hashValue)

}
