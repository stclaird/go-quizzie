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

func InitQuestions() (allQuestions []Question) {
	//import questions from a json file into database

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
				for i := range question.Answers {
					question.Answers[i].Id = strconv.Itoa(i)
				}
				allQuestions = append(allQuestions, question)
			}
		}
	}

	return allQuestions
}
