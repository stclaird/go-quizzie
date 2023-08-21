package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mongo "github.com/stclaird/go-quizzie/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Home(c *gin.Context) {
	//Home Page
	c.JSON(http.StatusOK, gin.H{"response": "home"})
}

// GET /question
// Get all questions
func Questions(c *gin.Context) {
	client, ctx, cancel, err := mongo.Connect("mongodb://mongoadmin:mongoadmin@mongo:27017")
	if err != nil {
		panic(err)
	}

	var filter, option interface{}
	option = bson.D{{"_id", 0}}

	subcategory := c.Param("subcategory")
	if subcategory == "" {
		filter = bson.M{"subcategory": subcategory}
	} else {
		filter = bson.D{{}}
	}

	defer mongo.Close(client, ctx, cancel)

	cursor, err := mongo.Query(client, ctx, "quizzie", "questions", filter, option)
	if err != nil {
		panic(err)
	}

	var results []bson.D
	if err := cursor.All(ctx, &results); err != nil {
		// handle the error
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"questions": results})
}

type CategorySubCategorys struct {
	CategoryName string   `json:"Category"`
	SubCategorys []string `json:"SubCategorys"`
}

func Categorys(c *gin.Context) {
	client, ctx, cancel, err := mongo.Connect("mongodb://mongoadmin:mongoadmin@mongo:27017")
	if err != nil {
		panic(err)
	}

	defer mongo.Close(client, ctx, cancel)
	filter := bson.D{{}}
	collection := client.Database("quizzie").Collection("questions")
	categories, err := collection.Distinct(ctx, "category", filter)
	if err != nil {
		panic(err)
	}

	var catSubCats []CategorySubCategorys

	var catfilter interface{}
	for _, cat := range categories {
		catStr := fmt.Sprintf("%v", cat)
		catfilter = bson.M{"category": catStr}
		subCatsResp, err := collection.Distinct(ctx, "subcategory", catfilter)
		if err != nil {
			panic(err)
		}

		var subCatsStr []string
		for _, x := range subCatsResp {
			subCatsStr = append(subCatsStr, fmt.Sprintf("%v", x))
		}

		catSubCat := CategorySubCategorys{
			CategoryName: catStr,
			SubCategorys: subCatsStr,
		}
		catSubCats = append(catSubCats, catSubCat)
	}

	c.JSON(http.StatusOK, &catSubCats)
}
