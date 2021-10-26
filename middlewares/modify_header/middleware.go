package modify_header

import (
	"github.com/gin-gonic/gin"
)

func Middlewares(engine *gin.Engine) {
	engine.Use(modifyHeader)
}

func modifyHeader(c *gin.Context) {
	// 下面等价于 c.Header("Server", "Tomato 1.0.0")
	w := c.Writer
	w.Header().Set("Server", "Tomato 1.0.0")
	// 必须要在 w.WriteHeader(statusCode int) 之前调用
	c.Next()
}
