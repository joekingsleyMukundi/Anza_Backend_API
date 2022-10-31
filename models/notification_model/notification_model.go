package notificationmodel

type Notificationmodel struct {
	Id       string `json:"_id" bson:"_id"`
	Category string `json:"category" bson:"category" binding:"required"` //payment , maintainance request, invited tenant,
	Message  string `json:"message" bson:"message" binding:"required"`
	UserId   string `json:"userId" bson:"userId"  `

	CreatedDate  string `json:"createdDate" bson:"createdDate" binding:"required"`
	CreatedTime  string `json:"createdTime" bson:"createdTime" binding:"required"`
	IsReadByUser bool   `json:"isReadByUser" bson:"isReadByUser"`
}

func (n Notificationmodel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"category":     n.Category,
		"message":      n.Message,
		"userId":       n.UserId,
		"createdDate":  n.CreatedDate,
		"createdTime":  n.CreatedTime,
		"isReadByUser": n.IsReadByUser,
	}
}
