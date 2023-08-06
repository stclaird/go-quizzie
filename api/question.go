package api

import (
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

	defer mongo.Close(client, ctx, cancel)
    var filter, option interface{}
     
    // filter  gets all document,
    filter = bson.D{{}}
     
    //  option remove id field from all documents
    option = bson.D{{"_id", 0}}
 
    // call the query method with client, context,
    // database name, collection  name, filter and option
    // This method returns momngo.cursor and error if any.
    cursor, err := mongo.Query(client, ctx, "quizzie", "questions", filter, option)
    // handle the errors.
    if err != nil {
        panic(err)
    }
 
    var results []bson.D
     
    // to get bson object  from cursor,
    // returns error if any.
    if err := cursor.All(ctx, &results); err != nil {
        // handle the error
        panic(err)
    }

	c.JSON(http.StatusOK, gin.H{"questions": results})

}
