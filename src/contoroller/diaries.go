package contoroller

import "github.com/gin-gonic/gin"

func GetDiariesList(c *gin.Context) {
	sort := c.DefaultQuery("sort", "createdAt ")
	// sample response
	c.JSON(200, gin.H{
		"sample": sort,
	})
}

func PostDiary(c *gin.Context) {
	title := c.Query("title")
	body := c.Query("body")
	userid := c.Query("author_id")
	// sample response
	c.JSON(201, gin.H{
		"title":  title,
		"body":   body,
		"userID": userid,
	})
}
