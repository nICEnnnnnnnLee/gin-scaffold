package demo_data_bind

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	group := e.Group("/group")
	group.GET("/login", getLoginHandler)
}
