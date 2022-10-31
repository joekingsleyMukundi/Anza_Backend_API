package commentsroutehandlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	commentmodel "github.com/kennedy-muthaura/anzaapi/models/comment_model"
	"github.com/kennedy-muthaura/anzaapi/services/mongodbapi"
	"github.com/kennedy-muthaura/anzaapi/utils/appconstants"
	helperfunctions "github.com/kennedy-muthaura/anzaapi/utils/helper_functions"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllVideoLessonComments(c *gin.Context) {
	comments, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.VideoCommentsCollection, bson.M{})

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to get lessons"})

	} else {
		c.JSON(200, comments)
	}
}
func GetVideoLessonComments(c *gin.Context) {
	var id = c.Param("id")
	comments, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.VideoCommentsCollection, bson.M{"videoId": id})

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to get lessons"})

	} else {
		c.JSON(200, comments)
	}
}
func GetVideoLessonCommentReplies(c *gin.Context) {
	var id = c.Param("id")
	comments, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.VideoCommentsCollection, bson.M{"videoId": id, "isReply": true})

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to get lessons"})

	} else {
		c.JSON(200, comments)
	}
}

func PostVideoCommentToDbHandler(c *gin.Context) {
	var comment commentmodel.CommentModel
	if err := c.ShouldBind(&comment); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(comment.ToMap(), mongodbapi.VideoCommentsCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdateVideoCommentHandler(c *gin.Context) {
	var comment commentmodel.CommentModel
	var id = c.Param("id")

	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(comment.ToMap(), mongodbapi.VideoCommentsCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

func DeleteVideoCommentHandler(c *gin.Context) {

	var id = c.Param("id")

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.VideoCommentsCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "delete success"})
}

func LikeVideoLessonComment(c *gin.Context) {
	var id = c.Param("id")

	comment, err := helperfunctions.GetCommentModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "lesson not found"})
		return
	}
	comment.Likes += 1

	_, err = mongodbapi.UpdateADocInCollection(comment.ToMap(), mongodbapi.VideoLessonsCollection, comment.CommentId)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to update, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "liked"})
}
func DisLikeVideoLessonComment(c *gin.Context) {
	var id = c.Param("id")

	comment, err := helperfunctions.GetCommentModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "lesson not found"})
		return
	}
	comment.Likes -= 1

	_, err = mongodbapi.UpdateADocInCollection(comment.ToMap(), mongodbapi.VideoLessonsCollection, comment.CommentId)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to update, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "disliked"})
}
