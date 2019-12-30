package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var(
	router *gin.Engine
)

func InitRouter(r *gin.Engine)  {
	if r == nil {
		return
	}
	router=r
	routers()
}

func routers()  {
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/html/handpage.html")
	})
	router.POST("/inquire",inquire)
}

func inquire(c *gin.Context)  {
	switch c.PostForm("select_engine"){
	case "baidu":
		c.Redirect(http.StatusMovedPermanently,"https://www.baidu.com/s?ie=UTF-8&wd="+c.PostForm("key"))
		break
	case "google":
		c.Redirect(http.StatusMovedPermanently,"https://www.google.com/search?q="+c.PostForm("key"))
		break
	case "bing":
		c.Redirect(http.StatusMovedPermanently,"https://www.bing.com/search?q="+c.PostForm("key"))

		break
	}
	
}