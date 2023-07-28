package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stclaird/go-quizzie/api"
	"github.com/stclaird/go-quizzie/pkg/models"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

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
