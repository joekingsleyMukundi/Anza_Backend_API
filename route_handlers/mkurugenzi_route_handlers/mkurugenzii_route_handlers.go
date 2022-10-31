package mkurugenziroutehandlers

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

func GetAllLMkurugenziEpisodesRouteHandler(c *gin.Context) {

	episodes, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.MkurugenziEpisodesCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No Mkurugenzi episodes found Available"})

	} else {
		c.JSON(200, episodes)

	}

}

func AddMkurugenziEpisodeToDbHandler(c *gin.Context) {
	var episode videolessons.MkurugenziEpisode
	if err := c.ShouldBind(&episode); err != nil {

		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(episode.ToMap(), mongodbapi.MkurugenziEpisodesCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdateMkurugenziEpisodeHandler(c *gin.Context) {
	var episode videolessons.MkurugenziEpisode
	var id = c.Param("id")

	if err := c.ShouldBind(&episode); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(episode.ToMap(), mongodbapi.MkurugenziEpisodesCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

func DeleteMkurugenziEpisodeHandler(c *gin.Context) {
	//todo remove thumbnail, and video ,, get lesson data from db , get the link and uuid from link, detele it in the assets folder

	var id = c.Param("id")

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.MkurugenziEpisodesCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "delete success"})
}

func GetMkurugenziEpisodeFromServer(c *gin.Context) {
	fileid := c.Param("id")

	c.Status(http.StatusAccepted)

	c.File(fmt.Sprintf("assets/mkurugenzi/videos/%v", fileid))
}

func GetMkurugenziThumbnailFromServer(c *gin.Context) {
	thumbnailName := c.Param("id")
	c.File(fmt.Sprintf("assets/mkurugenzi/thumbnails/%v", thumbnailName))
}
func GetAllMkurugenziFilesInTheServer(c *gin.Context) {
	files, err := ioutil.ReadDir("./assets/mkurugenzi/videos")
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
