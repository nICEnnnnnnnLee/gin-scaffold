package demo_router_middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func middlewareTestHandler(c *gin.Context) {
	fmt.Println("middlewareTestHandler开始执行")
	val := c.MustGet("example")
	result := fmt.Sprintf("来自中间件设置的值: %v", val)
	c.String(http.StatusOK, result)
	fmt.Println("middlewareTestHandler结束执行")
}
