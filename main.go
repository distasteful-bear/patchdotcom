package main

import (
	"distasteful-bear/patchdotcom/paths"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", paths.HomePage)

	r.GET("/ping", func(c *gin.Context) {
		c.Status(200)
	})

	r.Run()
}
