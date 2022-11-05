package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kennedy-muthaura/anzaapi/middlewares"
	adminroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/admin_route_handlers"
	appdataroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/appdata_route_handlers"
	authroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/auth_route_handlers"
	careertalksroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/career_talks_route_handlers"
	commentsroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/comments_route_handlers"
	examinertalksroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/examiner_talks_route_handlers"
	gamifiedquizesroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/gamified_quizes_route_handlers"
	liveclassesroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/live_classes_route_handlers"
	mkurugenziroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/mkurugenzi_route_handlers"
	notificationroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/notification_route_handlers"
	pastpapersroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/past_papers_route_handlers"
	schoolsroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/schools_route_handlers"
	setbooksroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/setbooks_route_handlers"
	teacherroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/teacher_route_handlers"
	userroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/user_route_handlers"
	videolessonsroutehandlers "github.com/kennedy-muthaura/anzaapi/route_handlers/video_lessons_route_handlers"
	appcronjobs "github.com/kennedy-muthaura/anzaapi/utils/app_cronjobs"
	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const productionMode = false

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://ANZA-ACADEMY-DB-MAIN-USER:vFfYuBZstuK9d7uX@anza-academy-db.wqrdqvl.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal("unable to connect to database")
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(aurora.BgRed("Db unable to connect"))
		fmt.Println(err)
	}
	fmt.Println(aurora.Blink(aurora.Green("Db Connected Successfully")))

	appcronjobs.StartServerCronJobs()
	fmt.Println(aurora.Blink(aurora.Green("automatic cron jobs running successfully")))

	router := gin.Default()

	// if productionMode {

	// 	// gin.SetMode(gin.ReleaseMode)

	// }
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		// AllowAllOrigins:  true,
		AllowOrigins: []string{"*"},

		MaxAge:          12 * time.Hour,
		AllowWebSockets: true,
	}))

	//User Router papers
	router.GET("/anzaapi/users", middlewares.CacheRequestMiddleware(), authroutehandlers.LoginAUserHandler)
	router.GET("/anzaapi/refresh_user/:id", authroutehandlers.RefreshUserHandler)
	router.PATCH("/anzaapi/user/:id", userroutehandlers.UpdateUserDetailsHandler)
	router.PATCH("/anzaapi/user_stats/:id", userroutehandlers.UpdateUserStatsHandler)
	router.DELETE("/anzaapi/user/:id", authroutehandlers.LoginAUserHandler)

	//video lessons
	router.GET("/anzaapi/lessons", middlewares.CacheRequestMiddleware(), videolessonsroutehandlers.GetAllVideoLessonsDocsHandler)
	router.GET("/anzaapi/video_lesson/add_like/:id", videolessonsroutehandlers.AddaVideoLikeInLesson)
	router.GET("/anzaapi/video_lesson/add_dislike/:id", videolessonsroutehandlers.AddaVideoDisLikeInLesson)
	router.GET("/anzaapi/video_lesson/add_view/:id", videolessonsroutehandlers.AddaVideoViewInLesson)
	router.POST("/anzaapi/lessons", videolessonsroutehandlers.PostVideoLessonsToDbHandler)
	router.PATCH("/anzaapi/lesson/:id", videolessonsroutehandlers.UpdateVideoLessonHandler)
	router.DELETE("/anzaapi/lesson/:id", videolessonsroutehandlers.DeleteVideoLessonHandler)                     //todo delete thumbnail, video file and notes
	router.GET("/anzaapi/view_video/lesson/:id", videolessonsroutehandlers.GetAVideoLessonFromServer)            //serves the video file
	router.GET("/anzaapi/view_thumbnail/lesson/:id", videolessonsroutehandlers.GetVideoClassThumbnailFromServer) //serves the video thumbnail
	router.POST("/anzaapi/upload_video/lesson", videolessonsroutehandlers.UploadVideoToServer)                   //upload the video file
	router.POST("/anzaapi/upload_thumbnail/lesson", videolessonsroutehandlers.UploadVideoThumbnailToServer)      //upload the video thumbnail
	router.GET("/anzaapi/uploads/exists/:filename", videolessonsroutehandlers.CheckifFileExistInServer)          //upload the video thumbnail
	router.GET("/anzaapi/lesson_files/all", videolessonsroutehandlers.GetAllLessonVideoFilesInTheServer)         //upload the video thumbnail

	//todo lesson notes
	router.GET("/anzaapi/lesson_note/:lessonid", middlewares.CacheRequestMiddleware(), videolessonsroutehandlers.GetLessonsNotesHandler)
	router.POST("/anzaapi/lesson_notes", videolessonsroutehandlers.AddLessonNotes)
	router.PATCH("/anzaapi/lesson_note/:id", videolessonsroutehandlers.UpdateLessonNotesHandler)
	router.DELETE("/anzaapi/lesson_note/:id", videolessonsroutehandlers.DeleteLessonNotesHandler)

	//todo comments routes
	router.GET("/anzaapi/comments/:id", middlewares.CacheRequestMiddleware(), commentsroutehandlers.GetVideoLessonComments)
	router.GET("/anzaapi/comment_replies/:id", commentsroutehandlers.GetVideoLessonCommentReplies)
	router.GET("/anzaapi/like_comment/:id", commentsroutehandlers.GetVideoLessonCommentReplies)
	router.GET("/anzaapi/dislike_comment/:id", commentsroutehandlers.GetVideoLessonCommentReplies)
	router.POST("/anzaapi/comments", commentsroutehandlers.PostVideoCommentToDbHandler)
	router.PATCH("/anzaapi/comment/:id", commentsroutehandlers.UpdateVideoCommentHandler)
	router.DELETE("/anzaapi/comment/:id", commentsroutehandlers.DeleteVideoCommentHandler)
	router.GET("/anzaapi/all_comments", commentsroutehandlers.GetAllVideoLessonComments)

	//todo schemes of work routes
	router.GET("/anzaapi/teacher/workschemes", middlewares.CacheRequestMiddleware(), teacherroutehandlers.GetAllSchemesOfWorkRouteHandler)
	router.POST("/anzaapi/teacher/workschemes", teacherroutehandlers.AddSchemeOfWorkDbHandler)
	router.PATCH("/anzaapi/teacher/workscheme/:id", teacherroutehandlers.UpdateSchemOfWorkHandler)
	router.DELETE("/anzaapi/teacher/workscheme/:id", teacherroutehandlers.DeleteSchemeOfWorkHandler)
	router.GET("/anzaapi/teacher/view_scheme/:id", teacherroutehandlers.GetWorkSchemeFromServer)

	//todo lesson plans  routes
	router.GET("/anzaapi/teacher/lesson_plans", middlewares.CacheRequestMiddleware(), teacherroutehandlers.GetAllLessonsPlansRouteHandler)
	router.POST("/anzaapi/teacher/lesson_plans", teacherroutehandlers.AddLessonPlanToDbHandler)
	router.PATCH("/anzaapi/teacher/lesson_plan/:id", teacherroutehandlers.UpdateLessonPlanHandler)
	router.DELETE("/anzaapi/teacher/lesson_plan/:id", teacherroutehandlers.DeleteLessonPlanHandler)
	router.GET("/anzaapi/teacher/view_plan/:id", teacherroutehandlers.GetLessonPlanFromServer)

	//todo setbooks  routes
	router.GET("/anzaapi/setbook_episodes", middlewares.CacheRequestMiddleware(), setbooksroutehandlers.GetAllLSetBookEpisodesRouteHandler)
	router.POST("/anzaapi/setbook_episodes", setbooksroutehandlers.AddSetBookEpisodeToDbHandler)
	router.PATCH("/anzaapi/setbook_episode/:id", setbooksroutehandlers.UpdateSetBookEpisodeHandler)
	router.DELETE("/anzaapi/setbook_episode/:id", setbooksroutehandlers.DeleteSetBookEpisodeHandler)
	router.GET("/anzaapi/setbook/view_episode/:id", setbooksroutehandlers.GetSetBookEpisodeFromServer)
	router.GET("/anzaapi/setbook_files/all", setbooksroutehandlers.GetAllSetbookFilesInTheServer)
	//todo mkurugenzi   routes
	router.GET("/anzaapi/mkurugenzi_episodes", middlewares.CacheRequestMiddleware(), mkurugenziroutehandlers.GetAllLMkurugenziEpisodesRouteHandler)
	router.POST("/anzaapi/mkurugenzi_episodes", mkurugenziroutehandlers.AddMkurugenziEpisodeToDbHandler)
	router.PATCH("/anzaapi/mkurugenzi_episode/:id", mkurugenziroutehandlers.UpdateMkurugenziEpisodeHandler)
	router.DELETE("/anzaapi/mkurugenzi_episode/:id", mkurugenziroutehandlers.DeleteMkurugenziEpisodeHandler)
	router.GET("/anzaapi/mkurugenzi/view_episode/:id", mkurugenziroutehandlers.GetMkurugenziEpisodeFromServer)
	router.GET("/anzaapi/mkurugenzi_files/all", mkurugenziroutehandlers.GetAllMkurugenziFilesInTheServer)
	router.GET("/anzaapi/view_thumbnail/mkurugenzi/:id", mkurugenziroutehandlers.GetMkurugenziThumbnailFromServer) //serves the video thumbnail

	//todo live classes   routes
	router.GET("/anzaapi/live_classes", middlewares.CacheRequestMiddleware(), liveclassesroutehandlers.GetAllLiveClassesRouteHandler)
	router.GET("/anzaapi/teacher/live_classes/:id", liveclassesroutehandlers.GetTeacherLiveClassesRouteHandler)
	router.POST("/anzaapi/live_classes", liveclassesroutehandlers.AddLiveClassToDbHandler)
	router.PATCH("/anzaapi/live_class/:id", liveclassesroutehandlers.UpdateLiveClassHandler)
	router.DELETE("/anzaapi/live_class/:id", liveclassesroutehandlers.DeleteLiveClassHandler)

	//todo past papers routes
	router.GET("/anzaapi/past_papers", middlewares.CacheRequestMiddleware(), pastpapersroutehandlers.GetAllPastPapersRouteHandlers)
	router.GET("/anzaapi/teacher/past_papers/:id", middlewares.CacheRequestMiddleware(), pastpapersroutehandlers.GetAllTeachersPastPapersRouteHandlers)
	router.POST("/anzaapi/past_papers", pastpapersroutehandlers.AddPastPaperToDbHandler)
	router.PATCH("/anzaapi/past_paper/:id", pastpapersroutehandlers.UpdatePastPaperHandler)
	router.DELETE("/anzaapi/past_paper/:id", pastpapersroutehandlers.DeletePastPaperHandler)
	router.GET("/anzaapi/view_past_paper/:paperid", pastpapersroutehandlers.GetPastPaperFromServer)
	//notes route
	router.GET("/anzaapi/notes/pdf_notes/:paperid", pastpapersroutehandlers.GetPdfNotesFromServer) //serves lesson pdf notes
	//quizes papers
	router.GET("/anzaapi/gamified_quizes", middlewares.CacheRequestMiddleware(), gamifiedquizesroutehandlers.GetAllGamifiedQuestionsRouteHandler)
	router.GET("/anzaapi/teacher/gamified_quizes/:id", middlewares.CacheRequestMiddleware(), gamifiedquizesroutehandlers.GetTeachersGamifiedQuestionsRouteHandler)
	router.POST("/anzaapi/gamified_quizes", gamifiedquizesroutehandlers.AddGamifiedQuizesToDbHandler)
	router.PATCH("/anzaapi/gamified_quiz/:id", gamifiedquizesroutehandlers.UpdateGamifiedQuizTestHandler)
	router.PATCH("/anzaapi/update_test/add_play/:id", gamifiedquizesroutehandlers.UpdateGamifiedTestPlaysHandler)
	router.PATCH("/anzaapi/update_test/add_like/:id", gamifiedquizesroutehandlers.LikeGamifiedQuizToDbHandler)
	router.PATCH("/anzaapi/update_test/publish/:id", gamifiedquizesroutehandlers.PublishGamifiedTestForStudentsHandler)
	router.PATCH("/anzaapi/update_test/un_publish/:id", gamifiedquizesroutehandlers.UnPublishGamifiedTestForStudentsHandler)
	router.DELETE("/anzaapi/gamified_quiz/:id", gamifiedquizesroutehandlers.DeleteGamifiedQuizTestHandler)
	router.GET("/anzaapi/rate/gamified_quizes/:id/:rate", gamifiedquizesroutehandlers.RateGamifiedQuizToDbHandler)
	router.GET("/anzaapi/like/gamified_quizes/:id", gamifiedquizesroutehandlers.LikeGamifiedQuizToDbHandler)
	router.GET("/anzaapi/unlike/gamified_quizes/:id", gamifiedquizesroutehandlers.UnLikeGamifiedQuizToDbHandler)
	router.GET("/anzaapi/quiz_test/add_play/:id/:pass", gamifiedquizesroutehandlers.AddPlayGamifiedQuizToDbHandler)
	router.GET("/anzaapi/view_thumbnail/quizes/:id", gamifiedquizesroutehandlers.GetTestThumbnailFromServer) //serves the video thumbnail

	//Examiner Talks routes
	router.GET("/anzaapi/examiner_talks", middlewares.CacheRequestMiddleware(), examinertalksroutehandlers.GetAllExaminerTalksRouteHandler)
	router.POST("/anzaapi/examiner_talks", examinertalksroutehandlers.AddExaminierTalkToDbHandler)
	router.PATCH("/anzaapi/examiner_talk/:id", examinertalksroutehandlers.UpdateExaminerTalkHandler)
	router.DELETE("/anzaapi/examiner_talk/:id", examinertalksroutehandlers.DeleteExaminerTalkHandler)
	//Career Talks routes
	router.GET("/anzaapi/career_talks", middlewares.CacheRequestMiddleware(), careertalksroutehandlers.GetAllCareerTalksRouteHandler)
	router.POST("/anzaapi/career_talks", careertalksroutehandlers.AddCareerTalkToDbHandler)
	router.PATCH("/anzaapi/career_talk/:id", careertalksroutehandlers.UpdateCareerTalkHandler)
	router.DELETE("/anzaapi/career_talk/:id", careertalksroutehandlers.DeleteCareerTalkHandler)
	router.GET("/anzaapi/view_video/career/:id", careertalksroutehandlers.GetACareerFromServer)
	router.GET("/anzaapi/view_thumbnail/career/:id", careertalksroutehandlers.GetCareerThumbnailFromServer)
	//files routes, notification routes

	//schools routes
	router.GET("/anzaapi/all_schools", middlewares.CacheRequestMiddleware(), schoolsroutehandlers.GetAllSchoolsRouteHandler)
	router.POST("/anzaapi/schools", schoolsroutehandlers.AddSchoolsRouteHandler)
	router.DELETE("/anzaapi/school/:id", schoolsroutehandlers.DeleteSchoolRouteHandler)
	//todo authentication routes
	router.POST("/anzaapi/auth/login", authroutehandlers.LoginAUserHandler)
	router.POST("/anzaapi/auth/google/login", authroutehandlers.GoogleLoginAUserHandler)
	router.POST("/anzaapi/auth/register", authroutehandlers.RegisterUserHandler)
	router.POST("/anzaapi/auth/change_password/:userId", authroutehandlers.ChangeUserPassword)
	router.GET("/anzaapi/auth/refresh_token", middlewares.RefreshTokenAuthMiddleware(), authroutehandlers.TokenRefreshHandler)
	router.POST("/anzaapi/auth/email_password_reset/code", authroutehandlers.SendPasswordResetCodeToMailHander)
	router.POST("/anzaapi/auth/reset_password/code", authroutehandlers.ResetPasswordWithCodeHandler)
	// router.GET("/anzaapi/auth/web_reset/:userId", authroutehandlers.GetTemplateForPassWordReset)
	// router.POST("/anzaapi/auth/email_password_reset/web", authroutehandlers.SendResetPasswordMailWeb)
	// router.POST("/anzaapi/auth/password_reset_weblink/:linkid", authroutehandlers.ReceiveNewPasswordFromWeb)
	// router.GET("anzaapi/auth/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"name":      "Kennedy muthaura",
	// 		"year":      "2021",
	// 		"dontmatch": "",
	// 	})
	// })
	//todo payments routes
	router.GET("/anzaapi/payments/mpesa/")

	//todo notification routes
	router.GET("/anzaapi/user_notifications/:userId", middlewares.CacheRequestMiddleware(), notificationroutehandlers.GetAllUnreadUserNotifications)
	router.POST("/anzaapi/notifications", middlewares.CacheRequestMiddleware(), notificationroutehandlers.AddNewUserNotification)
	router.PATCH("/anzaapi/notification/:id", notificationroutehandlers.UpdateAUserNotification)
	router.PATCH("/anzaapi/update_notifications/read", notificationroutehandlers.MarkNotificationsAdReadFromIdsList)

	//admin routes
	router.GET("/anzaapi/admin/stats", middlewares.CacheRequestMiddleware(), adminroutehandlers.GetAdminStats)
	router.GET("/anzaapi/app/data", middlewares.CacheRequestMiddleware(), appdataroutehandlers.GetAppData)
	//testmonials and faqs
	router.GET("/anzaapi/app/testmonials", middlewares.CacheRequestMiddleware(), appdataroutehandlers.GetAllTestimonialsRouteHandler)
	router.POST("/anzaapi/app/testmonials", appdataroutehandlers.AddTestmonialToDbHandler)
	router.GET("/anzaapi/app/faqs", middlewares.CacheRequestMiddleware(), appdataroutehandlers.GetAllFaqsRouteHandler)
	router.POST("/anzaapi/app/faqs", appdataroutehandlers.AddFaqToDbHandler)

	//todo tcorner routes ,, exam generator, mypastpapers,tutions , assignments,study groups,tutions and my quizes and timetable

	//todo scorner routes -- assignments , assignment helper

	//

	router.GET("/anzaapi/test", func(c *gin.Context) {
		c.JSON(200, map[string]string{"message": "Anza Api is Up And Running up and running ,, kudos to ken"})
	})

	port := ":8080"
	fmt.Println(aurora.BgBrightGreen(fmt.Sprintf("Server running on port %v", port)))

	if false {
		endless.ListenAndServe(port, router)
	} else {
		router.Run(port)
	}
}
