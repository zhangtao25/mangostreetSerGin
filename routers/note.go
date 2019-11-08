package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func C()  {
	var err error
	db, err = gorm.Open("mysql", "root:wjyy26303@tcp(114.55.145.3:3306)/test?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
}



func GetNote(c *gin.Context) {


	title := c.Param("title")

	type Note struct {
		Id   int
		Title string
	}
	var notes []Note
	db.Where("title = ?", title).Find(&notes)
	fmt.Println(title,notes)

	c.JSON(200,gin.H{
		"data": notes,
	})
}