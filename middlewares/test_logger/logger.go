package logger

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

//自定义中间件第1种定义方式
func Logger(c *gin.Context) {
	t := time.Now()
	fmt.Println("我是自定义中间件第1种定义方式---请求之前")
	//在gin上下文中定义一个变量
	c.Set("example", "CustomRouterMiddle1")
	//请求之前
	c.Next()
	fmt.Println("我是自定义中间件第1种定义方式---请求之后")
	//请求之后
	//计算整个请求过程耗时
	t2 := time.Since(t)
	log.Println(t2)
}

//自定义中间件第2种定义方式
func Logger2() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("我是自定义中间件第2种定义方式---请求之前")
		//在gin上下文中定义一个变量
		c.Set("example", "CustomRouterMiddle2")
		//请求之前
		c.Next()
		fmt.Println("我是自定义中间件第2种定义方式---请求之后")
		//请求之后
		//计算整个请求过程耗时
		t2 := time.Since(t)
		log.Println(t2)
	}
}

func LoggerInRouter(label string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("LoggerInRouter---请求之前 label:" + label)
		//在gin上下文中定义一个变量
		c.Set("example", "LoggerInRouter")
		//请求之前
		c.Next()
		fmt.Println("LoggerInRouter---请求之后 label:" + label)
		//请求之后
		//计算整个请求过程耗时
		t2 := time.Since(t)
		log.Println(t2)
	}
}
