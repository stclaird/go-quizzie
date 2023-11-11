package api

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	model "github.com/stclaird/go-quizzie/pkg/models"
)

//Send Answer
func Answers(c *gin.Context) {
	db,err := model.Open("./badger-db")
	if err != nil {
		log.Printf("error in func Answers %s", err)
	}
	questionId := c.Param("qid")
	// fmt.Println("QUID %s",questionId)
	submittedAnswer := c.Param("answer")
	question, err := model.GetItem(questionId, db)
	model.Close(db)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(question.Answers)
	isCorrect, answer := checkAnswer(question, submittedAnswer)

	response := model.AnswerResponse{
		IsCorrect : isCorrect,
		ActualAnswer: answer,
	}

	c.JSON(http.StatusOK, response)
}

//compare the real answer with the user submitted answer
func checkAnswer(question model.Question, submittedAnswer string) (bool, []model.Answer) {
	var answers []string
	var correctAnswers []model.Answer

	for _,v := range question.Answers{
		if v.IsCorrect == true {
			correctAnswer := model.Answer{
				Id : v.Id,
				Answer: v.Text,
			}
			correctAnswers = append(correctAnswers, correctAnswer)
			answers = append(answers, v.Id)
		}
	}
	fmt.Printf("answers %v\n", answers)

	sort.Strings(answers)
	submittedAnswerArr := strings.Split(submittedAnswer, "")
	sort.Strings(submittedAnswerArr)

	answersStr := strings.Join(answers, "")
	submittedAnswerStr := strings.Join(submittedAnswerArr, "")

	fmt.Printf("Final: %s,%s", answersStr, submittedAnswerStr)

	if answersStr != submittedAnswerStr {
		return false, correctAnswers
	}

	return true, correctAnswers
}