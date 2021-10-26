package demo_html_template

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	e.GET("/index", getIndexHandler)
}
