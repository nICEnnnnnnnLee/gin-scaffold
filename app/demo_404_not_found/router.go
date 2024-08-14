package demo_404_not_found

import (
	"gin_scaffold/core"
	"io"

	"github.com/gin-gonic/gin"
)

// 注册路由
func init() {
	core.IncludeRouter(Routers)
}

func Routers(e *gin.Engine) {
	e.GET("/404", NotFoundHandler)
}

func NotFoundHandler(c *gin.Context) {
	// c.Status(404)
	// c.Stream(func(w io.Writer) bool {
	// 	_, err := io.WriteString(w, "这是一个尝试使用stream来写的404消息")
	// 	return err != nil
	// })

	w := c.Writer
	const msg = "这是一个尝试使用stream来写的404消息~~"
	// w.Header().Set("Content-Length", strconv.Itoa(len(msg))) // 可以注释掉, 会自动处理
	w.WriteHeader(404)
	io.WriteString(w, msg)
	// // c.Writer.CloseNotify()
}
