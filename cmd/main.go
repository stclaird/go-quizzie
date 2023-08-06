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
	mongo "github.com/stclaird/go-quizzie/pkg/models"
)

func initQuestions() (questionsObj []models.Question) {
	//import questions from a json file into database

	jsonFile, err := os.Open("questions.json")
	if err != nil {
		log.Println("Error", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &questionsObj)

	return questionsObj

}

func main() {

	client, ctx, cancel, err := mongo.Connect("mongodb://mongoadmin:mongoadmin@mongo:27017")
    if err != nil {
        panic(err)
    }

	defer mongo.Close(client, ctx, cancel)

	questions := initQuestions()

	for _, doc := range questions{
		fmt.Println(doc)
		result, err := mongo.InsertOne(client, ctx, "quizzie", "questions", doc)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.InsertedID)

	}

	r := gin.Default()
	// Routes
	r.GET("/", api.Home)
	r.GET("/questions", api.Questions)
	// r.GET("/question/:id", api.FindBook)
	// r.POST("/question", api.CreateBook)
	// r.PATCH("/question/:id", api.UpdateBook)
	// r.DELETE("/question/:id", api.DeleteBook)

	// Run the server
	r.Run()

}
