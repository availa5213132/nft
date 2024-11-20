package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建 Gin 实例
	r := gin.Default()

	// 配置路由
	r.GET("/", func(c *gin.Context) {
		// 返回 JSON 格式的响应
		c.JSON(200, gin.H{
			"message": "Welcome to the NFT Trading Platform API!",
		})
	})

	// 启动 API 服务，默认监听 8080 端口
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Server failed to start, ", err)
	}

	// 打印 API 服务启动信息
	fmt.Println("API Server is running on port 8080")
}
