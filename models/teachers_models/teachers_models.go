package teachersmodels

type SchemesOfWorkModel struct {
	Id          string `json:"_id" bson:"_id"`
	Title       string `json:"title" bson:"title" binding:"required"`
	FileId      string `json:"fileId" bson:"fileId" binding:"required"`
	SubjectType string `json:"subjectType" bson:"subjectType" binding:"required"`
	ClassLevel  string `json:"classLevel" bson:"classLevel" binding:"required"`
	CreatedAt   string `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt   string `json:"updatedAt" bson:"updatedAt" binding:"required"`
	Term        string `json:"term" bson:"term" binding:"required"`
}

func (s *SchemesOfWorkModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":       s.Title,
		"fileId":      s.FileId,
		"subjectType": s.SubjectType,
		"classLevel":  s.ClassLevel,
		"createdAt":   s.CreatedAt,
		"updatedAt":   s.UpdatedAt,
		"term":        s.Term,
	}
}

type LessonPlansModel struct {
	Id            string `json:"_id" bson:"_id"`
	Title         string `json:"title" bson:"title" binding:"required"`
	FileId        string `json:"fileId" bson:"fileId" binding:"required"`
	ClassLevel    string `json:"classLevel" bson:"classLevel" binding:"required"`
	SubjectType   string `json:"subjectType" bson:"subjectType" binding:"required"`
	Chapter       string `json:"chapter" bson:"chapter" binding:"required"`
	ChapterNumber int64  `json:"chapterNumber" bson:"chapterNumber" binding:"required"`
	Topic         string `json:"topic" bson:"topic" binding:"required"`
	TopicNumber   int64  `json:"topicNumber" bson:"topicNumber" binding:"required"`
	CreatedAt     string `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt     string `json:"updatedAt" bson:"updatedAt" binding:"required"`
	Term          string `json:"term" bson:"term" binding:"required"`
}

func (s *LessonPlansModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":         s.Title,
		"fileId":        s.FileId,
		"subjectType":   s.SubjectType,
		"classLevel":    s.ClassLevel,
		"chapter":       s.Chapter,
		"chapterNumber": s.ChapterNumber,
		"topic":         s.Topic,
		"topicNumber":   s.TopicNumber,
		"createdAt":     s.CreatedAt,
		"updatedAt":     s.UpdatedAt,
		"term":          s.Term,
	}
}
