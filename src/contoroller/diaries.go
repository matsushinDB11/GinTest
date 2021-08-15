package contoroller

import "github.com/gin-gonic/gin"

func GetDiariesList(c *gin.Context) {
	sort := c.Param("sort")
	// sample response
	c.JSON(200, gin.H{
		"sample": sort,
	})
}

func PostDiary(c *gin.Context) {
	title := c.Param("title")
	body := c.Param("body")
	userid := c.Param("author_id")
	// sample response
	c.JSON(201, gin.H{
		"title":  title,
		"body":   body,
		"userID": userid,
	})
}
