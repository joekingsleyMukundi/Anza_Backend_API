package appdatamodel

type AppDataModel struct {
	Id          string            `json:"_id" bson:"_id"`
	Testmonials []TestmonialModel `json:"testmonials" `
	Stats       StatsModel        `json:"stats"`
	Faqs        []FaqModel        `json:"faqs"`
	Schools     []string          `json:"schools"`
}

type TestmonialModel struct {
	Id         string `json:"_id" bson:"_id"`
	ClientName string `json:"clientName" bson:"clientName" binding:"required"`
	Message    string `json:"message" bson:"message" binding:"required"`
	Occupation string `json:"occupation" bson:"occupation" binding:"required"`
	Category   string `json:"category" bson:"category" binding:"required"`
	Date       string `json:"date" bson:"date" binding:"required"`
	ImageUrl   string `json:"imageUrl" bson:"imageUrl" `
	Number     int    `json:"number" bson:"number" `
}

func (t TestmonialModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"clientName": t.ClientName,
		"message":    t.Message,
		"occupation": t.Occupation,
		"category":   t.Category,
		"date":       t.Date,
		"imageUrl":   t.ImageUrl,
		"number":     t.Number,
	}

}

type FaqModel struct {
	Id       string `json:"_id" bson:"_id"`
	Question string `json:"question" bson:"question" binding:"required"`
	Answer   string `json:"answer" bson:"answer" binding:"required"`
	Number   int64  `json:"number" bson:"number" binding:"required"`
}

func (f FaqModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"question": f.Question,
		"answer":   f.Answer,
		"number":   f.Number,
	}

}

type StatsModel struct {
	TotalTeachers int64 `json:"totalTeachers"`
	TotalStudents int64 `json:"totalStudents"`
	TotalLessons  int64 `json:"totalLessons" `
}
