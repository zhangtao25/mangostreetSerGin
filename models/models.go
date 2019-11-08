package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func Setup()  {
	var err error
	db, err = gorm.Open("mysql", "root:wjyy26303@tcp(114.55.145.3:3306)/test?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}
}

func CloseDB() {
	defer db.Close()
}
