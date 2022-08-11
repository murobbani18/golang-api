package handler

import (
	"net/http"
    "fmt"

	"github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
	"latihan1/book"
)


func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Mufti Robbani",
		"desc": "Kalibata",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "Hello World",
	})
}


func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"title": "Books",
		"id" : id,
	})
}

func QueryHandler(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"title": "Query",
		"id" : id,
	})
}

func PostHandler(c *gin.Context) {
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"id" : bookInput.Price,
	})
}