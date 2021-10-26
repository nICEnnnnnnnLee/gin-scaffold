package demo_data_bind

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//定义接收数据的结构体
type Login struct {
	//binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

type LoginStructError validator.ValidationErrors

func (errs LoginStructError) Error() string {
	/*
		输入空信息时
		errs[0].Tag() required
		errs[0].ActualTag() required
		errs[0].Namespace() Login.User
		errs[0].StructNamespace() Login.User
		errs[0].Field() User
		errs[0].StructField() User
		errs[0].Value()
		errs[0].Param()
		errs[0].StructField() User
	*/
	fmt.Printf("errs[0].Tag() %+v \n", errs[0].Tag())
	fmt.Printf("errs[0].ActualTag() %+v \n", errs[0].ActualTag())
	fmt.Printf("errs[0].Namespace() %+v \n", errs[0].Namespace())
	fmt.Printf("errs[0].StructNamespace() %+v \n", errs[0].StructNamespace())
	fmt.Printf("errs[0].Field() %+v \n", errs[0].Field())
	fmt.Printf("errs[0].StructField() %+v \n", errs[0].StructField())
	fmt.Printf("errs[0].Value() %+v \n", errs[0].Value())
	fmt.Printf("errs[0].Param() %+v \n", errs[0].Param())
	fmt.Printf("errs[0].StructField() %+v \n", errs[0].StructField())
	return fmt.Sprintf("%s 出现错误", errs[0].StructNamespace())
}

func getLoginHandler(c *gin.Context) {
	var login Login
	// 将request的body中的数据，自动按照json格式解析到结构体
	// c.ShouldBindJSON(&json)
	// c.Bind(&login)
	if err := c.ShouldBindQuery(&login); err != nil {
		// ?username=xxx&password=xx
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": LoginStructError(err.(validator.ValidationErrors)).Error()})
		return
	}

	//判断用户名密码是否正确
	if login.User != "root" || login.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	fmt.Println(login)
	// c.JSON(http.StatusOK, gin.H{"status": 200})
	c.YAML(200, gin.H{"name": login.User, "password": login.Pssword})

}
