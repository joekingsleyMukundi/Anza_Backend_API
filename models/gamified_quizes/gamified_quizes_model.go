package gamifiedquizes

type GamifiedQuizesModel struct {
	TestId            string    `json:"_id" bson:"_id"`
	SubjectType       string    `json:"subjectType" paperUrl:"subjectType" binding:"required"`
	ClassLevel        string    `json:"classLevel" paperUrl:"classLevel" binding:"required"`
	Title             string    `json:"title" bson:"title" binding:"required"`
	PassPercentage    int       `json:"passPercentage" paperUrl:"passPercentage" binding:"required"`
	CreatedAt         string    `json:"createdAt" paperUrl:"createdAt" binding:"required"`
	UpdatedAt         string    `json:"updatedAt" paperUrl:"updatedAt" binding:"required"`
	TotalPlays        int       `json:"totalPlays" paperUrl:"totalPlays"`
	TotalPasses       int       `json:"totalPasses" paperUrl:"totalPasses"`
	Likes             int64     `json:"likes" paperUrl:"likes" binding:"required"`
	Rating            float64   `json:"rating" paperUrl:"rating"`
	OwnerId           string    `json:"ownerId" bson:"ownerId" binding:"required"`
	TeacherName       string    `json:"teacherName" bson:"teacherName" `
	SchoolName        string    `json:"schoolName" bson:"schoolName" binding:"required"`
	ThumbnailUrl      string    `json:"thumbnailUrl" paperUrl:"thumbnailUrl" binding:"required"`
	Questions         []OneQuiz `json:"questions" bson:"questions" binding:"required"`
	Tags              []string  `json:"tags" bson:"tags" binding:"required"`
	IsPublished       bool      `json:"isPublished" bson:"isPublished"`
	TaggedVideoLesson string    `json:"taggedVideoLesson" bson:"taggedVideoLesson"`
}

func (t GamifiedQuizesModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"subjectType":       t.SubjectType,
		"classLevel":        t.ClassLevel,
		"title":             t.Title,
		"passPercentage":    t.PassPercentage,
		"createdAt":         t.CreatedAt,
		"updatedAt":         t.UpdatedAt,
		"totalPlays":        t.TotalPlays,
		"totalPasses":       t.TotalPasses,
		"likes":             t.Likes,
		"rating":            t.Rating,
		"ownerId":           t.OwnerId,
		"teacherName":       t.TeacherName,
		"schoolName":        t.SchoolName,
		"thumbnailUrl":      t.ThumbnailUrl,
		"questions":         t.Questions,
		"tags":              t.Tags,
		"isPublished":       t.IsPublished,
		"taggedVideoLesson": t.TaggedVideoLesson,
	}

}

type OneQuiz struct {
	Category          string   `json:"category" bson:"category" binding:"required"`
	Question          string   `json:"question" bson:"question" binding:"required"`
	ImageUrl          string   `json:"imageUrl" bson:"imageUrl"`
	AnswerOptions     []string `json:"answerOptions" bson:"answerOptions" binding:"required"`
	Answer            string   `json:"answer" bson:"answer" binding:"required"`
	AnswerExplanation string   `json:"answerExplanation" bson:"answerExplanation"`
	Points            int64    `json:"points" bson:"points" binding:"required"`
	TimerSeconds      int      `json:"timerSeconds" bson:"timerSeconds" binding:"required"`
	QuestionNumber    int      `json:"questionNumber" bson:"questionNumber" binding:"required"`
	QuestionType      string   `json:"questionType" bson:"questionType"`
	AnswerType        string   `json:"answerType" bson:"answerType"`
	ExplanationType   string   `json:"explanationType" bson:"explanationType"`
	QuestionImage     string   `json:"questionImage" bson:"questionImage"`
}
