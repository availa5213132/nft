package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WeChatAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetWeChatAccessToken(appID, appSecret, code string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", appID, appSecret, code)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var accessToken WeChatAccessToken
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return "", err
	}

	if accessToken.AccessToken == "" {
		return "", errors.New("failed to get access token")
	}

	return accessToken.AccessToken, nil
}

// WeChatUserInfo 表示从微信获取到的用户信息结构
type WeChatUserInfo struct {
	OpenID     string `json:"openid"`
	Nickname   string `json:"nickname"`
	HeadImgURL string `json:"headimgurl"`
	Sex        int    `json:"sex"`
	Language   string `json:"language"`
	Country    string `json:"country"`
	Province   string `json:"province"`
	City       string `json:"city"`
}

// GetWeChatUserInfo 获取微信用户信息
func GetWeChatUserInfo(accessToken string) (*WeChatUserInfo, error) {
	// 构建请求URL
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", accessToken)

	// 发送请求
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error sending request to WeChat API: %v", err)
		return nil, errors.New("failed to send request to WeChat API")
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, errors.New("failed to read response from WeChat API")
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error from WeChat API: %v", string(body))
		return nil, fmt.Errorf("WeChat API returned error: %s", string(body))
	}

	// 解析返回的 JSON 数据
	var userInfo WeChatUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		log.Printf("Error unmarshaling JSON response: %v", err)
		return nil, errors.New("failed to parse user info")
	}

	// 返回用户信息
	return &userInfo, nil
}

// GetWeChatQRCode 获取微信扫码二维码的图片
func GetWeChatQRCode(appID, secret, scene string) ([]byte, error) {
	// 获取 access_token
	tokenURL := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appID, secret)
	resp, err := http.Get(tokenURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenResult map[string]interface{}
	if err := json.Unmarshal(body, &tokenResult); err != nil {
		return nil, err
	}

	accessToken, ok := tokenResult["access_token"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to get access token")
	}

	// 创建二维码
	qrCreateURL := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s", accessToken)

	// 生成二维码的参数
	qrParams := map[string]interface{}{
		"action_name": "QR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_id": scene,
			},
		},
	}

	qrData, err := json.Marshal(qrParams)
	if err != nil {
		return nil, err
	}

	// 发送 POST 请求创建二维码
	resp, err = http.Post(qrCreateURL, "application/json", bytes.NewReader(qrData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var createResult map[string]interface{}
	if err := json.Unmarshal(body, &createResult); err != nil {
		return nil, err
	}

	// 获取二维码的 ticket
	ticket, ok := createResult["ticket"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to get ticket")
	}

	// 获取二维码的图片URL
	// 通过ticket来获取二维码图片
	qrImageURL := fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s", ticket)

	// 获取二维码图片
	resp, err = http.Get(qrImageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 返回图片内容
	return ioutil.ReadAll(resp.Body)
}
