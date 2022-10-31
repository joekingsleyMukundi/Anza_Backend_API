package talksmodels

type ExaminerTalkModel struct {
	TalkId              string `json:"_id" bson:"_id"`
	Title               string `json:"title" bson:"title" binding:"required"`
	VideoUrl            string `json:"videoUrl" bson:"videoUrl" binding:"required"`
	ThumbnailUrl        string `json:"thumbnailUrl" bson:"thumbnailUrl" binding:"required"`
	TotalViews          int64  `json:"totalViews" bson:"totalViews" `
	Likes               int64  `json:"likes" bson:"likes" `
	CreatedAt           string `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt           string `json:"updatedAt" bson:"updatedAt" binding:"required"`
	SubjectType         string `json:"subjectType" bson:"subjectType" binding:"required"`
	ExaminerSchool      string `json:"examinerSchool" bson:"examinerSchool"`
	ExaminerName        string `json:"examinerName" bson:"examinerName"`
	ExaminerDescription string `json:"examinerDescription" bson:"examinerDescription"`
}

func (t ExaminerTalkModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":               t.Title,
		"videoUrl":            t.VideoUrl,
		"thumbnailUrl":        t.ThumbnailUrl,
		"totalViews":          t.TotalViews,
		"likes":               t.Likes,
		"createdAt":           t.CreatedAt,
		"updatedAt":           t.UpdatedAt,
		"subjectType":         t.SubjectType,
		"examinerSchool":      t.ExaminerSchool,
		"examinerName":        t.ExaminerName,
		"examinerDescription": t.ExaminerDescription,
	}

}

type CareerTalkModel struct {
	TalkId       string `json:"_id" bson:"_id"`
	Title        string `json:"title" bson:"title" binding:"required"`
	VideoUrl     string `json:"videoUrl" bson:"videoUrl" binding:"required"`
	ThumbnailUrl string `json:"thumbnailUrl" bson:"thumbnailUrl" binding:"required"`
	TotalViews   int64  `json:"totalViews" bson:"totalViews" `
	Likes        int64  `json:"likes" bson:"likes" `
	CreatedAt    string `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt    string `json:"updatedAt" bson:"updatedAt" binding:"required"`
	Category     string `json:"category" bson:"category" binding:"required"`
	GuestName    string `json:"guestName" bson:"guestName" binding:"required"`
	Description  string `json:"description" bson:"description"`
}

func (t CareerTalkModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":        t.Title,
		"videoUrl":     t.VideoUrl,
		"thumbnailUrl": t.ThumbnailUrl,
		"totalViews":   t.TotalViews,
		"likes":        t.Likes,
		"createdAt":    t.CreatedAt,
		"updatedAt":    t.UpdatedAt,
		"category":     t.Category,
		"guestName":    t.GuestName,
		"description":  t.Description,
	}

}
