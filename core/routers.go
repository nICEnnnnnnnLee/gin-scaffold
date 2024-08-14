package core

import (
	"github.com/gin-gonic/gin"
)

var (
	rgDefault = &RouterGroup{
		id:            "default",
		routerOptions: []RouterOption{},
	}
)

type RouterOption func(*gin.Engine)
type RouterGroup struct {
	id            string
	routerOptions []RouterOption
}

func (rg *RouterGroup) Include(opts ...RouterOption) {
	rg.routerOptions = append(rg.routerOptions, opts...)
}

func (rg *RouterGroup) ApplyTo(engine *gin.Engine) {
	for _, opt := range rg.routerOptions {
		opt(engine)
	}
}

// var rgMap     = make(map[string]RouterGroup)
// func (rg *RouterGroup) Register() error {
// 	_, ok := rgMap[rg.id]
// 	if ok {
// 		return errors.New("routerGroupId 已经被注册")
// 	}
// 	rgMap[rg.id] = *rg
// 	return nil
// }
// // 注册路由组配置
// func RegisterRouterGroup(routerGroup *RouterGroup) error {
// 	_, ok := rgMap[routerGroup.id]
// 	if ok {
// 		return errors.New("routerGroupId 已经被注册")
// 	}
// 	rgMap[routerGroup.id] = *routerGroup
// 	return nil
// }

// 注册app的路由配置
func IncludeRouter(opts ...RouterOption) {
	rgDefault.Include(opts...)
}

// 将路由组配置实现
func ApplyRouterTo(engine *gin.Engine) {
	rgDefault.ApplyTo(engine)
}
