package demo_redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/redirect", RedirectHandler)
}

func RedirectHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/index")
}
