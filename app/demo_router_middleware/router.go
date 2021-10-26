package demo_router_middleware

import (
	logger "gin_scaffold/middlewares/test_logger"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/middleware",
		logger.LoggerInRouter("LoggerInRouter 1"),
		middlewareTestHandler,
		logger.LoggerInRouter("LoggerInRouter 2"),
	)

	group := e.Group("/middleGroup", logger.LoggerInRouter("LoggerInRouter 1"), logger.LoggerInRouter("LoggerInRouter 2"))
	// group.Use(logger.LoggerInRouter("LoggerInRouter 1"))
	group.GET("/middleware", middlewareTestHandler)
	group.GET("/middleware2",
		logger.LoggerInRouter("LoggerInRouter 3"),
		middlewareTestHandler,
		logger.LoggerInRouter("LoggerInRouter 4"),
	)
}
