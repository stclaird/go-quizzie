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
	r.GET("/questions/:prefix", api.Questions)
	r.GET("/categories/", api.Categories)
	r.GET("/answer/:qid/:answer", api.Answers)
	return r
}

func main() {
	// Import questions
	questions := model.InitQuestions()
	db,err := model.Open("./badger-quizzie")
	if err != nil {
		log.Printf("main %s", err)
	}

	for _, question := range questions {
		err := model.InsertOneItem(question, db)
		if err != nil {
			fmt.Println(err)
		}
	}

	model.Close(db)

	// Run the server
	r := setupRouter()
	r.Run()
}
