package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stclaird/quizzie/pkg/models"
)

// GET /question
// Get all questions
func FindQuestion(c *gin.Context) {
	var questions []models.Question
	models.DB.Find(&questions)

	c.JSON(http.StatusOK, gin.H{"data": questions})

}
