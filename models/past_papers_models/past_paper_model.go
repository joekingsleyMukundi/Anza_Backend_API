package pastpapersmodels

type PastPaperModel struct {
	PaperId          string `json:"_id" bson:"_id"`
	Title            string `json:"title" bson:"title" binding:"required"`
	PaperUrl         string `json:"paperUrl" paperUrl:"paperUrl" binding:"required"`
	MarkingSchemeUrl string `json:"markingSchemeUrl" paperUrl:"markingSchemeUrl"`
	ClassLevel       string `json:"classLevel" paperUrl:"classLevel" binding:"required"`
	SubjectType      string `json:"subjectType" paperUrl:"subjectType" binding:"required"`
	PaperNumber      string `json:"paperNumber" paperUrl:"paperNumber"`
	SchoolName       string `json:"schoolName" paperUrl:"schoolName" binding:"required"`
	OwnerId          string `json:"ownerId" paperUrl:"ownerId"`
	TeacherName      string `json:"teacherName" paperUrl:"teacherName" binding:"required"`
	CreatedAt        string `json:"createdAt" paperUrl:"createdAt" binding:"required"`
	UpdatedAt        string `json:"updatedAt" paperUrl:"updatedAt" binding:"required"`
	Year             int64  `json:"year" paperUrl:"year" `
	Term             string `json:"term" paperUrl:"term" `
}

func (p PastPaperModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":            p.Title,
		"paperUrl":         p.PaperUrl,
		"markingSchemeUrl": p.MarkingSchemeUrl,
		"classLevel":       p.ClassLevel,
		"subjectType":      p.SubjectType,
		"paperNumber":      p.PaperNumber,
		"schoolName":       p.SchoolName,
		"teacherName":      p.TeacherName,
		"ownerId":          p.OwnerId,
		"createdAt":        p.CreatedAt,
		"updatedAt":        p.UpdatedAt,
		"year":             p.Year,
		"term":             p.Term,
	}
}
