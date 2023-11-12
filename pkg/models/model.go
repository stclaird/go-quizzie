package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Question struct {
	Qid			string `json:"qid"`
	Text        string `json:"text"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	DateAdded   string
	Answers     []struct {
		Id        string
		Text      string `json:"text"`
		IsCorrect bool   `json:"iscorrect"`
	} `json:"answers"`
}

type QuestionNoAnswer struct {
	Qid			string `json:"qid"`
	Text        string `json:"text"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	DateAdded   string
	Answers     []struct {
		Id        string
		Text      string `json:"text"`
	} `json:"answers"`
}

type Category struct {
	Id string	`json:"id"`
	CategoryName string   `json:"Category"`
	SubCategories []Subcategory `json:"SubCategories"`
}

type Subcategory struct {
	SubCategoryName string `json:"SubCategoryName"`
	URLPrefix string `json:"URLPrefix"`
}

type AnswerResponse struct {
	IsCorrect bool
	ActualAnswer []Answer
}

type Answer struct {
	Id string
	Answer string
}

func createQid(question Question, k int) string{
	//Create qid (question ID)
	return fmt.Sprintf("%s-%s-%s", question.Category, question.Subcategory, strconv.Itoa(k))
}

func InitQuestions() (allQuestions []Question) {
	//import questions from a json file into database
	//returns a slice of question stuct types

	files, err := ioutil.ReadDir("questions/")
	if err != nil {
		log.Fatal(err)
	}

	for _, File := range files {
		fileExtension := filepath.Ext(File.Name())
		if fileExtension == ".json" {
			var questionsObj []Question
			fmt.Printf("Loading %s\n", File.Name())
			filePath := fmt.Sprintf("questions/%s", File.Name())
			jsonFile, err := os.Open(filePath)
			if err != nil {
				log.Println("Error", err)
			}
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)
			json.Unmarshal(byteValue, &questionsObj)
			for k, question := range questionsObj {
				question.Qid = createQid(question,k)
				for i := range question.Answers {
					question.Answers[i].Id = strconv.Itoa(i)
				}
				allQuestions = append(allQuestions, question)
			}
		}
	}
	return allQuestions
}

