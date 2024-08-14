package demo_data_bind

import (
	"gin_scaffold/core"

	"github.com/gin-gonic/gin"
)

// 注册路由
func init() {
	core.IncludeRouter(Routers)
}

func Routers(e *gin.Engine) {
	group := e.Group("/group")
	group.GET("/login", getLoginHandler)
}
