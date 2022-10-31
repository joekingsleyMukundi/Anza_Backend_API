package authmodels

type LoginPayLoadModel struct {
	//email andpassword only.
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber" binding:"required"`
	Password    string `json:"password" bson:"password"  binding:"required"`
}

type PasswordResetPayloadModel struct {
	//only email is required
	Email string `json:"email" bson:"email" binding:"required"`
}
type PasswordResetWithCodePayloadModel struct {
	//only email is required
	Email       string `json:"email" bson:"email" binding:"required"`
	ResetCode   string `json:"resetCode" bson:"resetCode" binding:"required"`
	NewPassword string `json:"newPassword" bson:"newPassword" binding:"required"`
}

type UserModel struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" `
}

type PassWordResetDbModel struct {
	RequestId string `json:"_id" bson:"_id" `
	Email     string `json:"email" bson:"email" binding:"required"`
	ResetCode string `json:"resetCode" bson:"resetCode" binding:"required"`
	UserId    string `json:"userId" bson:"userId" binding:"required"`
	ExpiresAt int64  `json:"expiresAt" bson:"expiresAt" binding:"required"`
}

func (rq PassWordResetDbModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"email":     rq.Email,
		"userId":    rq.UserId,
		"resetCode": rq.ResetCode,
		"expiresAt": rq.ExpiresAt,
	}

}

//used for storing the  token in redis
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}
