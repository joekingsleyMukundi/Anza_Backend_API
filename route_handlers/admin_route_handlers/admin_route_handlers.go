package adminroutehandlers

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	adminmodels "github.com/kennedy-muthaura/anzaapi/models/admin_models"
	"github.com/kennedy-muthaura/anzaapi/services/mongodbapi"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAdminStats(c *gin.Context) {
	newStats := adminmodels.AdminStatsModel{
		TotalUsers: 0, Students: 0, Teachers: 0, Institution: 0, ActiveSubscriptions: 0,
		TotalVideos: 0, Form1Videos: 0, Form2Videos: 0, Form3Videos: 0, Form4Videos: 0,
		TotalPapers: 0, Form1Papers: 0, Form2Papers: 0, Form3Papers: 0, Form4Papers: 0,
		TotalQuizzes: 0, Form1Quizes: 0, Form2Quizzes: 0, Form3Quizzes: 0, Form4Quizes: 0,
		ExaminerTalks: 0, CareerTalks: 0, FasihiEnglish: 0, FasihiKiswahili: 0,
	}
	ctx := context.Background()
	//get users
	cursor, err := mongodbapi.UsersCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user bson.M
		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		newStats.TotalUsers += 1
		if user["accountType"] == "teacher" {
			newStats.Teachers += 1
		}
		if user["accountType"] == "student" {
			newStats.Students += 1
		}
		if user["accountType"] == "school" {
			newStats.Institution += 1
		}
	}
	//get videos
	cursor, err = mongodbapi.VideoLessonsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var lesson bson.M
		if err = cursor.Decode(&lesson); err != nil {
			log.Fatal(err)
		}
		newStats.TotalVideos += 1
		if lesson["classLevel"] == "form 1" {
			newStats.Form1Videos += 1
		}
		if lesson["classLevel"] == "form 2" {
			newStats.Form2Videos += 1
		}
		if lesson["classLevel"] == "form 3" {
			newStats.Form3Videos += 1
		}
		if lesson["classLevel"] == "form 4" {
			newStats.Form4Videos += 1
		}
	}
	//get papers
	cursor, err = mongodbapi.PastPapersCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var paper bson.M
		if err = cursor.Decode(&paper); err != nil {
			log.Fatal(err)
		}
		newStats.TotalPapers += 1
		if paper["classLevel"] == "form 1" {
			newStats.Form1Papers += 1
		}
		if paper["classLevel"] == "form 2" {
			newStats.Form2Papers += 1
		}
		if paper["classLevel"] == "form 3" {
			newStats.Form3Papers += 1
		}
		if paper["classLevel"] == "form 4" {
			newStats.Form4Papers += 1
		}
	}
	//get quizzes
	cursor, err = mongodbapi.GamifiedQuizesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var quiz bson.M
		if err = cursor.Decode(&quiz); err != nil {
			log.Fatal(err)
		}
		newStats.TotalQuizzes += 1
		if quiz["classLevel"] == "form 1" {
			newStats.Form1Quizes += 1
		}
		if quiz["classLevel"] == "form 2" {
			newStats.Form2Quizzes += 1
		}
		if quiz["classLevel"] == "form 3" {
			newStats.Form3Quizzes += 1
		}
		if quiz["classLevel"] == "form 4" {
			newStats.Form4Quizes += 1
		}
	}
	//get other content stats
	docsnumber, err := mongodbapi.ExaminerTalksCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal("unable to count documents")
	}
	newStats.ExaminerTalks = int(docsnumber)
	careerTalksNumber, err := mongodbapi.CareerTalksCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal("unable to count documents")
	}
	newStats.CareerTalks = int(careerTalksNumber)

	c.JSON(200, bson.M{"stats": newStats})

}
