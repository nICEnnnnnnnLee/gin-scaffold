package demo_sync_async

import (
	"gin_scaffold/core"

	"github.com/gin-gonic/gin"
)

// 注册路由
func init() {
	core.IncludeRouter(Routers)
}

func Routers(e *gin.Engine) {
	e.GET("/long_async", asyncHandler)
	e.GET("/long_sync", syncHandler)
}
