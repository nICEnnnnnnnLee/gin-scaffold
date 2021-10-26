package logger

import "github.com/gin-gonic/gin"

func Middlewares(engine *gin.Engine) {
	engine.Use(Logger)
	engine.Use(Logger2())
}
