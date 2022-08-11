package main


import (

	"log"
	"fmt"
	"github.com/gin-gonic/gin"

	"latihan1/handler"
	"latihan1/book"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/latihan1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&book.Book{})

	fmt.Println("Database connection berhasil")

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostHandler)

	router.Run()
}
