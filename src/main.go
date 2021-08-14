package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/nyan", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "にゃーん",
		})
	})
	r.Run()
}
