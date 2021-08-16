package contoroller

import "github.com/gin-gonic/gin"

func GetDiary(c *gin.Context) {
	id := c.Param("id")
	// sample response
	c.JSON(200, gin.H{
		"id":         id,
		"authorID":   1111,
		"title":      "志望企業が決められない",
		"body":       "私はどこに向かっているのか",
		"likedCount": 3,
		"liked":      false,
		"createdAt":  "2019-09-02T16:16:22+09:00",
		"updatedAt":  "2021-09-04T16:16:22+09:00",
	})
}
func PatchDiary(c *gin.Context) {
	id := c.Param("id")
	// sample response
	c.JSON(200, gin.H{
		"patch": id,
	})
}

func DeleteDiary(c *gin.Context) {
	id := c.Param("id")
	// sample response
	c.JSON(200, gin.H{
		"delete": id,
	})
}
