package routers

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	//创建一个无中间件路由
	// router := gin.New()
	// 默认启动方式，包含 Logger、Recovery 中间件
	router := gin.Default()
	for _, opt := range options {
		opt(router)
	}
	return router
}

func ApplyTo(engine *gin.Engine) {
	for _, opt := range options {
		opt(engine)
	}
}
