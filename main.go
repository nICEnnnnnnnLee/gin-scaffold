package main

import (
	"context"
	"fmt"
	"gin_scaffold/app/demo_404_not_found"
	"gin_scaffold/app/demo_data_bind"
	"gin_scaffold/app/demo_html_template"
	"gin_scaffold/app/demo_params_get"
	"gin_scaffold/app/demo_redirect"
	"gin_scaffold/app/demo_router_middleware"
	"gin_scaffold/app/demo_sync_async"
	"gin_scaffold/middlewares"
	"gin_scaffold/middlewares/modify_header"
	recover_mid "gin_scaffold/middlewares/recover"
	logger "gin_scaffold/middlewares/test_logger"
	"gin_scaffold/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	//创建一个无中间件路由
	engine := gin.New()
	// 默认启动方式，包含 Logger、Recovery 中间件
	// engine := gin.Default()

	// 初始化中间件, 必须要在路由之前挂载
	{
		middlewares.Include(
			modify_header.Middlewares,
			logger.Middlewares,
			recover_mid.Middlewares,
		)
		middlewares.ApplyTo(engine)
	}

	// 初始化路由
	{
		routers.Include(
			demo_params_get.Routers,
			demo_data_bind.Routers,
			demo_html_template.Routers,
			demo_sync_async.Routers,
			demo_router_middleware.Routers,
			demo_redirect.Routers,
			demo_404_not_found.Routers,
		)
		routers.ApplyTo(engine)
	}

	// 设置404
	engine.NoRoute(demo_404_not_found.NotFoundHandler)
	// 设置模板路径
	engine.LoadHTMLGlob("public/templates/**/*")
	// 设置静态资源路径
	engine.Static("/assets", "./public/assets")

	// 开始运行
	// 使用gin默认方式
	// if err := engine.Run(":8081"); err != nil {
	// 	fmt.Printf("startup service failed, err:%v \n", err)
	// }

	// 自定义http.server
	server := &http.Server{
		Addr:           ":8081",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go server.ListenAndServe()
	// 设置优雅退出
	gracefulExitWeb(server)
}

func gracefulExitWeb(server *http.Server) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-done

	now := time.Now()
	// 10秒超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		fmt.Println("err", err)
	}

	// 看看实际退出所耗费的时间
	fmt.Println("------exited--------", time.Since(now))
}
