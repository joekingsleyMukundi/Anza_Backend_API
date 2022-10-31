package videolessonsroutehandlers

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	videolessons "github.com/kennedy-muthaura/anzaapi/models/video_lessons"
	"github.com/kennedy-muthaura/anzaapi/services/mongodbapi"
	"github.com/kennedy-muthaura/anzaapi/utils/appconstants"
	helperfunctions "github.com/kennedy-muthaura/anzaapi/utils/helper_functions"
	"go.mongodb.org/mongo-driver/bson"
)

var productionUrl = "https://api.thesigurd.com/anzaapi"

func GetAllVideoLessonsDocsHandler(c *gin.Context) {
	lessons, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.VideoLessonsCollection, bson.M{"classLevel": "form 3"})

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to get lessons"})

	} else {
		c.JSON(200, lessons)
	}
}

func PostVideoLessonsToDbHandler(c *gin.Context) {
	var lesson videolessons.VideoLessonModel
	if err := c.ShouldBind(&lesson); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(lesson.ToMap(), mongodbapi.VideoLessonsCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdateVideoLessonHandler(c *gin.Context) {
	var lesson videolessons.VideoLessonModel
	var id = c.Param("id")

	if err := c.ShouldBind(&lesson); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(lesson.ToMap(), mongodbapi.VideoLessonsCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

func DeleteVideoLessonHandler(c *gin.Context) {
	var id = c.Param("id")
	//todo remove thumbnail, and video ,, get lesson data from db , get the link and uuid from link, detele it in the assets folder
	//todo fix this onece demo lesson url and thumbnails are removed
	// lesson, err := helperfunctions.GetVideoLessonModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})

	// if err == nil {
	// 	splitfilename := strings.Split(lesson.VideoUrl, "/")
	// 	filename := splitfilename[len(splitfilename)-1]
	// 	helperfunctions.DeleteFileInServer(filename)

	// 		splitthumbanailname := strings.Split(lesson.ThumbnailUrl, "/")
	// 		msname := splitthumbanailname[len(splitthumbanailname)-1]
	// 		helperfunctions.DeleteFileInServer(msname)

	// }

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.VideoLessonsCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "delete success"})
}
func AddaVideoLikeInLesson(c *gin.Context) {
	var id = c.Param("id")

	lesson, err := helperfunctions.GetVideoLessonModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "lesson not found"})
		return
	}
	lesson.Likes += 1

	_, err = mongodbapi.UpdateADocInCollection(lesson.ToMap(), mongodbapi.VideoLessonsCollection, lesson.VideoId)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to update, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "liked"})
}
func AddaVideoDisLikeInLesson(c *gin.Context) {
	var id = c.Param("id")

	lesson, err := helperfunctions.GetVideoLessonModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "lesson not found"})
		return
	}
	lesson.DisLikes += 1

	_, err = mongodbapi.UpdateADocInCollection(lesson.ToMap(), mongodbapi.VideoLessonsCollection, lesson.VideoId)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to update, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "disliked"})
}
func AddaVideoViewInLesson(c *gin.Context) {
	var id = c.Param("id")

	lesson, err := helperfunctions.GetVideoLessonModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(id)})

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "lesson not found"})
		return
	}
	lesson.TotalViews += 1

	_, err = mongodbapi.UpdateADocInCollection(lesson.ToMap(), mongodbapi.VideoLessonsCollection, lesson.VideoId)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to update, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "sucess"})
}

//todo lesson notes route handlers
func GetLessonsNotesHandler(c *gin.Context) {
	var id = c.Param("lessonid")
	notes, err := mongodbapi.GetManyDocumentsFromACollection(mongodbapi.LessonNotesCollection, bson.M{"lessonId": id})

	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to get lessons"})

	} else {
		c.JSON(200, notes)
	}
}
func AddLessonNotes(c *gin.Context) {
	var notes videolessons.LessonNotes
	if err := c.ShouldBind(&notes); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}

	newid, err := mongodbapi.AddADocumentToCollection(notes.ToMap(), mongodbapi.LessonNotesCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(201, bson.M{"_id": newid})

}

func UpdateLessonNotesHandler(c *gin.Context) {
	var notes videolessons.LessonNotes
	var id = c.Param("id")

	if err := c.ShouldBind(&notes); err != nil {

		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "all fieds were not provided"})
		return

	}
	_, err := mongodbapi.UpdateADocInCollection(notes.ToMap(), mongodbapi.LessonNotesCollection, id)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to save, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "update success"})

}

func DeleteLessonNotesHandler(c *gin.Context) {
	//todo remove thumbnail, and video ,, get lesson data from db , get the link and uuid from link, detele it in the assets folder

	var id = c.Param("id")

	isDeleted := mongodbapi.DeleteDocInACollection(mongodbapi.LessonNotesCollection, id)
	if !isDeleted {
		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": "Failed to delete, try again later"})
		return
	}
	c.JSON(200, bson.M{"message": "delete success"})
}

//files handlers

func GetAVideoLessonFromServer(c *gin.Context) {
	videoName := c.Param("id")

	c.Status(http.StatusAccepted)

	c.File(fmt.Sprintf("assets/videoclasses/videos/%v", videoName))
}

func GetVideoClassThumbnailFromServer(c *gin.Context) {
	fmt.Println("getting thumbanail from server")
	thumbnailName := c.Param("id")

	c.File(fmt.Sprintf("assets/videoclasses/thumbnails/%v", thumbnailName))

}

type videoLesson struct {
	Lesson    *multipart.FileHeader `form:"lesson" binding:"required"`
	Type      string                `form:"type" binding:"required"` //lesson, examiner, career
	Extension string                `form:"extension"`               //lesson, examiner, career
}
type videoThumbnail struct {
	Photo *multipart.FileHeader `form:"photo" binding:"required"`
	Type  string                `form:"type" binding:"required"` //lesson, examiner, career
}

//size 540 *  360 max 150 kbs
func UploadVideoToServer(c *gin.Context) {

	var form videoLesson
	err := c.ShouldBind(&form)

	if err != nil {

		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "cannot bind the thumbnail form"})
		return
	}
	videoId := uuid.New()

	filename := fmt.Sprintf("%v.%v", videoId, "mp4")
	videolink := ""
	savePath := "assets/"
	if form.Type == "lesson" {
		savePath = fmt.Sprintf("%v%v", savePath, "videoclasses/videos")
		videolink = fmt.Sprintf("%v/view_video/lesson/%v", productionUrl, filename)
	}
	if form.Type == "career" {
		savePath = fmt.Sprintf("%v%v", savePath, "careertalks/videos")
		videolink = fmt.Sprintf("%v/view_video/career/%v", productionUrl, filename)
	}
	if form.Type == "examiner" {
		savePath = fmt.Sprintf("%v%v", savePath, "examinertalks/videos")
		videolink = fmt.Sprintf("%v/view_video/examiner/%v", productionUrl, filename)
	}
	if form.Type == "paper" {

		filename = fmt.Sprintf("%v.%v", videoId, "pdf")
		savePath = fmt.Sprintf("%v%v", savePath, "pastpapers/")
		videolink = fmt.Sprintf("%v/view_past_paper/%v", productionUrl, filename)
	}
	if form.Type == "ms" {

		filename = fmt.Sprintf("%v.%v", videoId, "pdf")
		savePath = fmt.Sprintf("%v%v", savePath, "pastpapers/")
		videolink = fmt.Sprintf("%v/view_past_paper/%v", productionUrl, filename)

	}
	if form.Type == "notes" {

		filename = fmt.Sprintf("%v.%v", videoId, "pdf")
		savePath = fmt.Sprintf("%v%v", savePath, "pdfnotes/")
		err = c.SaveUploadedFile(form.Lesson, fmt.Sprintf("%v/%v", savePath, filename))
		if err == nil {
			c.JSON(201, map[string]string{"id": filename})

		} else {
			c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to save pdf file"})

		}

		return

	}
	if form.Type == "work_sheme" {

		filename = fmt.Sprintf("%v.%v", videoId, form.Extension)
		savePath = fmt.Sprintf("%v%v", savePath, "workschemes/")
		err = c.SaveUploadedFile(form.Lesson, fmt.Sprintf("%v/%v", savePath, filename))
		if err == nil {
			c.JSON(201, map[string]string{"id": filename})

		} else {
			c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to save work scheme"})

		}

		return

	}
	if form.Type == "lesson_plan" {

		filename = fmt.Sprintf("%v.%v", videoId, form.Extension)
		savePath = fmt.Sprintf("%v%v", savePath, "lessonplans/")
		err = c.SaveUploadedFile(form.Lesson, fmt.Sprintf("%v/%v", savePath, filename))
		if err == nil {
			c.JSON(201, map[string]string{"id": filename})

		} else {
			c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to save lesson plan"})

		}

		return

	}
	if form.Type == "fun_quizes" {

		filename = fmt.Sprintf("%v.%v", videoId, form.Extension)
		savePath = fmt.Sprintf("%v%v", savePath, "quizes/")
		err = c.SaveUploadedFile(form.Lesson, fmt.Sprintf("%v/%v", savePath, filename))
		if err == nil {
			c.JSON(201, map[string]string{"id": filename})
		} else {
			c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to save quiz thumbnail"})

		}

		return

	}

	err = c.SaveUploadedFile(form.Lesson, fmt.Sprintf("%v/%v", savePath, filename))
	if err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to save file"})
		return
	}

	c.JSON(201, map[string]string{"url": videolink})
}

func UploadVideoThumbnailToServer(c *gin.Context) {

	var form videoThumbnail
	err := c.ShouldBind(&form)

	if err != nil {

		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "cannot bind the thumbnail form"})
		return
	}

	videoId := uuid.New()

	filename := fmt.Sprintf("%v.%v", videoId, "png")
	photolink := ""
	savePath := "assets/"
	if form.Type == "lesson" {
		savePath = fmt.Sprintf("%v%v", savePath, "videoclasses/thumbnails")
		photolink = fmt.Sprintf("%v/view_thumbnail/lesson/%v", productionUrl, filename)
	}
	if form.Type == "career" {
		savePath = fmt.Sprintf("%v%v", savePath, "careertalks/thumbnails")

		err = c.SaveUploadedFile(form.Photo, fmt.Sprintf("%v/%v", savePath, filename))
		if err != nil {
			c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to save photo"})
			return
		}
		c.JSON(201, map[string]string{"url": filename})
		return
	}
	if form.Type == "mkurugenzi" {
		savePath = fmt.Sprintf("%v%v", savePath, "mkurugenzi/thumbnails")

		err = c.SaveUploadedFile(form.Photo, fmt.Sprintf("%v/%v", savePath, filename))
		if err != nil {
			c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to save photo"})
			return
		}
		c.JSON(201, map[string]string{"url": filename})
		return
	}
	if form.Type == "examiner" {
		savePath = fmt.Sprintf("%v%v", savePath, "examinertalks/thumbnails")
		photolink = fmt.Sprintf("%v/view_thumbnail/examiner/%v", productionUrl, filename)
	}

	err = c.SaveUploadedFile(form.Photo, fmt.Sprintf("%v/%v", savePath, filename))
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "unable to save photo"})
		return
	}

	c.JSON(201, map[string]string{"url": photolink})
}

func CheckifFileExistInServer(c *gin.Context) {
	filename := c.Param("filename")
	fmt.Println(filename)

	folders := []string{"thumbnails", "videos"}
	isExist := false

	for _, v := range folders {

		if _, err := os.Stat(fmt.Sprintf("assets/videoclasses/%v/%v", v, filename)); err == nil {
			isExist = true
			break
		}
	}

	if isExist {
		c.JSON(200, map[string]string{"message": "exist"})

	} else {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "file does not exist in server"})

	}

}
func GetAllLessonVideoFilesInTheServer(c *gin.Context) {
	files, err := ioutil.ReadDir("./assets/videoclasses/videos")
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "files not  found"})
		return

	}
	allfiles := []string{}

	for _, f := range files {
		if !f.IsDir() {
			allfiles = append(allfiles, f.Name())
		}
	}
	c.JSON(200, map[string]interface{}{"files": allfiles})

}
