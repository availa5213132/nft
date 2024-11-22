package main

import (
	"nft/server/core"
	"nft/server/flags"
	global "nft/server/gloabl"
	"nft/server/routers"
)

func main() {

	core.InitConf()                    //读取配置文件
	global.Log = core.InitLogger()     //初始化日志
	global.DB = core.InitGorm()        //连接数据库
	global.Redis = core.ConnectRedis() //连接redis
	// 命令行参数绑定
	option := flags.Parse()
	if option.Run() {
		return
	}
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
