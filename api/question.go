package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/stclaird/go-quizzie/pkg/models"
)

type Answer struct {
	Qid string
	Answer string
}

func contains(s []model.CategorySubCategorys, e model.CategorySubCategorys) (bool, int) {
    for k, v := range s {
        if v.CategoryName == e.CategoryName {
            return true, k
        }
    }
    return false, -1
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
	var catSubCategories []model.CategorySubCategorys

	All, _ := model.GetAllItems(db)
	model.Close(db)

	for _, v := range All {
		var catSubCat model.CategorySubCategorys
		var subCategory model.Subcategory
		subCategory.SubCategoryName = v.Subcategory
		subCategory.URLPrefix = fmt.Sprintf("%s-%s", v.Category, v.Subcategory)

		catSubCat.CategoryName = v.Category

		exists, idx := contains(catSubCategories, catSubCat)
		if  exists {
			catSubCategories[idx].SubCategorys = append(catSubCategories[idx].SubCategorys, subCategory )
			break
		} else {
			catSubCat.SubCategorys = append(catSubCat.SubCategorys,subCategory)
		}
		catSubCategories = append(catSubCategories, catSubCat)
	}

	c.JSON(http.StatusOK, &catSubCategories)
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

	model.Close(db)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, response)
}
