package main

import (
	fmt
	log
	github.comgin-gonicgin
)

func main() {
	 创建Gin实例
	r = gin.Default()

	 配置路由
	r.GET(, func(c gin.Context) {
		c.JSON(200, gin.H{
			message Welcome to the NFT Trading Platform API!,
		})
	})

	 启动API服务
	err = r.Run(8080)  默认监听8080端口
	if err != nil {
		log.Fatal(Server failed to start, err)
	}

	fmt.Println(API Server is running on port 8080)
}
