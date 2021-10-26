package recover_mid

import (
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			c.JSON(200, gin.H{
				"code": "444",
				"msg":  "服务器内部错误",
			})
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}
