package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stclaird/go-quizzie/pkg/models"
)

func Home(c *gin.Context) {
	//Home Page
	c.JSON(http.StatusOK, gin.H{"response": "home"})
}

// GET /question
// Get all questions
func FindQuestion(c *gin.Context) {
	var questions []models.Question
	models.DB.Find(&questions)

	c.JSON(http.StatusOK, gin.H{"data": questions})

}
