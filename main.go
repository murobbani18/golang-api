package main


import (
    http "net/http"
    "fmt"
    // "log"
    "encoding/json"
    "github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	fmt.Println(http.StatusOK)
	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postHandler)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Mufti Robbani",
		"desc": "Kalibata",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "Hello World",
	})
}


func booksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"title": "Books",
		"id" : id,
	})
}

func queryHandler(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"title": "Query",
		"id" : id,
	})
}

type BookInput struct {
	Title string 	  `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

func postHandler(c *gin.Context) {
	var bookInput BookInput

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