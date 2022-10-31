package setbooksroutehandlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	videolessons "github.com/kennedy-muthaura/anzaapi/models/video_lessons"
	"github.com/kennedy-muthaura/anzaapi/services/mongodbapi"
	"github.com/kennedy-muthaura/anzaapi/utils/appconstants"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllLSetBookEpisodesRouteHandler(c *gin.Context) {

	episodes, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.SetBooksEpisodesCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No setbook episodes found Available"})

	} else {
		c.JSON(200, episodes)

	}

}

func AddSetBookEpisodeToDbHandler(c *gin.Context) {
	var episode videolessons.SetBookEpisode
	if err := c.ShouldBind(&episode); err != nil {

		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(episode.ToMap(), mongodbapi.SetBooksEpisodesCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdateSetBookEpisodeHandler(c *gin.Context) {
	var episode videolessons.SetBookEpisode
	var id = c.Param("id")

	if err := c.ShouldBind(&episode); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(episode.ToMap(), mongodbapi.SetBooksEpisodesCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

func DeleteSetBookEpisodeHandler(c *gin.Context) {
	//todo remove thumbnail, and video ,, get lesson data from db , get the link and uuid from link, detele it in the assets folder

	var id = c.Param("id")

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.SetBooksEpisodesCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "delete success"})
}

func GetSetBookEpisodeFromServer(c *gin.Context) {
	fileid := c.Param("id")

	c.Status(http.StatusAccepted)

	c.File(fmt.Sprintf("assets/setbooks/%v", fileid))
}

func GetAllSetbookFilesInTheServer(c *gin.Context) {
	files, err := ioutil.ReadDir("./assets/setbooks")
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "files not  found"})
		return

	}
	allfiles := []string{}

	for _, f := range files {
		if !f.IsDir() {
			allfiles = append(allfiles, f.Name())
		}
	}
	c.JSON(200, map[string]interface{}{"files": allfiles})

}
