package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	model "github.com/stclaird/go-quizzie/pkg/models"
)

type Answer struct {
	Qid string
	Answer string
}

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

func Home(c *gin.Context) {
	//Home Page
	c.JSON(http.StatusOK, gin.H{"response": "home"})
}

func Categories(c *gin.Context) {
	db,err := model.Open("./badger-quizzie")
	if err != nil {
		log.Printf("func Questions %s\n", err)
	}

	All, _ := model.GetAllItems(db)


	Categories := make(map[string]*model.Category)
	subCategories := make(map[string]model.Subcategory)

	//Populate a map of all Categories and a second map of all subcategories
	for k, v := range All {
		var category model.Category
		category.Id = strconv.Itoa(k)
		category.CategoryName = v.Category
		Categories[ v.Category] = &category

		var subCategory model.Subcategory
		subCategory.SubCategoryName = v.Subcategory
		subCategory.URLPrefix = fmt.Sprintf("%s-%s", v.Category, v.Subcategory)
		subCategories[subCategory.URLPrefix] = subCategory

	}
	//Apply all subcategories to appropriate category
	for _,v := range subCategories {
		splt := strings.Split(v.URLPrefix, "-")
		cat := splt[0]
		for _, value := range Categories {
			if value.CategoryName == cat {
				Categories[cat].SubCategories = append(Categories[cat].SubCategories, v)
			}
		}
	}

	var response []*model.Category
	for _,v := range Categories {
		response = append(response, v)
	}

	c.JSON(http.StatusOK, &response)
}

//Retrieve Questions from specific category "prefix"
func Questions(c *gin.Context) {
	prefix := c.Param("prefix")
	db,err := model.Open("./badger-quizzie")
	if err != nil {
		log.Printf("func Questions %s", err)
	}

	var response []model.QuestionNoAnswer
	response, err = model.GetItemsbyPrefix(prefix, db)


	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, response)
}
