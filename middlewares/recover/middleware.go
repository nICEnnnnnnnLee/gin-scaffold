package recover_mid

import "github.com/gin-gonic/gin"

func Middlewares(engine *gin.Engine) {
	engine.Use(Recover)
}
