package demo_html_template

import (
	"gin_scaffold/core"

	"github.com/gin-gonic/gin"
)

// 注册路由
func init() {
	core.IncludeRouter(Routers)
}

func Routers(e *gin.Engine) {

	e.GET("/index", getIndexHandler)
}
