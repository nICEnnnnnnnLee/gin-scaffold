package core

import "github.com/gin-gonic/gin"

var (
	midDefault = &MiddlewareGroup{
		id:                "default",
		MiddlewareOptions: []MiddlewareOption{},
	}
)

type MiddlewareOption func(*gin.Engine)
type MiddlewareGroup struct {
	id                string
	MiddlewareOptions []MiddlewareOption
}

func (rg *MiddlewareGroup) Include(opts ...MiddlewareOption) {
	rg.MiddlewareOptions = append(rg.MiddlewareOptions, opts...)
}

func (rg *MiddlewareGroup) ApplyTo(engine *gin.Engine) {
	for _, opt := range rg.MiddlewareOptions {
		opt(engine)
	}
}

func IncludeMiddleware(opts ...MiddlewareOption) {
	midDefault.Include(opts...)
}

func ApplyMiddlewareTo(engine *gin.Engine) {
	midDefault.ApplyTo(engine)
}
