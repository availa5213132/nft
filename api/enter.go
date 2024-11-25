package api

import (
	"nft/server/api/user_api"
)

type ApiGroup struct {
	UserApi   user_api.UserApi
	WeChatApi user_api.WeChatApi
}

var ApiGroupApp = new(ApiGroup)
