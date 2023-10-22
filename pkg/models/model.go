package models

import (
	"crypto/rand"
	"encoding/hex"
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

type CategorySubCategorys struct {
	CategoryName string   `json:"Category"`
	SubCategorys []Subcategory `json:"SubCategorys"`
}

type Subcategory struct {
	SubCategoryName string `json:"SubCategoryName"`
	URLPrefix string `json:"URLPrefix"`
}

func createQid(question Question) string{
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%s-%s-%s", question.Category, question.Subcategory, hex.EncodeToString(b))
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
			fmt.Printf("Loading %s", File.Name())
			filePath := fmt.Sprintf("questions/%s", File.Name())
			jsonFile, err := os.Open(filePath)
			if err != nil {
				log.Println("Error", err)
			}
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)
			json.Unmarshal(byteValue, &questionsObj)
			for _, question := range questionsObj {
				question.Qid = createQid(question)
				for i := range question.Answers {
					question.Answers[i].Id = strconv.Itoa(i)
				}
				allQuestions = append(allQuestions, question)
			}
		}
	}

	return allQuestions
}

