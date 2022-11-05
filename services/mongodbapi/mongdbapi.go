package mongodbapi

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Todo always remember to change the database  and collections of the project
var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

// defer cancel()
var client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://ANZA-ACADEMY-DB-MAIN-USER:vFfYuBZstuK9d7uX@anza-academy-db.wqrdqvl.mongodb.net/?retryWrites=true&w=majority"))

var anzaacademydb = client.Database("anza-academy-db")

//football collections
var UsersCollection = anzaacademydb.Collection("usersCollection")
var AccountSubscriptionCollection = anzaacademydb.Collection("accountSubscriptionCollection")
var UserStatsCollection = anzaacademydb.Collection("userStatsCollection")
var VideoLessonsCollection = anzaacademydb.Collection("videoLessonsCollection")
var LessonNotesCollection = anzaacademydb.Collection("LessonNotesCollection")
var VideoCommentsCollection = anzaacademydb.Collection("videoCommentsCollection")
var PastPapersCollection = anzaacademydb.Collection("pastPapersCollection")
var GamifiedQuizesCollection = anzaacademydb.Collection("gamifiedQuizesCollection")
var ExaminerTalksCollection = anzaacademydb.Collection("examinerTalksCollection")
var CareerTalksCollection = anzaacademydb.Collection("careerTalksCollection")
var NotificationsCollection = anzaacademydb.Collection("notificationsCollection")
var PassWordResetRequestsCollection = anzaacademydb.Collection("passWordResetRequestsCollection")
var ReferBonusesDataCollection = anzaacademydb.Collection("referBonusesDataCollection")
var SchoolsCollection = anzaacademydb.Collection("schoolsCollection")
var TestmonialsCollection = anzaacademydb.Collection("testmonialsCollection")
var FaqsCollection = anzaacademydb.Collection("faqsCollection")
var LessonPlansCollection = anzaacademydb.Collection("lessonPlansCollection")
var SchemesOfWorkCollection = anzaacademydb.Collection("schemesOfWorkCollection")
var SetBooksEpisodesCollection = anzaacademydb.Collection("setBooksEpisodesCollection")
var MkurugenziEpisodesCollection = anzaacademydb.Collection("mkurugenziEpisodesCollection")
var LiveClassesCollection = anzaacademydb.Collection("liveClassesCollection")

func AddADocumentToCollection(doc interface{}, collection *mongo.Collection) (string, error) {
	newdoc, err := collection.InsertOne(context.Background(), doc)

	if err != nil {
		fmt.Println("unable to add document in database")
		return "", errors.New("unable to save the document")
	}

	return newdoc.InsertedID.(primitive.ObjectID).Hex(), nil
}

func DeleteDocInACollection(collection *mongo.Collection, docid string) bool {
	id, err := primitive.ObjectIDFromHex(docid)

	if err != nil {
		fmt.Println("cannot convert the parsed id to mongodb id --deleting")
		return false
	}

	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})

	if err != nil {

		fmt.Println("unable to delete the document in db")
		fmt.Println(err)
		return false
	}

	if res.DeletedCount > 0 {

		return true
	} else {
		return false
	}

}

func UpdateADocInCollection(doc interface{}, collection *mongo.Collection, docid string) (string, error) {
	id, err := primitive.ObjectIDFromHex(docid)

	if err != nil {
		fmt.Println("cannot convert the parsed id to mongodb id")
		return "", errors.New("id provided is invalid")
	}

	res, err := collection.ReplaceOne(context.Background(), bson.M{"_id": id}, doc)

	if err != nil {
		fmt.Println(err)
		return "", errors.New("unable to update the document")
	}

	if res.UpsertedID == nil {
		return docid, nil
	}

	return fmt.Sprintf("%v", res.UpsertedID), nil

}

func RetrieveOneDocumentInACollectionWithId(collection *mongo.Collection, docid string) (bson.M, error) {
	id, err := primitive.ObjectIDFromHex(docid)

	if err != nil {
		fmt.Println("cannot convert the parsed id to mongodb id -retrieving doc")

	}
	var resdoc bson.M

	doc := collection.FindOne(context.Background(), bson.M{"_id": id})

	if doc.Err() != nil {

		fmt.Println(doc.Err().Error())
		return nil, errors.New("unable to get the document")
	}

	doc.Decode(&resdoc)

	return resdoc, nil

}

func GetManyDocumentsFromACollection(collection *mongo.Collection, filter bson.M) ([]bson.M, error) {

	var findOptions options.FindOptions = options.FindOptions{Sort: bson.M{"createdAt": -1}}

	cursor, err := collection.Find(context.Background(), filter, &findOptions)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("unable to get the documents")
	}
	var docs []bson.M
	if err = cursor.All(context.Background(), &docs); err != nil {
		fmt.Println("unable to retrieve many documents")

		log.Fatal(err)

		return nil, errors.New("unable to decode docs to slice")
	}

	return docs, nil

}

func RetrieveOneDocumentInACollection(collection *mongo.Collection, filter map[string]interface{}) (bson.M, error) {
	var resdoc bson.M

	doc := collection.FindOne(context.Background(), filter)

	if doc.Err() != nil {

		fmt.Println(doc.Err().Error())
		return nil, errors.New("unable to get the document")
	}

	doc.Decode(&resdoc)

	return resdoc, nil

}
