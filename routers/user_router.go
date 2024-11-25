package routers

import "nft/server/api"

func (router RouterGroup) UserRouter() {
	app1 := api.ApiGroupApp.UserApi
	app2 := api.ApiGroupApp.WeChatApi
	router.GET("set", app1.QQLoginView)
	router.GET("wechat/login", app2.WeChatLoginView)     // 微信登录页面
	router.GET("wechat/callback", app2.WeChatCallback)   // 微信回调
	router.GET("wechat/qr", app2.WeChatQRCode)           // 获取微信扫码二维码
	router.GET("wechat/wechatVerify", app2.WechatVerify) // 获取微信用户信息
	//router.GET("wechat/refresh_token", app2.WeChatRefreshToken) // 刷新微信 access_token
	//router.GET("wechat/error", app2.WeChatLoginError)           // 处理授权失败

}
