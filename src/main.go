package main

import (
	"gin-api/contoroller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	serviceEngine := engine.Group("/service")
	{
		serviceEngine.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "This is gin_test",
			})
		})
		diaries := serviceEngine.Group("diaries")
		{
			diaries.GET("/", contoroller.GetDiariesList)
			diaries.POST("/", contoroller.PostDiary)
			diaries.GET("/:id", contoroller.GetDiary)
			diaries.PATCH("/:id", contoroller.PatchDiary)
			diaries.DELETE("/:id", contoroller.DeleteDiary)
			diaries.POST("/:id/likes", contoroller.PostDiaryLikes)
			diaries.DELETE("/:id/likes", contoroller.DeleteDiaryLikes)
		}
	}
	engine.Run(":3001")
}
