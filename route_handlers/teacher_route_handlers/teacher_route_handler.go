package teacherroutehandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	teachersmodels "github.com/kennedy-muthaura/anzaapi/models/teachers_models"
	"github.com/kennedy-muthaura/anzaapi/services/mongodbapi"
	"github.com/kennedy-muthaura/anzaapi/utils/appconstants"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllSchemesOfWorkRouteHandler(c *gin.Context) {

	schemes, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.SchemesOfWorkCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No Schemes of work Available"})

	} else {
		c.JSON(200, schemes)

	}

}

func AddSchemeOfWorkDbHandler(c *gin.Context) {
	var scheme teachersmodels.SchemesOfWorkModel
	if err := c.ShouldBind(&scheme); err != nil {
		fmt.Println(scheme)
		fmt.Println(scheme.FileId)
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(scheme.ToMap(), mongodbapi.SchemesOfWorkCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdateSchemOfWorkHandler(c *gin.Context) {
	var scheme teachersmodels.SchemesOfWorkModel
	var id = c.Param("id")

	if err := c.ShouldBind(&scheme); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(scheme.ToMap(), mongodbapi.SchemesOfWorkCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

func DeleteSchemeOfWorkHandler(c *gin.Context) {
	//todo remove thumbnail, and video ,, get lesson data from db , get the link and uuid from link, detele it in the assets folder

	var id = c.Param("id")

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.SchemesOfWorkCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "delete success"})
}
func GetAllLessonsPlansRouteHandler(c *gin.Context) {

	plans, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.LessonPlansCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No lesson plans Available"})

	} else {
		c.JSON(200, plans)

	}

}

func AddLessonPlanToDbHandler(c *gin.Context) {
	var plan teachersmodels.LessonPlansModel
	if err := c.ShouldBind(&plan); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(plan.ToMap(), mongodbapi.LessonPlansCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdateLessonPlanHandler(c *gin.Context) {
	var plan teachersmodels.LessonPlansModel
	var id = c.Param("id")

	if err := c.ShouldBind(&plan); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(plan.ToMap(), mongodbapi.LessonPlansCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

func DeleteLessonPlanHandler(c *gin.Context) {
	//todo remove thumbnail, and video ,, get lesson data from db , get the link and uuid from link, detele it in the assets folder

	var id = c.Param("id")

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.LessonPlansCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "delete success"})
}

func GetLessonPlanFromServer(c *gin.Context) {
	fileid := c.Param("id")

	c.Status(http.StatusAccepted)

	c.File(fmt.Sprintf("assets/lessonplans/%v", fileid))
}
func GetWorkSchemeFromServer(c *gin.Context) {
	fileid := c.Param("id")

	c.Status(http.StatusAccepted)

	c.File(fmt.Sprintf("assets/workschemes/%v", fileid))
}
