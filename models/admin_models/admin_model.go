package adminmodels

type AdminStatsModel struct {
	//users
	TotalUsers          int `json:"totalUsers"`
	Students            int `json:"students"`
	Teachers            int `json:"teachers"`
	Institution         int `json:"institution"`
	ActiveSubscriptions int `json:"activeSubscriptions"`
	//videos
	TotalVideos int `json:"totalVideos"`
	Form1Videos int `json:"form1Videos"`
	Form2Videos int `json:"form2Videos"`
	Form3Videos int `json:"form3Videos"`
	Form4Videos int `json:"form4Videos"`

	//past papers
	TotalPapers int `json:"totalPapers"`
	Form1Papers int `json:"form1Papers"`
	Form2Papers int `json:"form2Papers"`
	Form3Papers int `json:"form3Papers"`
	Form4Papers int `json:"form4Papers"`

	//Gamified Questions Stats
	TotalQuizzes int `json:"totalQuizzes"`
	Form1Quizes  int `json:"form1Quizes"`
	Form2Quizzes int `json:"form2Quizzes"`
	Form3Quizzes int `json:"form3Quizzes"`
	Form4Quizes  int `json:"form4Quizes"`

	//other content stats
	ExaminerTalks   int `json:"examinerTalks"`
	CareerTalks     int `json:"careerTalks"`
	FasihiEnglish   int `json:"fasihiEnglish"`
	FasihiKiswahili int `json:"fasihiKiswahili"`
}
