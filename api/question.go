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

//Retrieve Questions from specific category "prefix"
func Questions(c *gin.Context) {

	prefix := c.Param("prefix")

	db,err := model.Open("./badger-quizzie")
	if err != nil {
		log.Printf("func Questions %s", err)
	}
	var results []model.Question

	results, err = model.GetItemsbyPrefix(prefix, db)
	model.Close(db)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, results)
}

//Retrieve Categories
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


//Retrieve Answer to Question
// func Answers(c *gin.Context) {
// 	client, ctx, cancel, err := mongo.Connect("mongodb://mongoadmin:mongoadmin@mongo:27017")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer mongo.Close(client, ctx, cancel)

// 	var req Answer
// 	c.BindJSON(&req)
// 	fmt.Println("Posted", req)

// 	var filter, option interface{}
// 	option = bson.D{
// 		{"_id", 0},
// 	}

// 	cursor, err := mongo.Query(client, ctx, "quizzie", "questions", filter, option)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var question mongo.Question

// 	if err := cursor.All(ctx, &question); err != nil {
// 		// handle the error
// 		panic(err)
// 	}

// 	c.JSON(200, req)

// }