package liveclassesroutehandlers

import (
	liveclassesmodels "github.com/Anza2022/Anza_Backend_API/models/live_classes_models"
	"github.com/Anza2022/Anza_Backend_API/services/mongodbapi"
	"github.com/Anza2022/Anza_Backend_API/utils/appconstants"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllLiveClassesRouteHandler(c *gin.Context) {

	classes, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.LiveClassesCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No live classes  found Available"})

	} else {
		c.JSON(200, classes)

	}

}
func GetTeacherLiveClassesRouteHandler(c *gin.Context) {
	var id = c.Param("id")
	classes, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.LiveClassesCollection, bson.M{"ownerId": id})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No live classes  found Available"})

	} else {
		c.JSON(200, classes)

	}
}

func AddLiveClassToDbHandler(c *gin.Context) {
	var class liveclassesmodels.LiveClassModel
	if err := c.ShouldBind(&class); err != nil {

		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(class.ToMap(), mongodbapi.LiveClassesCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdateLiveClassHandler(c *gin.Context) {
	var class liveclassesmodels.LiveClassModel
	var id = c.Param("id")

	if err := c.ShouldBind(&class); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(class.ToMap(), mongodbapi.LiveClassesCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

func DeleteLiveClassHandler(c *gin.Context) {
	//todo remove thumbnail, and video ,, get lesson data from db , get the link and uuid from link, detele it in the assets folder

	var id = c.Param("id")

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.LiveClassesCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "delete success"})
}
