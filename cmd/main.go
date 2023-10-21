package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stclaird/go-quizzie/api"
	model "github.com/stclaird/go-quizzie/pkg/models"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/", api.Home)
	r.GET("/questions", api.Questions)
	// r.GET("/questions/:subcategory", api.Questions)
	r.GET("/categories/", api.Categories)
	// r.POST("/answer", api.Answers)
	return r
}

func main() {

	// Import questions from JSON files
	questions := model.InitQuestions()

	db,err := model.Open("./badger-quizzie")
	if err != nil {
		log.Printf("main %s", err)
	}

	for _, question := range questions {
		err := model.InsertOne(question, db)
		if err != nil {
			fmt.Println(err)
		}
	}

	model.Close(db)

	// Run the server
	r := setupRouter()
	r.Run()

}
