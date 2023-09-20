package main

import (
	"github.com/gin-gonic/gin"
	"study/server"
)

type Server struct {
	router *gin.Engine
}

func (s *Server) setupRouter() {
	// 添加更多的路由和处理函数...
	s.router.POST("/Test", server.Test)
}

func main() {
	// 创建一个默认的Gin引擎
	server := &Server{
		router: gin.Default(),
	}

	// 设置路由和处理函数
	server.setupRouter()

	server.router.Run(":8080")
}
