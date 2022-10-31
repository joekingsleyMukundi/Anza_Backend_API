package videolessons

type VideoLessonModel struct {
	VideoId       string         `json:"_id" bson:"_id"`
	Title         string         `json:"title" bson:"title" binding:"required"`
	VideoUrl      string         `json:"videoUrl" bson:"videoUrl" binding:"required"`
	ThumbnailUrl  string         `json:"thumbnailUrl" bson:"thumbnailUrl" binding:"required"`
	TopicName     string         `json:"topicName" bson:"topicName" binding:"required"`
	ClassLevel    string         `json:"classLevel" bson:"classLevel" binding:"required"`
	SubjectType   string         `json:"subjectType" bson:"subjectType" binding:"required"`
	TeacherName   string         `json:"teacherName" bson:"teacherName" binding:"required"`
	TotalViews    int64          `json:"totalViews" bson:"totalViews" `
	TopicPriority int64          `json:"topicPriority" bson:"topicPriority"`
	VideoPriority int64          `json:"videoPriority" bson:"videoPriority"`
	CreatedAt     string         `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt     string         `json:"updatedAt" bson:"updatedAt" binding:"required"`
	CommentsOff   bool           `json:"commentsOff" bson:"commentsOff"`
	IsPublished   bool           `json:"isPublished" bson:"isPublished"`
	TeacherSchool string         `json:"teacherSchool" bson:"teacherSchool"`
	NotesId       string         `json:"notesId" bson:"notesId"`
	LessonQuizes  []InLessonQuiz `json:"lessonQuizes" bson:"lessonQuizes"`
	LessonBrief   string         `json:"lessonBrief" bson:"lessonBrief"`
	Chapter       string         `json:"chapter" bson:"chapter"`
	ChapterNumber int64          `json:"chapterNumber" bson:"chapterNumber"`
	Likes         int64          `json:"likes" bson:"likes"`
	DisLikes      int64          `json:"disLikes" bson:"disLikes"`
	VideoLength   string         `json:"videoLength" bson:"videoLength"`
	IsPractical   bool           `json:"isPractical" bson:"isPractical"`
}

func (l VideoLessonModel) ToMap() map[string]interface{} {

	var lessonquizes = []map[string]interface{}{}

	if len(l.LessonQuizes) > 0 {
		for i := 0; i < len(l.LessonQuizes); i++ {
			lessonquizes = append(lessonquizes, l.LessonQuizes[i].ToMap())
		}

	}
	return map[string]interface{}{
		"title":         l.Title,
		"videoUrl":      l.VideoUrl,
		"thumbnailUrl":  l.ThumbnailUrl,
		"topicName":     l.TopicName,
		"classLevel":    l.ClassLevel,
		"subjectType":   l.SubjectType,
		"teacherName":   l.TeacherName,
		"totalViews":    l.TotalViews,
		"topicPriority": l.TopicPriority,
		"videoPriority": l.VideoPriority,
		"createdAt":     l.CreatedAt,
		"updatedAt":     l.UpdatedAt,
		"commentsOff":   l.CommentsOff,
		"isPublished":   l.IsPublished,
		"teacherSchool": l.TeacherSchool,
		"notesId":       l.NotesId,
		"lessonQuizes":  lessonquizes,
		"lessonBrief":   l.LessonBrief,
		"chapter":       l.Chapter,
		"chapterNumber": l.ChapterNumber,
		"likes":         l.Likes,
		"disLikes":      l.DisLikes,
		"videoLength":   l.VideoLength,
		"isPractical":   l.IsPractical,
	}
}

type InLessonQuiz struct {
	Question      string   `json:"question" bson:"question"`
	CorrectAnswer string   `json:"correctAnswer" bson:"correctAnswer"`
	AnswerOptions []string `json:"answerOptions" bson:"answerOptions"`
	Time          int64    `json:"time" bson:"time"`
}

func (l InLessonQuiz) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"question":      l.Question,
		"correctAnswer": l.CorrectAnswer,
		"answerOptions": l.AnswerOptions,
		"time":          l.Time,
	}

}

type LessonNotes struct {
	NotesId  string `json:"_id" bson:"_id"`
	Notes    string `json:"notes" bson:"notes" binding:"required"`
	LessonId string `json:"lessonId" bson:"lessonId" binding:"required"`
}

func (l LessonNotes) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"notes":    l.Notes,
		"lessonId": l.LessonId,
	}

}

type SetBookEpisode struct {
	Id string `json:"_id" bson:"_id"`

	Title         string `json:"title" bson:"title" binding:"required"`
	Book          string `json:"book" bson:"book" binding:"required"`
	FileName      string `json:"fileName" bson:"fileName" binding:"required"`
	EpisodeNumber int64  `json:"episodeNumber" bson:"episodeNumber" binding:"required"`
	CreatedAt     string `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt     string `json:"updatedAt" bson:"updatedAt" binding:"required"`
	Description   string `json:"description" bson:"description"`
	Likes         int64  `json:"likes" bson:"likes"`
	Views         int64  `json:"views" bson:"views"`
}

func (l *SetBookEpisode) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":         l.Title,
		"book":          l.Book,
		"episodeNumber": l.EpisodeNumber,
		"fileName":      l.FileName,
		"createdAt":     l.CreatedAt,
		"updatedAt":     l.UpdatedAt,
		"views":         l.Views,
		"likes":         l.Likes,
		"description":   l.Description,
	}

}

type MkurugenziEpisode struct {
	Id            string `json:"_id" bson:"_id"`
	Title         string `json:"title" bson:"title" binding:"required"`
	FileName      string `json:"fileName" bson:"fileName" binding:"required"`
	ThumbnailUrl  string `json:"thumbnailUrl" bson:"thumbnailUrl" binding:"required"`
	EpisodeNumber int64  `json:"episodeNumber" bson:"episodeNumber" binding:"required"`
	CreatedAt     string `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt     string `json:"updatedAt" bson:"updatedAt" binding:"required"`
	Description   string `json:"description" bson:"description"`
	Likes         int64  `json:"likes" bson:"likes"`
	Views         int64  `json:"views" bson:"views"`
}

func (l *MkurugenziEpisode) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":         l.Title,
		"episodeNumber": l.EpisodeNumber,
		"fileName":      l.FileName,
		"thumbnailUrl":  l.ThumbnailUrl,
		"createdAt":     l.CreatedAt,
		"updatedAt":     l.UpdatedAt,
		"views":         l.Views,
		"likes":         l.Likes,
		"description":   l.Description,
	}

}
