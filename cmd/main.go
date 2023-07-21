package main

import (
	"github.com/stclaird/gin-bookstore/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	r := gin.Default()

	// Routes
	r.GET("/question", controllers.FindBooks)
	r.GET("/question/:id", controllers.FindBook)
	r.POST("/question", controllers.CreateBook)
	r.PATCH("/question/:id", controllers.UpdateBook)
	r.DELETE("/question/:id", controllers.DeleteBook)

	// Run the server
	r.Run()

}
