package userroutehandlers

import (
	"fmt"
	"mime/multipart"
	"os"

	usermodels "github.com/Anza2022/Anza_Backend_API/models/user_models"
	"github.com/Anza2022/Anza_Backend_API/services/mongodbapi"
	"github.com/Anza2022/Anza_Backend_API/utils/appconstants"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func GetOneUserHandler(c *gin.Context) {
	id := c.Param("userId")

	u, err := mongodbapi.RetrieveOneDocumentInACollectionWithId(mongodbapi.UsersCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "user not found"})
		return
	}
	bounsData, err := mongodbapi.RetrieveOneDocumentInACollection(mongodbapi.ReferBonusesDataCollection, map[string]interface{}{"userId": id})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to refresh user"})
		return
	}

	c.JSON(200, map[string]interface{}{"userData": u, "bonusData": bounsData})

}

func GetAllUsersHandler(c *gin.Context) {

	users, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.UsersCollection, bson.M{})

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, gin.H{
			"message": "Failed to load the Users",
		})

		return
	}
	// for i := 0; i < 600; i++ {
	// 	games = append(games, games[4])
	// }

	c.JSON(200, users)
}

func UpdateUserDetailsHandler(c *gin.Context) {
	//must be admin or have a token with that users email.
	var updatedUser usermodels.UserModel

	err := c.ShouldBind(&updatedUser)
	id := c.Param("id")

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, gin.H{
			"message": "All User Fields for updating Were not Provided/correct",
		})

		return
	}

	updatedId, err := mongodbapi.UpdateADocInCollection(updatedUser.ToMap(), mongodbapi.UsersCollection, id)

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, gin.H{
			"message": "Unable to update the User",
		})

		return
	}

	c.JSON(200, gin.H{
		"_id": updatedId,
	})
}
func UpdateUserStatsHandler(c *gin.Context) {
	//must be admin or have a token with that users email.
	var updatedStats usermodels.UserStats

	err := c.ShouldBind(&updatedStats)
	id := c.Param("id")

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, gin.H{
			"message": "All User Fields for updating Were not Provided/correct",
		})

		return
	}

	updatedId, err := mongodbapi.UpdateADocInCollection(updatedStats.ToMap(), mongodbapi.UserStatsCollection, id)

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, gin.H{
			"message": "Unable to update the User",
		})

		return
	}

	c.JSON(200, gin.H{
		"_id": updatedId,
	})
}

func DeleteOneUserHandler(c *gin.Context) {

}

//files handlers
// var productionUrl = "http://192.168.137.195:8085/anzaapi"

var productionUrl = "https://api.thesigurd.com/anzaapi"

type ClientPhoto struct {
	ClientPhoto *multipart.FileHeader `form:"photo" binding:"required"`
	Extension   string                `form:"extension" binding:"required"`
}

// size 540 *  360 max 150 kbs
func AddAHousePhoto(c *gin.Context) {
	var photoType = c.Query("type")

	var form ClientPhoto
	// in this case proper binding will be automatically selected
	err := c.ShouldBind(&form)

	if err != nil {

		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "cannot bind the thumbnail form"})
		return
	}

	imageId := uuid.New()

	filename := fmt.Sprintf("%v.%v", imageId, form.Extension)
	savePath := "photos/"
	if photoType == "housePhoto" {
		savePath = fmt.Sprintf("%v%v", savePath, "housePhotos")

	}
	if photoType == "tenantPhoto" {
		savePath = fmt.Sprintf("%v%v", savePath, "tenantPhotos")

	}
	if photoType == "tenantDocPhoto" {
		savePath = fmt.Sprintf("%v%v", savePath, "tenantDocPhotos")

	}
	if photoType == "profilePhoto" {
		savePath = fmt.Sprintf("%v%v", savePath, "profilePhotos")

	}

	err = c.SaveUploadedFile(form.ClientPhoto, fmt.Sprintf("%v/%v", savePath, filename))
	if err != nil {

		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to save photo"})
		return
	}
	// db.Save(&form)
	c.JSON(201, map[string]string{"url": fmt.Sprintf("%v/file/%vs/%v", productionUrl, photoType, filename)})
}

func GetAProfilePhotoFromServer(c *gin.Context) {
	filename := c.Param("filename")

	c.File(fmt.Sprintf("photos/profilePhotos/%v", filename))
	// c.File("assets/videoclasses/thumbnails/greetings.png")

}

func DeletePhotoFromServer(c *gin.Context) {
	var filename = c.Param("filename")

	picdirs := [3]string{"housePhotos", "tenantPhotos", "profilePhotos"}
	for _, v := range picdirs {
		_ = os.Remove(fmt.Sprintf("photos/%v/%v", v, filename))

	}

	c.JSON(200, "")

}
