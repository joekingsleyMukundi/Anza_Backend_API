package liveclassesmodels

type LiveClassModel struct {
	Id            string   `json:"_id" bson:"_id"`
	Title         string   `json:"title" bson:"title" binding:"required"`
	DayDate       string   `json:"dayDate" bson:"dayDate" binding:"required"`
	DayTime       string   `json:"dayTime" bson:"dayTime" binding:"required"`
	Duration      int64    `json:"duration" bson:"duration"`
	BbbUrl        string   `json:"bbbUrl" bson:"bbbUrl"`
	Price         int64    `json:"price" bson:"price" binding:"required"`
	StudentIds    []string `json:"studentsIds" bson:"studentsIds"` //students who have booked the class
	CreatedAt     string   `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt     string   `json:"updatedAt" bson:"updatedAt" binding:"required"`
	IsVerified    bool     `json:"isVerified" bson:"isVerified"`
	OwnerId       string   `json:"ownerId" bson:"ownerId" binding:"required"`
	TeacherName   string   `json:"teacherName" bson:"teacherName" binding:"required"`
	TeacherSchool string   `json:"teacherSchool" bson:"teacherSchool" binding:"required"`
}

func (c *LiveClassModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":         c.Title,
		"price":         c.Price,
		"dayDate":       c.DayDate,
		"dayTime":       c.DayTime,
		"duration":      c.Duration,
		"bbbUrl":        c.BbbUrl,
		"studentsIds":   c.StudentIds,
		"createdAt":     c.CreatedAt,
		"updatedAt":     c.UpdatedAt,
		"isVerified":    c.IsVerified,
		"ownerId":       c.OwnerId,
		"teacherName":   c.TeacherName,
		"teacherSchool": c.TeacherSchool,
	}

}
