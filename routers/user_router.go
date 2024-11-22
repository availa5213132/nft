package routers

import "nft/server/api"

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.GET("set", app.QQLoginView)
}
