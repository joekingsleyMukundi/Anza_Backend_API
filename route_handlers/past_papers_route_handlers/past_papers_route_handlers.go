package pastpapersroutehandlers

import (
	"fmt"
	"net/http"
	"strings"

	pastpapersmodels "github.com/Anza2022/Anza_Backend_API/models/past_papers_models"
	"github.com/Anza2022/Anza_Backend_API/services/mongodbapi"
	"github.com/Anza2022/Anza_Backend_API/utils/appconstants"
	helperfunctions "github.com/Anza2022/Anza_Backend_API/utils/helper_functions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllPastPapersRouteHandlers(c *gin.Context) {

	papers, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.PastPapersCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No Past Papers Available"})

	} else {
		c.JSON(200, papers)

	}

}
func GetAllTeachersPastPapersRouteHandlers(c *gin.Context) {
	var id = c.Param("id")
	papers, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.PastPapersCollection, bson.M{"ownerId": id})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No Past Papers Available"})

	} else {
		c.JSON(200, papers)

	}

}

func AddPastPaperToDbHandler(c *gin.Context) {
	var paper pastpapersmodels.PastPaperModel
	if err := c.ShouldBind(&paper); err != nil {

		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(paper.ToMap(), mongodbapi.PastPapersCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdatePastPaperHandler(c *gin.Context) {
	var paper pastpapersmodels.PastPaperModel
	var id = c.Param("id")

	if err := c.ShouldBind(&paper); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(paper.ToMap(), mongodbapi.PastPapersCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

func DeletePastPaperHandler(c *gin.Context) {

	var id = c.Param("id")
	paper, err := helperfunctions.GetPastPaperModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})
	if err == nil {
		splitfilename := strings.Split(paper.PaperUrl, "/")
		filename := splitfilename[len(splitfilename)-1]
		helperfunctions.DeleteFileInServer(filename)
		if paper.MarkingSchemeUrl != "" {
			splitmsname := strings.Split(paper.PaperUrl, "/")
			msname := splitmsname[len(splitmsname)-1]
			helperfunctions.DeleteFileInServer(msname)
		}

	}

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.PastPapersCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}

	c.JSON(200, bson.M{"message": "delete success"})
}

func GetPastPaperFromServer(c *gin.Context) {
	paperId := c.Param("paperid")

	c.Status(http.StatusAccepted)

	c.File(fmt.Sprintf("assets/pastpapers/%v", paperId))
}
func GetPdfNotesFromServer(c *gin.Context) {
	paperId := c.Param("paperid")

	c.Status(http.StatusAccepted)

	c.File(fmt.Sprintf("assets/pdfnotes/%v", paperId))
}
