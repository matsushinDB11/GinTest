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
			diary := diaries.Group(":id")
			{
				diary.GET("/:id", contoroller.GetDiary)
				diary.PATCH("/:id", contoroller.PatchDiary)
				diary.DELETE("/:id", contoroller.DeleteDiary)
				diaryLikes := diary.Group("likes")
				{
					diaryLikes.POST("/:id/likes", contoroller.PostDiaryLikes)
					diaryLikes.DELETE("/:id/likes", contoroller.DeleteDiaryLikes)
				}
			}
		}
	}
	engine.Run(":3001")
}
