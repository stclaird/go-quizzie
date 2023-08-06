package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stclaird/go-quizzie/api"
	"github.com/stclaird/go-quizzie/pkg/models"
	"gorm.io/gorm"
)

func initQuestions(db *gorm.DB) {
	//import questions from a json file into database
	var questionsObj []models.ImportQuestion

	jsonFile, err := os.Open("questions.json")
	if err != nil {
		log.Println("Error", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &questionsObj)

	for _, q := range questionsObj {
		newQuestion := models.Question{
			Text:        q.Text,
			Type:        q.Type,
			Category:    q.Category,
			Subcategory: q.Subcategory,
		}
		insertedQ := models.DB.FirstOrCreate(&newQuestion, models.Question{Text: q.Text})

		fmt.Println(insertedQ)
		for _, a := range q.Answers {
			newAnswer := models.Answer{
				Text:       a.Text,
				IsCorrect:  a.IsCorrect,
				QuestionID: newQuestion.ID,
			}
			models.DB.FirstOrCreate(&newAnswer, models.Answer{Text: a.Text})
		}
	}

}

func main() {

	models.ConnectDatabase()

	initQuestions(models.DB)

	r := gin.Default()
	// Routes
	r.GET("/", api.Home)
	r.GET("/question", api.FindQuestion)
	// r.GET("/question/:id", api.FindBook)
	// r.POST("/question", api.CreateBook)
	// r.PATCH("/question/:id", api.UpdateBook)
	// r.DELETE("/question/:id", api.DeleteBook)

	// Run the server
	r.Run()

}
