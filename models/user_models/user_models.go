package usermodels

import videolessons "github.com/Anza2022/Anza_Backend_API/models/video_lessons"

type UserModel struct {
	UserId        string `json:"_id" bson:"_id"`
	AccountType   string `json:"accountType" bson:"accountType" binding:"required"`
	SchoolName    string `json:"schoolName" bson:"schoolName"`
	ClassLevel    string `json:"classLevel" bson:"classLevel"`
	UserName      string `json:"userName" bson:"userName" `
	Email         string `json:"email" bson:"email"`
	PhoneNumber   string `json:"phoneNumber" bson:"phoneNumber" binding:"required"`
	PassWord      string `json:"password" bson:"password" `
	ProfilePicUrl string `json:"profilePicUrl" bson:"profilePicUrl"`

	IsAdmin               bool   `json:"isAdmin" bson:"isAdmin"`
	AdminType             string `json:"adminType" bson:"adminType"`
	IsSuperAdmin          bool   `json:"isSuperAdmin" bson:"isSuperAdmin"`
	IsPhoneNumberVerified bool   `json:"isPhoneNumberVerified" bson:"isPhoneNumberVerified"`
	IsEmailVerified       bool   `json:"isEmailVerified" bson:"isEmailVerified"`
	IsFirstTimeLogin      bool   `json:"isFirstTimeLogin" bson:"isFirstTimeLogin"`
	IsGoogleSignUp        bool   `json:"isGoogleSignUp" bson:"isGoogleSignUp"`
	IsLoggedIn            bool   `json:"isLoggedIn" bson:"isLoggedIn"`
	CreatedAt             string `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt             string `json:"updatedAt" bson:"updatedAt" binding:"required"`
	TscId                 string `json:"tscId" bson:"tscId"`
	PrimarySubject        string `json:"primarySubject" bson:"primarySubject"`
	SecondarySubject      string `json:"secondarySubject" bson:"secondarySubject"`

	IdNumber             string `json:"idNumber" bson:"idNumber"`
	IsAcceptedByAdmin    bool   `json:"isAcceptedByAdmin" bson:"isAcceptedByAdmin"`
	LastMpesaPhoneNumber string `json:"lastMpesaPhoneNumber" bson:"lastMpesaPhoneNumber"`
	DateOfBirth          string `json:"dateOfBirth" bson:"dateOfBirth"`
}

func (u UserModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"accountType":           u.AccountType,
		"schoolName":            u.SchoolName,
		"userName":              u.UserName,
		"classLevel":            u.ClassLevel,
		"email":                 u.Email,
		"phoneNumber":           u.PhoneNumber,
		"password":              u.PassWord,
		"profilePicUrl":         u.ProfilePicUrl,
		"isAdmin":               u.IsAdmin,
		"adminType":             u.AdminType,
		"isSuperAdmin":          u.IsSuperAdmin,
		"isEmailVerified":       u.IsEmailVerified,
		"isPhoneNumberVerified": u.IsPhoneNumberVerified,
		"isGoogleSignUp":        u.IsGoogleSignUp,
		"isFirstTimeLogin":      u.IsFirstTimeLogin,
		"isLoggedIn":            u.IsLoggedIn,
		"createdAt":             u.CreatedAt,
		"updatedAt":             u.UpdatedAt,
		"tscId":                 u.TscId,
		"primarySubject":        u.PrimarySubject,
		"secondarySubject":      u.SecondarySubject,

		"lastMpesaPhoneNumber": u.LastMpesaPhoneNumber,
		"isAcceptedByAdmin":    u.IsAcceptedByAdmin,
		"idNumber":             u.IdNumber,
		"dateOfBirth":          u.DateOfBirth,
	}
}

type AccountSubscription struct {
	SubscriptionId          string `json:"_id" bson:"_id"`
	UserId                  string `json:"userId" bson:"userId"`
	SubscriptionStartDate   string `json:"subscriptionStartDate"  bson:"subscriptionStartDate"`
	SubscriptionEndDate     string `json:"subscriptionEndDate"  bson:"subscriptionEndDate"`
	CurrentSubscriptionPlan string `json:"currentSubscriptionPlan"  bson:"currentSubscriptionPlan"`
	Resubscriptions         int32  `json:"resubscriptions"  bson:"resubscriptions"`
	AccountLockDate         string `json:"accountLockDate"  bson:"accountLockDate"`
	IsRefererPaid           bool   `json:"isRefererPaid" bson:"isRefererPaid"`
}

func (s AccountSubscription) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"userId":                  s.UserId,
		"subscriptionStartDate":   s.SubscriptionStartDate,
		"subscriptionEndDate":     s.SubscriptionEndDate,
		"currentSubscriptionPlan": s.CurrentSubscriptionPlan,
		"resubscriptions":         s.Resubscriptions,
		"accountLockDate":         s.AccountLockDate,
		"isRefererPaid":           s.IsRefererPaid,
	}

}

type UserStats struct {
	StatsId              string                        `json:"_id" bson:"_id"`
	UserId               string                        `json:"userId" bson:"userId" binding:"required"`
	LastMpesaPhonenumber string                        `json:"lastMpesaPhonenumber"  bson:"lastMpesaPhonenumber"`
	FavoriteLessonIds    []string                      `json:"favoriteLessonIds"  bson:"favoriteLessonIds"`
	TotalVideosWatched   int64                         `json:"totalVideosWatched"  bson:"totalVideosWatched"`
	TotalTestTaken       int64                         `json:"totalTestTaken"  bson:"totalTestTaken"`
	TotalMarksScored     int64                         `json:"totalMarksScored"  bson:"totalMarksScored"`
	LastWatchedVideo     videolessons.VideoLessonModel `json:"lastWatchedVideo"  bson:"lastWatchedVideo"`
	LastVideoTimeSecs    int64                         `json:"lastVideoTimeSecs"  bson:"lastVideoTimeSecs"`
}

func (s UserStats) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"userId":               s.UserId,
		"lastMpesaPhonenumber": s.LastMpesaPhonenumber,
		"favoriteLessonIds":    s.FavoriteLessonIds,
		"totalVideosWatched":   s.TotalVideosWatched,
		"totalTestTaken":       s.TotalTestTaken,
		"totalMarksScored":     s.TotalMarksScored,
		"lastVideoTimeSecs":    s.LastVideoTimeSecs,
		"lastWatchedVideo":     s.LastWatchedVideo.ToMap(),
	}

}

type LoginUserPayload struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	PassWord string `json:"password" bson:"password" binding:"required"`
}

//add isRead and isResponed fields for admin panel
type ClientFeedbackModel struct {
	ClientName  string `json:"clientName" bson:"clientName" binding:"required"`
	ClientEmail string `json:"clientEmail" bson:"clientEmail" binding:"required"`
	Reason      string `json:"reason" bson:"reason" binding:"required"`
	Message     string `json:"message" bson:"message" binding:"required"`
	CreatedAt   string `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt   string `json:"updatedAt" bson:"updatedAt" binding:"required"`
}

type AdminStatsModel struct {
	TotalUsers          int64 `json:"totalUsers"`
	TotalOwners         int64 `json:"totalOwners"`
	TotalTenants        int64 `json:"totalTenants"`
	TotalProperties     int64 `json:"totalProperties"`
	TotalListings       int64 `json:"totalListings"`
	TotalTransactions   int64 `json:"totalTransactions"`
	TotalMaintainances  int64 `json:"totalMaintainances"`
	OnlineLoggedInUsers int64 `json:"onlineLoggedInUsers"`
}
