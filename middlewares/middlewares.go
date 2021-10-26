package middlewares

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的中间件配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

func ApplyTo(engine *gin.Engine) {
	for _, opt := range options {
		opt(engine)
	}
}
