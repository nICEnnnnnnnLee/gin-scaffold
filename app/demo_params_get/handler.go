package demo_params_get

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func getPathParamHandler(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	action = strings.Trim(action, "/")
	result := fmt.Sprintf("name: %v, action: %v", name, action)
	c.String(http.StatusOK, result)
}
func getQueryParamHandler(c *gin.Context) {
	key := c.DefaultQuery("key", "defaultQueryValue")
	result := fmt.Sprintf("key: %v", key)
	c.String(http.StatusOK, result)
}
