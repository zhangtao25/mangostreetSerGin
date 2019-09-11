package main

import "fmt"

// 类声明
type Book struct {
	Title    string
	Author   string
	DateTime string
}

// 类方法声明-传递值对象
func (b Book) B1() {
	b.Title = "Book001"
	b.Author = "ErWan"
}

// 类方法声明-传递指针对象
func (b *Book) B2() {
	b.Title = "Book002"
	b.Author = "Tinywan"
}

func main() {
	/*声明一个 Book 类型的变量 b ，并调用 B1() 和 B2()*/
	b := Book{"Def-Book", "Def-Author", "Def-DateTime"}

	fmt.Println("B1 调用前：", b.Title, b.Author, b.DateTime)
	b.B1()
	fmt.Println("B1 调用后：", b.Title)

	fmt.Println("------------------\r\n")
	fmt.Println("B2 调用前：", b.Title)
	b.B2()
	fmt.Println("B2 调用后：", b.Title)

}