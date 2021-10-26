package demo_sync_async

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/long_async", asyncHandler)
	e.GET("/long_sync", syncHandler)
}
