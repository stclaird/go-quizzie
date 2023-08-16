package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stclaird/go-quizzie/api"
	mongomodel "github.com/stclaird/go-quizzie/pkg/models"
)

func main() {

	client, ctx, cancel, err := mongomodel.Connect("mongodb://mongoadmin:mongoadmin@mongo:27017")
	if err != nil {
		panic(err)
	}

	defer mongomodel.Close(client, ctx, cancel)

	db := client.Database("quizzie")
	qCollection := db.Collection("questions")
	if err = qCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	questions := mongomodel.InitQuestions()

	for _, doc := range questions {
		_, err := mongomodel.InsertOne(client, ctx, "quizzie", "questions", doc)
		if err != nil {
			fmt.Println(err)
		}
	}

	r := gin.Default()
	// Routes
	r.GET("/", api.Home)
	r.GET("/questions", api.Questions)
	r.GET("/categorys/", api.Categorys)
	r.GET("/category/:category", api.Category)

	// Run the server
	r.Run()

}
