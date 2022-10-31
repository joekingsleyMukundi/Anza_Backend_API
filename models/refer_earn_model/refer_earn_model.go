package referearnmodel

type ReferalBonus struct {
	Id               string   `json:"_id" bson:"_id"`
	UserId           string   `json:"userId" bson:"userId" binding:"required"`
	InviteCode       string   `json:"inviteCode" bson:"inviteCode"`
	BonusCode        string   `json:"bonusCode" bson:"bonusCode"  binding:"required" `
	AmountEarned     int64    `json:"amountEarned" bson:"amountEarned"   `
	WidthdrawnAmount int64    `json:"withdrawnAmount" bson:"withdrawnAmount"   `
	ReferredUsersIds []string `json:"referredUsersIds" bson:"referredUsersIds"`
}

func (r ReferalBonus) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"userId":           r.UserId,
		"inviteCode":       r.InviteCode,
		"bonusCode":        r.BonusCode,
		"amountEarned":     r.AmountEarned,
		"withdrawnAmount":  r.WidthdrawnAmount,
		"referredUsersIds": r.ReferredUsersIds,
	}
}
