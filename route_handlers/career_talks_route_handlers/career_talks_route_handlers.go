package careertalksroutehandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	talksmodels "github.com/kennedy-muthaura/anzaapi/models/talks_models"
	"github.com/kennedy-muthaura/anzaapi/services/mongodbapi"
	"github.com/kennedy-muthaura/anzaapi/utils/appconstants"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCareerTalksRouteHandler(c *gin.Context) {

	quizes, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.CareerTalksCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No Career Talks Available"})

	} else {
		c.JSON(200, quizes)

	}

}

func AddCareerTalkToDbHandler(c *gin.Context) {
	var talk talksmodels.CareerTalkModel
	if err := c.ShouldBind(&talk); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(talk.ToMap(), mongodbapi.CareerTalksCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdateCareerTalkHandler(c *gin.Context) {
	var talk talksmodels.CareerTalkModel
	var id = c.Param("id")

	if err := c.ShouldBind(&talk); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(talk.ToMap(), mongodbapi.CareerTalksCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

func DeleteCareerTalkHandler(c *gin.Context) {
	//todo remove thumbnail, and video ,, get lesson data from db , get the link and uuid from link, detele it in the assets folder

	var id = c.Param("id")

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.CareerTalksCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "delete success"})
}

func GetACareerFromServer(c *gin.Context) {
	videoName := c.Param("id")

	c.Status(http.StatusAccepted)

	c.File(fmt.Sprintf("assets/careertalks/videos/%v", videoName))
}

func GetCareerThumbnailFromServer(c *gin.Context) {
	fmt.Println("getting thumbanail from server")
	thumbnailName := c.Param("id")

	c.File(fmt.Sprintf("assets/careertalks/thumbnails/%v", thumbnailName))

}
