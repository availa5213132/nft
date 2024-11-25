package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
	"net/url"
	"nft/server/config"
	"nft/server/models"
	"nft/server/service/user_ser"
	"nft/server/service/wechat_ser"
	"nft/server/utils" // 假设 utils 包下有一些工具函数，如发起 HTTP 请求等
	"time"
)

// WeChatLoginView 微信登录的视图，用户点击登录时会跳转到这个路由
func (api *WeChatApi) WeChatLoginView(c *gin.Context) {
	// 微信的授权地址，跳转到微信的授权页面
	redirectUri := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=wechat&connect_redirect=1#wechat_redirect",
		config.GetWechatConfig().AppID,
		config.GetWechatConfig().RedirectURI)

	c.Redirect(http.StatusFound, redirectUri)
}

// WeChatQRCode 微信扫码登录，生成微信授权二维码
func (api *WeChatApi) WeChatQRCode(c *gin.Context) {
	// 获取微信的 AppID（从配置中获取）
	appID := config.GetWechatConfig().AppID
	redirectURI := config.GetWechatConfig().RedirectURI // 授权回调地址
	state := "123456"                                   // 防止 CSRF 攻击的随机字符串

	// 生成微信授权 URL
	authURL, err := generateWeChatAuthURL(appID, redirectURI, state)
	if err != nil {
		// 如果生成授权 URL 失败，返回 500 错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate auth url"})
		return
	}

	// 生成二维码并返回给用户
	err = generateQRCode(authURL, c)
	if err != nil {
		// 如果生成二维码失败，返回 500 错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate qr code"})
		return
	}
}

// generateWeChatAuthURL 生成微信授权 URL
func generateWeChatAuthURL(appID, redirectURI, state string) (string, error) {
	// 参数校验
	if appID == "" {
		return "", fmt.Errorf("appID 不能为空")
	}
	if redirectURI == "" {
		return "", fmt.Errorf("redirectURI 不能为空")
	}
	if state == "" {
		state = "123456" // 提供默认 state，避免 CSRF 校验失败
	}

	// URL 编码
	encodedRedirectURI := url.QueryEscape(redirectURI)

	// 构建微信授权 URL
	authURL := fmt.Sprintf(
		"https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=%s#wechat_redirect",
		appID, encodedRedirectURI, state,
	)

	return authURL, nil
}

// generateQRCode 生成二维码并将其发送给用户
func generateQRCode(url string, c *gin.Context) error {
	// 生成二维码并返回字节数据
	qrImage, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return err
	}

	// 设置响应的 Content-Type 为图片格式
	c.Header("Content-Type", "image/png")
	c.Status(http.StatusOK)
	c.Writer.Write(qrImage) // 直接返回二维码图片字节流
	return nil
}

// WeChatCallback 微信回调，获取 `code` 并使用 `code` 换取 `access_token`
func (api *WeChatApi) WeChatCallback(c *gin.Context) {
	// 获取 code
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}

	// 使用 code 获取 access_token
	accessToken, err := utils.GetWeChatAccessToken(config.GetWechatConfig().AppID, config.GetWechatConfig().AppSecret, code)
	if err != nil {
		log.Printf("Failed to get access token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get access token"})
		return
	}

	// 获取用户信息
	userInfo, err := utils.GetWeChatUserInfo(accessToken)
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user info"})
		return
	}

	// 查找用户（通过 OpenID）
	user, err := user_ser.GetUserByOpenID(userInfo.OpenID)
	if err != nil {
		log.Printf("Error retrieving user by OpenID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user"})
		return
	}

	// 如果用户不存在，进行注册
	if user == nil {
		// 创建新用户
		user = &models.UserModel{
			NickName: userInfo.Nickname,
			UserName: userInfo.Nickname, // 可以考虑使用昵称作为用户名
			Avatar:   userInfo.HeadImgURL,
			Password: "",                                  // 如果需要支持密码，初始化为空字符串
			Email:    "",                                  // 默认值为空
			Phone:    "",                                  // 默认值为空
			Token:    "",                                  // 不需要直接保存 token，JWT token 会单独生成
			RoleID:   func() *int { r := 2; return &r }(), // 默认为普通用户角色
			//Status_ID:     func() *int { s := 3; return &s }(),                // 假设 3 表示通过微信扫码登录
			LastLogin:     nil,                                                // 初始登录时间为空
			RegisteredAt:  func() *time.Time { t := time.Now(); return &t }(), // 注册时间
			WalletAddress: "",                                                 // 默认为空
			IsDelete:      "0",                                                // 用户未被删除
		}
		// 保存到数据库
		err = user_ser.RegisterUser(user)
		if err != nil {
			log.Printf("Error registering new user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
			return
		}
	}

	// 生成 JWT Token
	token, err := wechat_ser.GenerateJWT(user)
	if err != nil {
		log.Printf("Failed to generate JWT token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// 将 token 赋值给用户并返回用户信息
	user.Token = token

	// 返回成功的登录信息，包括用户信息和生成的 token
	c.JSON(http.StatusOK, gin.H{
		"message":   "微信登录成功",
		"user_info": user,
		"token":     token, // 返回 JWT token
	})
}

// WechatVerify 微信 Token 验证
func (api *WeChatApi) WechatVerify(c *gin.Context) {
	// 获取微信传递的参数
	signature := c.DefaultQuery("signature", "")
	timestamp := c.DefaultQuery("timestamp", "")
	nonce := c.DefaultQuery("nonce", "")
	echostr := c.DefaultQuery("echostr", "")

	// 校验 signature 参数
	if wechat_ser.CheckSignature(signature, timestamp, nonce) {
		// 验证成功，返回 echostr 参数
		c.String(http.StatusOK, echostr)
	} else {
		// 验证失败
		c.String(http.StatusUnauthorized, "Unauthorized")
	}
}
