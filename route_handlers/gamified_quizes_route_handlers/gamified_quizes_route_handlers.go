package gamifiedquizesroutehandlers

import (
	"fmt"
	"strconv"

	gamifiedquizes "github.com/Anza2022/Anza_Backend_API/models/gamified_quizes"
	"github.com/Anza2022/Anza_Backend_API/services/mongodbapi"
	"github.com/Anza2022/Anza_Backend_API/utils/appconstants"
	helperfunctions "github.com/Anza2022/Anza_Backend_API/utils/helper_functions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllGamifiedQuestionsRouteHandler(c *gin.Context) {

	quizes, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.GamifiedQuizesCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No Quizes Available"})

	} else {
		c.JSON(200, quizes)

	}

}
func GetTeachersGamifiedQuestionsRouteHandler(c *gin.Context) {
	var id = c.Param("id")
	quizes, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.GamifiedQuizesCollection, bson.M{"ownerId": id})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No Quizes Available"})

	} else {
		c.JSON(200, quizes)

	}

}

func AddGamifiedQuizesToDbHandler(c *gin.Context) {
	var test gamifiedquizes.GamifiedQuizesModel
	if err := c.ShouldBind(&test); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(test.ToMap(), mongodbapi.GamifiedQuizesCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdateGamifiedQuizTestHandler(c *gin.Context) {
	var test gamifiedquizes.GamifiedQuizesModel
	var id = c.Param("id")

	if err := c.ShouldBind(&test); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(test.ToMap(), mongodbapi.GamifiedQuizesCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}
func PublishGamifiedTestForStudentsHandler(c *gin.Context) {
	var id = c.Param("id")
	test, err := helperfunctions.GetGamifiedQuizModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Test not found in database"})
		return

	}
	test.IsPublished = true

	_, err = mongodbapi.UpdateADocInCollection(test.ToMap(), mongodbapi.GamifiedQuizesCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})
}
func UnPublishGamifiedTestForStudentsHandler(c *gin.Context) {
	var id = c.Param("id")
	test, err := helperfunctions.GetGamifiedQuizModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Test not found in database"})
		return

	}
	test.IsPublished = false
	_, err = mongodbapi.UpdateADocInCollection(test.ToMap(), mongodbapi.GamifiedQuizesCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})
}
func UpdateGamifiedTestPlaysHandler(c *gin.Context) {
	var id = c.Param("id")
	test, err := helperfunctions.GetGamifiedQuizModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Test not found in database"})
		return

	}
	test.TotalPlays += 1

	_, err = mongodbapi.UpdateADocInCollection(test.ToMap(), mongodbapi.GamifiedQuizesCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})
}
func AddGamifiedTestLikeHandler(c *gin.Context) {
	var id = c.Param("id")
	test, err := helperfunctions.GetGamifiedQuizModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Test not found in database"})
		return

	}
	test.Likes += 1

	_, err = mongodbapi.UpdateADocInCollection(test.ToMap(), mongodbapi.GamifiedQuizesCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})
}

func DeleteGamifiedQuizTestHandler(c *gin.Context) {

	var id = c.Param("id")

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.GamifiedQuizesCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	//todo delete gamified quiz images and thumbnails
	c.JSON(200, bson.M{"message": "delete success"})
}

func LikeGamifiedQuizToDbHandler(c *gin.Context) {
	var id = c.Param("id")
	quiz, err := helperfunctions.GetGamifiedQuizModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})

	if err == nil {
		quiz.Likes += 1
		_, err = mongodbapi.UpdateADocInCollection(quiz.ToMap(), mongodbapi.GamifiedQuizesCollection, quiz.TestId)
		if err != nil {
			fmt.Println("unable to update likes in gamefied ")
		}

	}
	c.JSON(200, bson.M{"message": "like success"})

}
func UnLikeGamifiedQuizToDbHandler(c *gin.Context) {
	var id = c.Param("id")
	quiz, err := helperfunctions.GetGamifiedQuizModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})

	if err == nil {
		quiz.Likes -= 1
		_, err = mongodbapi.UpdateADocInCollection(quiz.ToMap(), mongodbapi.GamifiedQuizesCollection, quiz.TestId)
		if err != nil {
			fmt.Println("unable to update likes in gamefied ")
		}

	}
	c.JSON(200, bson.M{"message": "unlike success"})

}
func RateGamifiedQuizToDbHandler(c *gin.Context) {
	var id = c.Param("id")
	rate, _ := strconv.ParseFloat(c.Param("rate"), 64)
	fmt.Println("rate quiz id: ", id, rate)
	c.JSON(200, bson.M{"message": "success"})

}

func AddPlayGamifiedQuizToDbHandler(c *gin.Context) {
	var id = c.Param("id")
	isPass, _ := strconv.ParseBool(c.Param("pass"))

	quiz, err := helperfunctions.GetGamifiedQuizModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})

	if err == nil {
		quiz.TotalPlays += 1
		if isPass {
			quiz.TotalPasses += 1
		}
		_, err = mongodbapi.UpdateADocInCollection(quiz.ToMap(), mongodbapi.GamifiedQuizesCollection, quiz.TestId)
		if err != nil {
			fmt.Println("unable to update likes in gamefied ")
		}

	}
	c.JSON(200, bson.M{"message": "success"})

}

func GetTestThumbnailFromServer(c *gin.Context) {
	thumbnailName := c.Param("id")
	c.File(fmt.Sprintf("assets/quizes/%v", thumbnailName))
}
