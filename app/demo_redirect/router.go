package demo_redirect

import (
	"gin_scaffold/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册路由
func init() {
	core.IncludeRouter(Routers)
}

func Routers(e *gin.Engine) {
	e.GET("/redirect", RedirectHandler)
}

func RedirectHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/index")
}
