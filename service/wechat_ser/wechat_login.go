package wechat_ser

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"nft/server/models"
	"nft/server/utils/jwts"
	"sort"
)

// GenerateJWT 使用 jwts.GenToken 生成 JWT Token
func GenerateJWT(user *models.UserModel) (string, error) {
	// 生成 JwtPayLoad
	payload := jwts.JwtPayLoad{
		Username: user.UserName, // 使用用户名
		NickName: user.NickName, // 使用昵称
		Role: func() int {
			if user.RoleID != nil {
				return *user.RoleID
			}
			return 0 // 默认角色值
		}(),
		UserID: user.ID, // 使用用户 ID
	}

	// 调用 jwts.GenToken 创建 token
	token, err := jwts.GenToken(payload)
	if err != nil {
		log.Printf("JWT Token 生成失败: %v", err)
		return "", err
	}

	return token, nil
}

const (
	// Token 填写你自己在微信公众平台设置的 Token
	Token = "123456"
)

// CheckSignature 校验 signature 参数
func CheckSignature(signature, timestamp, nonce string) bool {
	// 将 token、timestamp 和 nonce 按字典序排序
	arr := []string{Token, timestamp, nonce}
	sort.Strings(arr)

	// 拼接成一个字符串
	str := arr[0] + arr[1] + arr[2]

	// 对拼接后的字符串进行 sha1 加密
	hash := sha1.New()
	hash.Write([]byte(str))
	hashSum := hash.Sum(nil)
	hashStr := hex.EncodeToString(hashSum)

	// 将加密后的结果与 signature 比较
	return hashStr == signature
}
