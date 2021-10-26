package demo_html_template

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index.html", gin.H{"title": "我是gin", "name": "you"})
}
