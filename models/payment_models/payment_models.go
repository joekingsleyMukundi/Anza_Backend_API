package paymentmodels

type PaymentRequestPayload struct {
	PhoneNumber uint64 `json:"phoneNumber" binding:"required"`
	Amount      uint64 `json:"amount" binding:"required"`
	Months      string `json:"months" binding:"required"`
	Plan        string `json:"plan" binding:"required"`
	InviteCode  string `json:"inviteCode" binding:"required"`
	PropertyId  string `json:"propertyId" binding:"required"` //knows which user
	UserId      string `json:"userId" binding:"required"`     //knows which user
	Time        string `json:"time" binding:"required"`
}

type WithdrawRequestPayload struct {
	PhoneNumber uint64 `json:"phoneNumber" binding:"required"`
	Amount      uint64 `json:"amount" binding:"required"`
	Time        string `json:"time" binding:"required"`
	FromWallet  bool   `json:"fromWallet"`
	PropertyId  string `json:"propertyId"`
	UserId      string `json:"userId" binding:"required"` //knows which user

}

type SucessfulMpesaPaymentTranscation struct {
	AccountUserName   string `json:"accountUserName" bson:"accountUserName" binding:"required"`
	PhoneNumber       string `json:"phoneNumber" bson:"phoneNumber" binding:"required"`
	Amount            uint64 `json:"amount" bson:"amount" binding:"required"`
	CheckoutRequestID string `json:"checkoutRequestID" bson:"checkoutRequestID" binding:"required"`
	MerchantRequestID string `json:"merchantRequestID" bson:"merchantRequestID" binding:"required"`
	CreatedAt         string `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt         string `json:"updatedAt" bson:"updatedAt" binding:"required"`
}
type SucessfulMpesaWithdrawTranscation struct {
	AccountUserName string `json:"accountUserName" bson:"accountUserName" binding:"required"`
	PhoneNumber     string `json:"phoneNumber" bson:"phoneNumber" binding:"required"`
	Amount          uint64 `json:"amount" bson:"amount" binding:"required"`
	ConversationID  string `json:"conversationID" bson:"conversationID" binding:"required"`
	TransactionID   string `json:"transactionID" bson:"transactionID" binding:"required"`
	CreatedAt       string `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt       string `json:"updatedAt" bson:"updatedAt" binding:"required"`
}

//todo add paypal creditcard payment models

//for tenants
type TenantMpesaPaymentRequestPayload struct {
	PhoneNumber     uint64 `json:"phoneNumber" binding:"required"`
	PayerName       string `json:"payerName" binding:"required"`
	PayerTime       string `json:"payerTime" binding:"required"`
	Amount          uint64 `json:"amount" binding:"required"`
	PropertyId      string `json:"propertyId" binding:"required"`
	HouseNumber     string `json:"houseNumber" binding:"required"`
	BillName        string `json:"billName" binding:"required"`
	RequiredAmount  uint64 `json:"requiredAmount" binding:"required"`
	RemainingAmount uint64 `json:"remainingAmount" binding:"required"`
	IsFullPayment   bool   `json:"isFullPayment" binding:"required"`
	IsCleared       bool   `json:"isCleared" binding:"required"`
	TenantId        string `json:"tenantId" binding:"required"`
	InvitedBy       string `json:"invitedBy" binding:"required"`
}

type TestMpesaRequestPayload struct {
	PhoneNumber    uint64 `json:"phoneNumber" binding:"required"`
	UserId         string `json:"userId" binding:"required"` //knows which user
	Time           string `json:"time" binding:"required"`
	ConsumerKey    string `json:"consumerKey" binding:"required"`
	ConsumerSecret string `json:"consumerSecret" binding:"required"`
	PassKey        string `json:"passKey" binding:"required"`
	ShortCode      uint   `json:"shortCode" binding:"required"`
	IsPaybill      bool   `json:"isPaybill" binding:"required"`
}
