package appdataroutehandlers

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	appdatamodel "github.com/kennedy-muthaura/anzaapi/models/app_data_model"
	"github.com/kennedy-muthaura/anzaapi/services/mongodbapi"
	"github.com/kennedy-muthaura/anzaapi/utils/appconstants"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAppData(c *gin.Context) {
	testmonials, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.TestmonialsCollection, bson.M{})
	if err != nil {
		fmt.Println("unable to count documents")
	}
	faqs, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.FaqsCollection, bson.M{})
	if err != nil {
		fmt.Println("unable to count documents")
	}
	schools, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.SchoolsCollection, bson.M{})
	if err != nil {
		fmt.Println("unable to count documents")
	}

	newStats := appdatamodel.StatsModel{TotalTeachers: 0, TotalStudents: 0, TotalLessons: 0}
	ctx := context.Background()
	students, err := mongodbapi.UsersCollection.CountDocuments(ctx, bson.M{"accountType": "student"})
	if err != nil {
		log.Fatal("unable to count documents")
	}
	teachers, err := mongodbapi.UsersCollection.CountDocuments(ctx, bson.M{"accountType": "teacher"})
	if err != nil {
		log.Fatal("unable to count documents")
	}
	lessons, err := mongodbapi.VideoLessonsCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal("unable to count documents")
	}
	newStats.TotalStudents = students
	newStats.TotalTeachers = teachers
	newStats.TotalLessons = lessons
	c.JSON(200, bson.M{
		"testmonials": testmonials,
		"stats":       newStats,
		"faqs":        faqs,
		"schools":     schools,
	})
}

func GetAllTestimonialsRouteHandler(c *gin.Context) {

	testmonials, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.TestmonialsCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No Examiner Talks Available"})

	} else {
		c.JSON(200, testmonials)

	}

}

func AddTestmonialToDbHandler(c *gin.Context) {
	var testmonial appdatamodel.TestmonialModel
	if err := c.ShouldBind(&testmonial); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(testmonial.ToMap(), mongodbapi.TestmonialsCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}
func GetAllFaqsRouteHandler(c *gin.Context) {

	faqs, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.FaqsCollection, bson.M{})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "No faqs Available"})

	} else {
		c.JSON(200, faqs)

	}

}

func AddFaqToDbHandler(c *gin.Context) {
	var faq appdatamodel.FaqModel
	if err := c.ShouldBind(&faq); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(faq.ToMap(), mongodbapi.FaqsCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}
