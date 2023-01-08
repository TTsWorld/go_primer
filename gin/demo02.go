package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func panicrecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				println("xxxxx")
				c.JSON(http.StatusInternalServerError, "server error")
				c.Request.Context().Done()
			}
		}()
		c.Next()
	}

}

func main() {
	r := gin.Default()
	//	r.Use(logBody())
	r.Use(panicrecover())

	r.GET("/", func(c *gin.Context) {
		a := []int{}
		println(a[0])
		c.JSON(200, "首页")
	})

	adminGroup := r.Group("/admin")
	adminGroup.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))

	adminGroup.POST("/index", func(c *gin.Context) {
		byteBody, _ := ioutil.ReadAll(c.Request.Body)
		println()
		println(string(byteBody))
		c.JSON(200, "后台首页")
	})

	r.Run(":8080")
}
