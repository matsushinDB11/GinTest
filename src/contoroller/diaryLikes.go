package contoroller

import "github.com/gin-gonic/gin"

func PostDiaryLikes(c *gin.Context) {
	id := c.Param("id")
	// sample response
	c.JSON(200, gin.H{
		"liked": id,
	})
}

func DeleteDiaryLikes(c *gin.Context) {
	id := c.Param("id")
	// sample response
	c.JSON(200, gin.H{
		"delete like": id,
	})
}
