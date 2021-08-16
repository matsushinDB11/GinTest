package main

import (
	"gin-api/contoroller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.String(200, "This is gin_test")
	})
	serviceEngine := engine.Group("/service")
	{
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
