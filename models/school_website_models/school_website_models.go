package schoolwebsitemodels

type SchoolWebsiteModel struct {
	Id                    string            `json:"_id" bson:"_id"`
	SchoolName            string            `json:"schoolName" bson:"schoolName" binding:"required"`
	SchoolTagline         string            `json:"schoolTagline" bson:"schoolTagline"`
	AccountId             string            `json:"accountId" bson:"accountId" binding:"required"` //id of the  school account operating the website
	LogoFile              string            `json:"logoFile" bson:"logoFile"`
	Mission               string            `json:"mission" bson:"mission"`
	Vision                string            `json:"vision" bson:"vision"`
	MainColor             string            `json:"mainColor" bson:"mainColor"`
	SecondaryColor        string            `json:"secondaryColor" bson:"secondaryColor"`
	Values                []string          `json:"values" bson:"values"` //array of ValueModel thar has value and description
	Testmonials           []TestmonialModel `json:"testmonials" bson:"testmonials"`
	UpcomingEvents        []SchoolEvent     `json:"upcomingEvents" bson:"upcomingEvents"` //array of eventmodel --date,title, event, descriotion
	OurTeam               []TeamMember      `json:"ourTeam" bson:"ourTeam"`               //array of teamMember model
	Assignments           []AssignmentModel `json:"assignments" bson:"assignments"`       //array of Assignments model
	ContactDetails        ContactDetails    `json:"contactDetails" bson:"contactDetails"` //ContactDetailsModels -social links, email,phonumber location
	LandingCarouselImages []string          `json:"landingCarouselImages" bson:"landingCarouselImages"`
	PartnersLogos         []string          `json:"partnersLogos" bson:"partnersLogos"`
	Faqs                  []FaqModel        `json:"faqs" bson:"faqs"`
	TermDates             TermDatesModel    `json:"termDates" bson:"termDates"`

	CreatedAt string `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt string `json:"updatedAt" bson:"updatedAt" binding:"required"`
}

func (s *SchoolWebsiteModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"schoolName":            s.SchoolName,
		"schoolTagline":         s.SchoolTagline,
		"accountId":             s.AccountId,
		"logoFile":              s.LogoFile,
		"mission":               s.Mission,
		"vision":                s.Vision,
		"mainColor":             s.MainColor,
		"values":                s.Values,
		"secondaryColor":        s.SecondaryColor,
		"testmonials":           s.Testmonials,
		"upcomingEvents":        s.UpcomingEvents,
		"ourTeam":               s.OurTeam,
		"assignments":           s.Assignments,
		"contactDetails":        s.ContactDetails,
		"landingCarouselImages": s.LandingCarouselImages,
		"partnersLogos":         s.PartnersLogos,
		"termDates":             s.TermDates,
		"createdAt":             s.CreatedAt,
		"updatedAt":             s.UpdatedAt,
	}

}

type TestmonialModel struct {
	FullName   string `json:"fullName" bson:"fullName" binding:"required"`
	Message    string `json:"message" bson:"message" binding:"required"`
	Occupation string `json:"occupation" bson:"occupation" binding:"required"`
	ImageUrl   string `json:"imageUrl" bson:"imageUrl" `
	Number     int    `json:"number" bson:"number" `
}

type FaqModel struct {
	Question string `json:"question" bson:"question" binding:"required"`
	Answer   string `json:"answer" bson:"answer" binding:"required"`
	Number   int64  `json:"number" bson:"number" binding:"required"`
}

type ContactDetails struct {
	PrimaryPhoneNumber   string `json:"primaryPhoneNumber" bson:"primaryPhoneNumber"`
	SecondaryPhoneNumber string `json:"secondaryPhoneNumber" bson:"secondaryPhoneNumber"`
	Email                string `json:"email" bson:"email"`
	Location             string `json:"location" bson:"location"`
	FacebookLink         string `json:"facebookLink" bson:"facebookLink"`
	TwitterLink          string `json:"twitterLink" bson:"twitterLink"`
	WhatsappNumber       string `json:"whatsappNumber" bson:"whatsappNumber"`
}

type AssignmentModel struct {
	ClassLevel  string `json:"classLevel" bson:"classLevel"`
	SubjectType string `json:"subjectType" bson:"subjectType"`
	FileName    string `json:"fileName" bson:"fileName"`
}
type TeamMember struct {
	Category   string `json:"category" bson:"category"`
	Occupation string `json:"occupation" bson:"occupation"`
	FullName   string `json:"fullName" bson:"fullName"`
	ImageUrl   string `json:"imageUrl" bson:"imageUrl"`
}

type SchoolEvent struct {
	Title           string `json:"title" bson:"title"`
	Date            string `json:"date" bson:"date"`
	Time            string `json:"time" bson:"time"`
	Description     string `json:"description" bson:"description"`
	EventAgendaFile string `json:"eventAgendaFile" bson:"eventAgendaFile"`
}

type ImageModel struct {
	Category string `json:"category" bson:"category"` //partners, carosel, gallery, aluminis, sports
	FileName string `json:"fileName" bson:"fileName"`
}

type TermDatesModel struct {
	Term              string `json:"term" bson:"term"`
	OpeningDate       string `json:"openingDate" bson:"openingDate"`
	MidtermBreakDate  string `json:"midtermBreakDate" bson:"midtermBreakDate"`
	MidtermReopenDate string `json:"midtermReopenDate" bson:"midtermReopenDate"`
	ClosingDate       string `json:"closingDate" bson:"closingDate"`
}
