package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stclaird/go-quizzie/api"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
    r := gin.Default()
	r.GET("/ping", api.Ping)
	r.GET("/questions", api.Questions)
	r.GET("/questions/:prefix", api.Questions)
	r.GET("/categories", api.Categories)
	r.GET("/answer/:qid/:answer", api.Answers)
    return r
}

func TestPingRoute(t *testing.T) {
	router := SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}


func TestCategoriesRoute(t *testing.T) {
	router := SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/categories", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Category")
}