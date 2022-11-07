package helperfunctions

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"

	commentmodel "github.com/Anza2022/Anza_Backend_API/models/comment_model"
	gamifiedquizes "github.com/Anza2022/Anza_Backend_API/models/gamified_quizes"
	notificationmodel "github.com/Anza2022/Anza_Backend_API/models/notification_model"
	pastpapersmodels "github.com/Anza2022/Anza_Backend_API/models/past_papers_models"
	usermodels "github.com/Anza2022/Anza_Backend_API/models/user_models"
	videolessons "github.com/Anza2022/Anza_Backend_API/models/video_lessons"
	"github.com/Anza2022/Anza_Backend_API/services/mongodbapi"
	"github.com/google/uuid"
	"github.com/jinzhu/now"
	"github.com/jordan-wright/email"
	"github.com/kevinburke/twilio-go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ComposeEmailForPassWordReset(useremail string, resetcode string) error {
	//demo email
	e := email.NewEmail()
	e.From = "R Manager <kennedymuthaura99@gmail.com>"
	e.To = []string{useremail}
	// e.ReplyTo=""  add noreply

	e.Subject = "Rental Manager Password Reset"
	e.Text = []byte("Your Password Reset Code is ")
	e.HTML = []byte(fmt.Sprintf("<h2 >%v</h2>", resetcode)) //links goes here to html and are formatted.  also code is here h1 to be bigger and padding
	// e.HTML = []byte(fmt.Sprintf("<a href='https://darasa.co.ke/anzaapi/auth/web_reset/%v'>Fancy HTML is supported, too!</a>", userid)) //links goes here to html and are formatted.  also code is here h1 to be bigger and padding
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "kennedymuthaura99@gmail.com", "014017387570", "smtp.gmail.com"))
	if err != nil {
		return err
	}
	return nil

}

func LogToAFileInServer(text string, level string) {
	file, err := os.OpenFile("serverlogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	logger := log.New(file, strings.ToUpper(fmt.Sprintf("%v\t", level)), log.LstdFlags)
	logger.Println(fmt.Sprintf("\t\t %v", text))

}

func parseTime(timestring string) time.Time {
	t, _ := now.Parse("1999-12-12 12")
	return t

}

func GenUUID() string {
	id := uuid.New()
	return strings.Split(strings.ToUpper(id.String()), "-")[0]
}
func GenResetCode() string {
	rand.Seed(time.Now().UnixNano())
	min := 111111
	max := 999999
	resetcode := strconv.Itoa(rand.Intn(max-min+1) + min)
	return resetcode

}

func GetUserModelWithFilterFromDb(filter map[string]interface{}) (usermodels.UserModel, error) {
	ctx := context.Background()
	var usermodel usermodels.UserModel
	if err := mongodbapi.UsersCollection.FindOne(ctx, filter).Decode(&usermodel); err != nil {
		return usermodel, fmt.Errorf("document does not exist")
	}

	return usermodel, nil

}
func GetVideoLessonModelWithFilterFromDb(filter map[string]interface{}) (videolessons.VideoLessonModel, error) {
	ctx := context.Background()
	var lesson videolessons.VideoLessonModel
	if err := mongodbapi.VideoLessonsCollection.FindOne(ctx, filter).Decode(&lesson); err != nil {
		return lesson, fmt.Errorf("document does not exist")
	}

	return lesson, nil

}
func GetGamifiedQuizModelWithFilterFromDb(filter map[string]interface{}) (gamifiedquizes.GamifiedQuizesModel, error) {
	ctx := context.Background()
	var quiz gamifiedquizes.GamifiedQuizesModel
	if err := mongodbapi.GamifiedQuizesCollection.FindOne(ctx, filter).Decode(&quiz); err != nil {
		return quiz, fmt.Errorf("document does not exist")
	}

	return quiz, nil

}
func GetPastPaperModelWithFilterFromDb(filter map[string]interface{}) (pastpapersmodels.PastPaperModel, error) {
	ctx := context.Background()
	var paper pastpapersmodels.PastPaperModel
	if err := mongodbapi.PastPapersCollection.FindOne(ctx, filter).Decode(&paper); err != nil {
		return paper, fmt.Errorf("document does not exist")
	}

	return paper, nil

}
func GetCommentModelWithFilterFromDb(filter map[string]interface{}) (commentmodel.CommentModel, error) {
	ctx := context.Background()
	var comment commentmodel.CommentModel
	if err := mongodbapi.VideoCommentsCollection.FindOne(ctx, filter).Decode(&comment); err != nil {
		return comment, fmt.Errorf("document does not exist")
	}

	return comment, nil

}

func GetNotificationModelWithFilterFromDb(filter map[string]interface{}) (notificationmodel.Notificationmodel, error) {

	ctx := context.Background()
	var notificationmodel notificationmodel.Notificationmodel
	if err := mongodbapi.NotificationsCollection.FindOne(ctx, filter).Decode(&notificationmodel); err != nil {
		return notificationmodel, fmt.Errorf("document does not exist")
	}
	return notificationmodel, nil

}

func GetStringFromMongoId(id primitive.ObjectID) string {
	return fmt.Sprintf("%v", id)
}

func GetMongoidFromString(id string) primitive.ObjectID {
	stringid, _ := primitive.ObjectIDFromHex(id)
	return stringid
}

func ComposeSendGridEmail(useremail, username, resetcode string) error {
	fmt.Println(os.Getenv("SENDGRID_API_KEY"))
	from := mail.NewEmail("Rentals Manager", "kennedymuthaura99@gmail.com")
	subject := "Password Reset at rentals manager"
	to := mail.NewEmail(username, useremail)
	plainTextContent := "rentals manager help landlords manager houses with ease"
	htmlContent := fmt.Sprintf("<div><strong>the reset code for your password is</strong><h2>%v</h2> </div>", resetcode)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("unabe to send the email")
	} else {
		fmt.Println("no error when sending")
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		return nil
	}
}

func SendPhoneMessageTwillio() error {
	var sid = os.Getenv("TWILLIO_SID_KEY")
	var token = os.Getenv("TWILLIO_AUTH_TOKEN")

	client := twilio.NewClient(sid, token, nil)

	// Send a message
	msg, err := client.Messages.SendMessage(os.Getenv("TWILLIO_SID_KEY"), "+254704285929", " Testing Sms Service For Rentals Manager \n  \n \tSent via go :) ✓", nil)

	if err != nil {
		return err
	} else {
		fmt.Println(msg)
		return nil
	}
}

func GetCurrrentDate() string {
	return strings.Split(time.Now().String(), " ")[0]
}
func GetCurrrentTime() string {
	return strings.Split(strings.Split(time.Now().String(), " ")[1], ".")[0]
}

func CreateNewNotificationAndSaveIt(userId string, message string, category string, time string) bool {
	newNofication := notificationmodel.Notificationmodel{Category: category, Message: message, UserId: userId, CreatedDate: GetCurrrentDate(), CreatedTime: time, IsReadByUser: false}
	_, err := mongodbapi.AddADocumentToCollection(newNofication.ToMap(), mongodbapi.NotificationsCollection)
	if err != nil {
		return false
	} else {
		return true
	}
}

func CreateNewNotificationAndSaveItUsingEmail(email string, message string, category string, time string) bool {

	dbuser, _ := GetUserModelWithFilterFromDb(bson.M{"email": email})

	newNofication := notificationmodel.Notificationmodel{Category: category, Message: message, UserId: dbuser.UserId, CreatedDate: GetCurrrentDate(), CreatedTime: time, IsReadByUser: false}
	_, err := mongodbapi.AddADocumentToCollection(newNofication.ToMap(), mongodbapi.NotificationsCollection)
	if err != nil {
		return false
	} else {
		return true
	}

}

func GetDaysDifferenceFromNow(createdDate string) int64 {

	t1 := time.Now()
	t2, err := now.Parse(createdDate)
	if err != nil {
		fmt.Println("cannot parse provided date")
	}

	days := t1.Sub(t2).Hours() / 24
	d := int64(days)

	return d

}

func DeleteFileInServer(filename string) {
	videoFolders := []string{"assets/videoclasses/videos/", "assets/examinertalks/videos/", "assets/careertalks/videos/"}
	photoFolders := []string{"assets/videoclasses/thumbnails/", "assets/examinertalks/thumbnails/", "assets/careertalks/thumbnails/", "assets/gamifiedtest/thumbnails/"}
	pdffolders := []string{"assets/pastpapers/"}
	ext := strings.Split(filename, ".")[1]

	if strings.ToLower(ext) == "png" {
		for _, v := range photoFolders {
			err := os.Remove(fmt.Sprintf("%v%v", v, filename))
			if err == nil {
				break
			} else {
				continue
			}
		}

	} else if strings.ToLower(ext) == "mp4" || strings.ToLower(ext) == "mkv" {
		for _, v := range videoFolders {
			err := os.Remove(fmt.Sprintf("%v%v", v, filename))
			if err == nil {
				break
			} else {
				continue
			}
		}
	} else if strings.ToLower(ext) == "pdf" {
		for _, v := range pdffolders {
			err := os.Remove(fmt.Sprintf("%v%v", v, filename))
			if err == nil {
				break
			} else {
				continue
			}
		}
	} else {
		fmt.Println("file extension not known")
	}
}
