package api

import (
	"nft/server/api/user_api"
)

type ApiGroup struct {
	UserApi user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
