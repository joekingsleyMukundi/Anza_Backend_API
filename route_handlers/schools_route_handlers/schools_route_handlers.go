package schoolsroutehandlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	schoolmodels "github.com/kennedy-muthaura/anzaapi/models/school_models"
	"github.com/kennedy-muthaura/anzaapi/services/mongodbapi"
	"github.com/kennedy-muthaura/anzaapi/utils/appconstants"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllSchoolsRouteHandler(c *gin.Context) {

	schools, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.SchoolsCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No Schools Available"})

	} else {
		c.JSON(200, schools)

	}

}
func AddSchoolsRouteHandler(c *gin.Context) {

	//this takes json of schoolname and checks if the school exist in db as lowercase if not add it to the collection
	var school schoolmodels.SchoolModel
	if err := c.ShouldBind(&school); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(school.ToMap(), mongodbapi.SchoolsCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func DeleteSchoolRouteHandler(c *gin.Context) {
	//todo remove thumbnail, and video ,, get lesson data from db , get the link and uuid from link, detele it in the assets folder

	var id = c.Param("id")

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.SchoolsCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "delete success"})
}
