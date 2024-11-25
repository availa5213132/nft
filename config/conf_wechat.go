package config

type WechatConfig struct {
	AppID       string
	AppSecret   string
	RedirectURI string
}

// GetWechatConfig 获取微信相关配置
func GetWechatConfig() *WechatConfig {
	return &WechatConfig{
		AppID:       "wx55b91f835d27d393",
		AppSecret:   "79e5eb42d91cdaf04dfcd73f5705de16",
		RedirectURI: "http://45i75j.natappfree.cc/api/wechat/callback",
	}
}
