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
				diary.GET("", contoroller.GetDiary)
				diary.PATCH("", contoroller.PatchDiary)
				diary.DELETE("", contoroller.DeleteDiary)
				diaryLikes := diary.Group("likes")
				{
					diaryLikes.POST("", contoroller.PostDiaryLikes)
					diaryLikes.DELETE("", contoroller.DeleteDiaryLikes)
				}
			}
		}
	}
	engine.Run(":3001")
}
