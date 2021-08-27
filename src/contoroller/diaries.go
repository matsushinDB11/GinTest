package contoroller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

//type Diaries struct {
//	Intractor diaries.Interactor
//}

func GetDiariesList(c *gin.Context) {
	sort := c.DefaultQuery("sort", "createdAt ")
	_ = sort
	// sample response
	res := map[string]interface{}{}
	resJSON := `{"diaries":[{"id":1,"authorID":1111,"title":"志望企業が決められない","body":"私はどこに向かっているのか","likedCount":3,"liked": false,"createdAt":"2019-09-02T16:16:22+09:00","updatedAt": "2021-09-04T16:16:22+09:00"},{"id":2,"authorID":2212,"title":"最終面接落ちたんご","body":"どないすればええんや","likedCount":6,"liked": true,"createdAt":"2021-09-04T16:20:12+09:00","updatedAt": "2021-09-12T16:16:22+09:00"},{"id":3,"authorID":7656,"title":"内定でたぞ","body":"よっしゃー","likedCount":5,"liked": false,"createdAt":"2021-09-02T17:09:11+09:00","updatedAt": "2019-09-23T16:16:22+09:00"}]}`
	_ = json.Unmarshal([]byte(resJSON), &res)
	c.JSON(200, res)
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
