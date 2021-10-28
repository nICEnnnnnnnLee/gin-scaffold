package demo_file_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/file/*path", FileserverHandler)
}

func FileserverHandler(c *gin.Context) {
	handler := http.FileServer(http.Dir("./public/assets"))
	handler = http.StripPrefix("/file/", handler)
	handler.ServeHTTP(c.Writer, c.Request)

}
