package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stclaird/go-quizzie/api"
)

func main() {
	router := gin.Default()

	r := gin.Default()

	// Routes
	r.GET("/question", api.FindBooks)
	// r.GET("/question/:id", api.FindBook)
	// r.POST("/question", api.CreateBook)
	// r.PATCH("/question/:id", api.UpdateBook)
	// r.DELETE("/question/:id", api.DeleteBook)

	// Run the server
	r.Run()

}
