package schoolmodels

type SchoolModel struct {
	Id                  string `json:"_id" bson:"_id"`
	SchoolName          string `json:"schoolName" bson:"schoolName" binding:"required"`
	SchoolEmail         string `json:"schoolEmail" bson:"schoolEmail" `
	PhoneNumber         string `json:"phoneNumber" bson:"phoneNumber" `
	IsAuthorisedByAdmin bool   `json:"isAuthorisedByAdmin" bson:"isAuthorisedByAdmin" `
	CreatedAt           string `json:"createdAt" bson:"createdAt" `
}

func (s SchoolModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"schoolName":          s.SchoolName,
		"schoolEmail":         s.SchoolEmail,
		"phoneNumber":         s.PhoneNumber,
		"isAuthorisedByAdmin": s.IsAuthorisedByAdmin,
		"createdAt":           s.CreatedAt,
	}

}
