package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/stclaird/go-quizzie/api"
	model "github.com/stclaird/go-quizzie/pkg/models"
)


func CORSConfig() cors.Config {
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"http://localhost:5000"}
    corsConfig.AllowCredentials = true
    corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
    corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
    return corsConfig
}


func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	r.Use(cors.New(CORSConfig()))

	// Routes
	r.GET("", api.Home)
	r.GET("/questions", api.Questions)
	r.GET("/questions/:prefix", api.Questions)
	r.GET("/categories", api.Categories)
	r.GET("/answer/:qid/:answer", api.Answers)
	return r
}

func main() {
	// Import questions
	questions := model.InitQuestions()
	db,err := model.Open("./badger-db")
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

	// Serve frontend static files

	r.Run(":5000")
	fmt.Println("Running")
}
