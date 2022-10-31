package notificationroutehandlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	notificationmodel "github.com/kennedy-muthaura/anzaapi/models/notification_model"
	"github.com/kennedy-muthaura/anzaapi/services/mongodbapi"
	"github.com/kennedy-muthaura/anzaapi/utils/appconstants"
	helperfunctions "github.com/kennedy-muthaura/anzaapi/utils/helper_functions"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUnreadUserNotifications(c *gin.Context) {


	userId := c.Param("userId")

	notifications, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.NotificationsCollection, bson.M{"userId": userId})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "failed:try again later"})
		return

	}

	c.JSON(200, notifications)
}

func AddNewUserNotification(c *gin.Context) {

	var newNotification notificationmodel.Notificationmodel
	if err := c.ShouldBind(&newNotification); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(newNotification.ToMap(), mongodbapi.NotificationsCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}

	c.JSON(201, bson.M{"_id": newid})

}
func AddNewNotificationWithUserEmail(c *gin.Context) {
	useremail := c.Query("email")
	var newNotification notificationmodel.Notificationmodel
	if err := c.ShouldBind(&newNotification); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	usermodel, err := helperfunctions.GetUserModelWithFilterFromDb(bson.M{"email": useremail})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "notification cannot be added to unexisting email"})
		return

	}
	newNotification.UserId = usermodel.UserId

	newid, err := mongodbapi.AddADocumentToCollection(newNotification.ToMap(), mongodbapi.NotificationsCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}

	c.JSON(201, bson.M{"_id": newid})

}

func UpdateAUserNotification(c *gin.Context) {
	var id = c.Param("id")
	var updatedNotification notificationmodel.Notificationmodel
	if err := c.ShouldBind(&updatedNotification); err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(updatedNotification.ToMap(), mongodbapi.NotificationsCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to update, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

type readNotificationsRequest struct {
	Ids []string `json:"ids" binding:"required"`
}

func MarkNotificationsAdReadFromIdsList(c *gin.Context) {

	var payload readNotificationsRequest
	if err := c.ShouldBind(&payload); err != nil {

		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided--maintainance property ids"})
		return

	}

	for _, id := range payload.Ids {
		notification, err := helperfunctions.GetNotificationModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})
		if err != nil {
			continue
		}
		notification.IsReadByUser = true
		mongodbapi.UpdateADocInCollection(notification.ToMap(), mongodbapi.NotificationsCollection, notification.Id)

	}

	c.JSON(200, bson.M{"message": "notifications updated as read successfully"})
}
