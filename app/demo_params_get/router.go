package demo_params_get

import (
	"gin_scaffold/core"

	"github.com/gin-gonic/gin"
)

// 注册路由
func init() {
	core.IncludeRouter(Routers)
}

func Routers(e *gin.Engine) {
	e.GET("/user/:name/*action", getPathParamHandler)
	e.GET("/query", getQueryParamHandler)
}
