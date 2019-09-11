package notes

import "github.com/gin-gonic/gin"

func GetAll(c *gin.Context) {
	//res := models.GetNotes()
	c.JSON(200, gin.H{
		"notes": "res",
	})
	return
}