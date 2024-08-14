package logger

import (
	"gin_scaffold/core"

	"github.com/gin-gonic/gin"
)

func init() {
	core.IncludeMiddleware(Middlewares)
}

func Middlewares(engine *gin.Engine) {
	engine.Use(Logger)
	engine.Use(Logger2())
}
