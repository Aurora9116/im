package router

import (
	"github.com/gin-gonic/gin"
	"im/middlewares"
	"im/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 用户登录
	r.POST("/login", service.Login)
	// 发送验证码
	r.POST("/send/code", service.SendCode)

	auth := r.Group("/u", middlewares.AuthCheck())

	// 用户详情
	auth.GET("/user/detail", service.UserDail)
	// 发送接收信息
	auth.GET("/websocket/message", service.WebsocketMessage)
	// 聊天记录列表
	auth.GET("/chat/list", service.ChatList)
	return r
}
