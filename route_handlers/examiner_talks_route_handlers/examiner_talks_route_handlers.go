package examinertalksroutehandlers

import (
	"fmt"

	talksmodels "github.com/Anza2022/Anza_Backend_API/models/talks_models"
	"github.com/Anza2022/Anza_Backend_API/services/mongodbapi"
	"github.com/Anza2022/Anza_Backend_API/utils/appconstants"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllExaminerTalksRouteHandler(c *gin.Context) {

	talks, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.ExaminerTalksCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No Examiner Talks Available"})

	} else {
		c.JSON(200, talks)

	}

}

func AddExaminierTalkToDbHandler(c *gin.Context) {
	var talk talksmodels.ExaminerTalkModel
	if err := c.ShouldBind(&talk); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(talk.ToMap(), mongodbapi.ExaminerTalksCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdateExaminerTalkHandler(c *gin.Context) {
	var talk talksmodels.ExaminerTalkModel
	var id = c.Param("id")

	if err := c.ShouldBind(&talk); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(talk.ToMap(), mongodbapi.ExaminerTalksCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

func DeleteExaminerTalkHandler(c *gin.Context) {
	//todo remove thumbnail, and video ,, get lesson data from db , get the link and uuid from link, detele it in the assets folder

	var id = c.Param("id")

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.ExaminerTalksCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "delete success"})
}
