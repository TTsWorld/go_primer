package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func logBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyBytes []byte // 我们需要的body内容

		// 从原有Request.Body读取
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil || len(bodyBytes) > 2048 {
			c.JSON(400, "args error")
			return
		}

		// 新建缓冲区并替换原有Request.body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// 当前函数可以使用body内容
		a := bodyBytes
		fmt.Printf("%s", a)

		c.Next()
	}

}

func main() {
	r := gin.Default()
	r.Use(logBody())

	r.POST("/", func(c *gin.Context) {
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
