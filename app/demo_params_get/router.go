package demo_params_get

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/user/:name/*action", getPathParamHandler)
	e.GET("/query", getQueryParamHandler)
}
