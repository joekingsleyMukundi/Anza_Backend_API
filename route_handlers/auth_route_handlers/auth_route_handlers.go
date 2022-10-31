package authroutehandlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	authmodels "github.com/kennedy-muthaura/anzaapi/models/auth_models"
	referearnmodel "github.com/kennedy-muthaura/anzaapi/models/refer_earn_model"
	usermodels "github.com/kennedy-muthaura/anzaapi/models/user_models"
	"github.com/kennedy-muthaura/anzaapi/services/mongodbapi"
	"github.com/kennedy-muthaura/anzaapi/utils/appconstants"
	helperfunctions "github.com/kennedy-muthaura/anzaapi/utils/helper_functions"
	tokenhelperfunctions "github.com/kennedy-muthaura/anzaapi/utils/token_helper_functions"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginAUserHandler(c *gin.Context) {

	//logins in the user , returns the user profile with   authtoken and refresh token
	var payload authmodels.LoginPayLoadModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "All Login Fields were not provided"})
		return
	}

	dbuserdoc, err := helperfunctions.GetUserModelWithFilterFromDb(bson.M{"phoneNumber": payload.PhoneNumber})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "user does not exist"})
		return
	}
	if dbuserdoc.IsGoogleSignUp {
		c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "please login using google"})
		return
	}

	// //compare password with hash
	var isValid = helperfunctions.VerifyPasswordHash(payload.Password, dbuserdoc.PassWord)
	if isValid {
		token, err := tokenhelperfunctions.CreateToken(dbuserdoc.UserId)
		refreshToken, _ := tokenhelperfunctions.CreateRefreshToken(dbuserdoc.UserId)
		if err != nil {
			helperfunctions.LogToAFileInServer("required to create Auth Tokens", "ERROR")
			c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Authentication failed, try again later"})
			return
		}
		//get the bonus data
		bounsData, err := mongodbapi.RetrieveOneDocumentInACollection(mongodbapi.ReferBonusesDataCollection, map[string]interface{}{"userId": dbuserdoc.UserId})
		if err != nil {
			helperfunctions.LogToAFileInServer("unable to retrieve bonus data during login", "ERROR")
			c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Authentication failed, try again later"})
			return
		}
		userstats, err := mongodbapi.RetrieveOneDocumentInACollection(mongodbapi.UserStatsCollection, map[string]interface{}{"userId": dbuserdoc.UserId})
		if err != nil {
			helperfunctions.LogToAFileInServer("unable to retrieve user stats data during login", "ERROR")
			c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Authentication failed, try again later"})
			return
		}
		subscriptiondata, err := mongodbapi.RetrieveOneDocumentInACollection(mongodbapi.AccountSubscriptionCollection, map[string]interface{}{"userId": dbuserdoc.UserId})
		if err != nil {
			helperfunctions.LogToAFileInServer("unable to retrieve account subscriptiond ata during login", "ERROR")
			c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Authentication failed, try again later"})
			return
		}

		c.JSON(200, map[string]interface{}{
			"userData":         dbuserdoc,
			"bonusData":        bounsData,
			"subscriptionData": subscriptiondata,
			"userStats":        userstats,
			"access_token":     token,
			"refresh_token":    refreshToken})
		return
		//create acesstoken&refresh token and add them to the map with db user doc, return the json of that user.

	} else {

		c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Incorrect phone number or password"})
		return
	}

}
func RefreshUserHandler(c *gin.Context) {
	userid := c.Param("id")

	dbuserdoc, err := helperfunctions.GetUserModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(userid)})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "user does not exist"})
		return
	}

	// //compare password with hash

	if true {
		token, err := tokenhelperfunctions.CreateToken(dbuserdoc.UserId)
		refreshToken, _ := tokenhelperfunctions.CreateRefreshToken(dbuserdoc.UserId)
		if err != nil {
			helperfunctions.LogToAFileInServer("required to create Auth Tokens", "ERROR")
			c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Authentication failed, try again later"})
			return
		}
		//get the bonus data
		bounsData, err := mongodbapi.RetrieveOneDocumentInACollection(mongodbapi.ReferBonusesDataCollection, map[string]interface{}{"userId": dbuserdoc.UserId})
		if err != nil {
			helperfunctions.LogToAFileInServer("unable to retrieve bonus data during login", "ERROR")
			c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Authentication failed, try again later"})
			return
		}
		userstats, err := mongodbapi.RetrieveOneDocumentInACollection(mongodbapi.UserStatsCollection, map[string]interface{}{"userId": dbuserdoc.UserId})
		if err != nil {
			helperfunctions.LogToAFileInServer("unable to retrieve user stats data during login", "ERROR")
			c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Authentication failed, try again later"})
			return
		}
		subscriptiondata, err := mongodbapi.RetrieveOneDocumentInACollection(mongodbapi.AccountSubscriptionCollection, map[string]interface{}{"userId": dbuserdoc.UserId})
		if err != nil {
			helperfunctions.LogToAFileInServer("unable to retrieve account subscriptiond ata during login", "ERROR")
			c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Authentication failed, try again later"})
			return
		}

		c.JSON(200, map[string]interface{}{
			"userData":         dbuserdoc,
			"bonusData":        bounsData,
			"subscriptionData": subscriptiondata,
			"userStats":        userstats,
			"access_token":     token,
			"refresh_token":    refreshToken})
		return
		//create acesstoken&refresh token and add them to the map with db user doc, return the json of that user.

	} else {

		c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Incorrect email or password"})
		return
	}

}

func RegisterUserHandler(c *gin.Context) {
	//todo create referalbonus, user stats,
	//unique key will be email
	var invitecode = c.Query("inviteCode")

	var userModel usermodels.UserModel

	err := c.ShouldBind(&userModel)

	if err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, gin.H{
			"message": "All User Fields were not Provided/correct",
		})

		return
	}

	duplicateDoc := mongodbapi.UsersCollection.FindOne(context.Background(), bson.M{"phoneNumber": userModel.PhoneNumber})

	if duplicateDoc.Err() == nil {
		c.JSON(appconstants.ErrorStatusCode, gin.H{
			"message": "user with this phone number already exist",
		})

		return

	}
	newHash, _ := helperfunctions.HashPassword(userModel.PassWord)

	userModel.PassWord = newHash
	newUserId, err := mongodbapi.AddADocumentToCollection(userModel.ToMap(), mongodbapi.UsersCollection)

	// userModel.PassWord = ""
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, gin.H{
			"message": "unable to save the user",
		})

		return

	}
	userModel.UserId = newUserId

	if len(invitecode) > 1 {
		ctx := context.Background()
		//get the document, updates the client ids with this, save the document
		var bonusDoc referearnmodel.ReferalBonus
		if err = mongodbapi.ReferBonusesDataCollection.FindOne(ctx, bson.M{"bonusCode": invitecode}).Decode(&bonusDoc); err != nil {
			invitecode = ""

		} else {
			bonusDoc.ReferredUsersIds = append(bonusDoc.ReferredUsersIds, newUserId)

			bonusDoc.AmountEarned += 100

			mongodbapi.UpdateADocInCollection(bonusDoc.ToMap(), mongodbapi.ReferBonusesDataCollection, bonusDoc.Id)

		}
	}
	//create a referer bonus model and save it
	var newReferModel = referearnmodel.ReferalBonus{UserId: newUserId, BonusCode: helperfunctions.GenUUID(), InviteCode: invitecode, AmountEarned: 0, ReferredUsersIds: []string{}, WidthdrawnAmount: 0}
	referid, err := mongodbapi.AddADocumentToCollection(newReferModel.ToMap(), mongodbapi.ReferBonusesDataCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, gin.H{
			"message": "unable to save the user",
		})

		return
	}
	newReferModel.Id = referid

	//create account subscription model and save it
	var newsubs = usermodels.AccountSubscription{SubscriptionId: "", UserId: userModel.UserId, SubscriptionStartDate: helperfunctions.GetCurrrentDate(), SubscriptionEndDate: helperfunctions.GetCurrrentDate(), CurrentSubscriptionPlan: "Free Trial", Resubscriptions: 0, AccountLockDate: helperfunctions.GetCurrrentDate(), IsRefererPaid: false}
	subid, err := mongodbapi.AddADocumentToCollection(newsubs.ToMap(), mongodbapi.AccountSubscriptionCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, gin.H{
			"message": "unable to save the user",
		})

		return
	}
	newsubs.SubscriptionId = subid

	//create account subscription model and save it
	var userstats = usermodels.UserStats{StatsId: "", UserId: userModel.UserId, LastMpesaPhonenumber: "", FavoriteLessonIds: []string{}, TotalVideosWatched: 0, TotalTestTaken: 0, TotalMarksScored: 0}
	statsid, err := mongodbapi.AddADocumentToCollection(userstats.ToMap(), mongodbapi.UserStatsCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, gin.H{
			"message": "unable to save the user",
		})

		return
	}
	userstats.StatsId = statsid

	//generate auth tokens
	accessToken, err := tokenhelperfunctions.CreateToken(newUserId)
	if err != nil {
		helperfunctions.LogToAFileInServer("Unable To create Auth Tokens", "ERROR")
	}
	refreshToken, err := tokenhelperfunctions.CreateRefreshToken(newUserId)
	if err != nil {
		helperfunctions.LogToAFileInServer("Unable To create Auth Tokens", "ERROR")
	}
	c.JSON(201, map[string]interface{}{
		"userData":         userModel,
		"bonusData":        newReferModel,
		"subscriptionData": newsubs,
		"userStats":        userstats,
		"access_token":     accessToken,
		"refresh_token":    refreshToken,
	})

}

func TokenRefreshHandler(c *gin.Context) {
	userid, _ := c.Get("userId")

	accessToken, err := tokenhelperfunctions.CreateToken(userid.(string))
	if err != nil {
		helperfunctions.LogToAFileInServer("Unable To create Auth Tokens", "ERROR")
	}
	refreshToken, err := tokenhelperfunctions.CreateRefreshToken(userid.(string))
	if err != nil {
		helperfunctions.LogToAFileInServer("Unable To create Auth Tokens", "ERROR")
	}
	c.JSON(200, map[string]interface{}{

		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

type ChangePasswordPayload struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

func ChangeUserPassword(c *gin.Context) {
	id := c.Param("userId")

	var payload ChangePasswordPayload

	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "All fields were not provided"})
		return
	}

	dbuser, err := helperfunctions.GetUserModelWithFilterFromDb(map[string]interface{}{"_id": helperfunctions.GetMongoidFromString(id)})

	if err != nil {

		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "user does not exist"})
		return
	}

	if helperfunctions.VerifyPasswordHash(payload.OldPassword, dbuser.PassWord) {
		newHash, err := helperfunctions.HashPassword(payload.NewPassword)
		if err != nil {
			c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "hash failedd,try again later"})
			return
		}

		dbuser.PassWord = newHash

		_, err = mongodbapi.UpdateADocInCollection(dbuser.ToMap(), mongodbapi.UsersCollection, dbuser.UserId)
		if err != nil {
			c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "failed to update user"})
			return
		} else {
			c.JSON(200, map[string]string{"hash": newHash})
			return
		}
	} else {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "old password is Wrong"})
		return
	}

}

func ResetPasswordWithCodeHandler(c *gin.Context) {

	var payload authmodels.PasswordResetWithCodePayloadModel

	err := c.ShouldBind(&payload)
	if err != nil {

		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "All Password Recovery Fields were not provides"})
		return

	}
	ctx := context.Background()
	//get the request from database and user data for
	var requestDoc authmodels.PassWordResetDbModel
	if err := mongodbapi.PassWordResetRequestsCollection.FindOne(ctx, bson.M{"email": payload.Email}).Decode(&requestDoc); err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "reset code for this email does not exist"})
		return
	}
	if requestDoc.ExpiresAt < time.Now().Unix() {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "reset code is expired"})
		return

	}
	if requestDoc.ResetCode != payload.ResetCode {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "invalid reset code"})
		return

	}

	var dbuser usermodels.UserModel
	if err := mongodbapi.UsersCollection.FindOne(ctx, bson.M{"email": requestDoc.Email}).Decode(&dbuser); err != nil {

		fmt.Println("error getting  user in db from email in password request doc")
		return
	}

	newhash, err := helperfunctions.HashPassword(payload.NewPassword)
	if err != nil {
		helperfunctions.LogToAFileInServer("unable to hash the new password ", "ERROR")
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "failed: try again later"})

		return
	}

	dbuser.PassWord = newhash
	_, err = mongodbapi.UpdateADocInCollection(dbuser.ToMap(), mongodbapi.UsersCollection, dbuser.UserId)

	if err != nil {
		if appconstants.IsDebug {
			fmt.Println("unable to  update user with the new password")
		}
		helperfunctions.LogToAFileInServer("unable to update user with new password from reset code", "ERROR")

		return
	}

	_, _ = mongodbapi.PassWordResetRequestsCollection.DeleteOne(context.Background(), map[string]interface{}{"email": dbuser.Email})

	c.JSON(200, map[string]string{"message": "password reset success"})

}
func SendPasswordResetCodeToMailHander(c *gin.Context) {
	//send a html email with code for reset password

	var resetModel authmodels.PasswordResetPayloadModel

	if err := c.ShouldBind(&resetModel); err != nil {
		fmt.Println(err)
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "all fields for password reset were not provided"})
		return
	}
	ctx := context.Background()
	var dbuser usermodels.UserModel
	if err := mongodbapi.UsersCollection.FindOne(ctx, bson.M{"email": resetModel.Email}).Decode(&dbuser); err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "user with this email does not exist"})
		return
	}
	resetcode := helperfunctions.GenResetCode()

	var newResetDbModel = authmodels.PassWordResetDbModel{
		UserId:    dbuser.UserId,
		Email:     dbuser.Email,
		ResetCode: resetcode,
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
	}

	err := helperfunctions.ComposeSendGridEmail(dbuser.Email, dbuser.UserName, resetcode)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "Failed sending email: try again later"})
		return
	}
	_, _ = mongodbapi.PassWordResetRequestsCollection.DeleteOne(context.Background(), map[string]interface{}{"email": dbuser.Email})

	_, err = mongodbapi.AddADocumentToCollection(newResetDbModel.ToMap(), mongodbapi.PassWordResetRequestsCollection)
	if err != nil {
		if appconstants.IsDebug {
			fmt.Println("unable to save new password request to db")
		}
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "Failed : try again later"})

		return
	}

	c.JSON(200, map[string]string{"message": "Password reset code was sent successfully"})
}

//not in use now, update them when you want their functionality

func SendResetPasswordMailWeb(c *gin.Context) {

	var payload authmodels.PasswordResetPayloadModel

	err := c.ShouldBind(&payload)
	if err != nil {

		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "All Password Recovery Fields were not provides"})
		return

	}

	//get user
	var newRequest authmodels.PassWordResetDbModel = authmodels.PassWordResetDbModel{Email: payload.Email, ExpiresAt: time.Now().Add(time.Hour * 24).Unix()}
	newId, err := mongodbapi.AddADocumentToCollection(newRequest.ToMap(), mongodbapi.PassWordResetRequestsCollection)
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "Service Not available, try again later"})
		return

	}
	newRequest.RequestId = newId

	//Todo compose a new email without reset code
	helperfunctions.ComposeEmailForPassWordReset(newId, "no code")

	//send a html email template with a link to reset password and expiration date.
	c.JSON(200, map[string]string{"message": "email sent successfully"})

}

func SendResetPasswordMailLink(c *gin.Context) {
	//send a html email template with a link to reset password and expiration date.
	var resetModel authmodels.PasswordResetPayloadModel

	if err := c.ShouldBind(&resetModel); err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "all fields for password reset were not provided"})
		return
	}
	ctx := context.Background()
	var dbuser usermodels.UserModel
	if err := mongodbapi.PassWordResetRequestsCollection.FindOne(ctx, bson.M{"email": resetModel.Email}).Decode(&dbuser); err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]string{"message": "user with this email does not exist"})
		return
	}
	resetcode := helperfunctions.GenUUID()

	var newResetDbModel = authmodels.PassWordResetDbModel{
		UserId:    dbuser.UserId,
		Email:     dbuser.Email,
		ResetCode: resetcode,
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
	}

	_, err := mongodbapi.AddADocumentToCollection(newResetDbModel.ToMap(), mongodbapi.PassWordResetRequestsCollection)
	if err != nil {
		fmt.Println("unable to save new password request to db")
		return
	}

	helperfunctions.ComposeEmailForPassWordReset(dbuser.Email, resetcode)

	c.JSON(200, map[string]string{"message": "Password reset link was sent successfully"})
}
func GetTemplateForPassWordReset(c *gin.Context) {
	//send a html email template with a link to reset password and expiration date.
	//returns the password reset form  template
	//takes the parameter and gets  the details from db
	requestuserId := c.Param("userId")
	ctx := context.Background()
	var dbrequest authmodels.PassWordResetDbModel
	if err := mongodbapi.PassWordResetRequestsCollection.FindOne(ctx, bson.M{"userId": requestuserId}).Decode(&dbrequest); err != nil {

		fmt.Println("error getting  password request in db from userid")
		return
	}
	if dbrequest.ExpiresAt < time.Now().Unix() {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"name": strings.Split(dbrequest.Email, "@")[0],
			"year": time.Now().Year(),
		})
	} else {
		c.HTML(http.StatusOK, "expired_link_template.tmpl", gin.H{
			"name": strings.Split(dbrequest.Email, "@")[0],
			"year": time.Now().Year(),
		})
	}

	c.JSON(200, map[string]string{"message": "tThis is the route to get template to using email link"})
}

func ReceiveNewPasswordFromWeb(c *gin.Context) {
	requestId := c.Param("linkid")
	var newpass NewPassWordModel

	if err := c.ShouldBind(&newpass); err != nil {
		fmt.Println("the form cannot bind")
		return
	}
	ctx := context.Background()
	//get the request from database and user data for
	var requestDoc authmodels.PassWordResetDbModel
	if err := mongodbapi.PassWordResetRequestsCollection.FindOne(ctx, bson.M{"_id": requestId}).Decode(&requestDoc); err != nil {

		fmt.Println("error getting  request doc")
		return
	}
	var dbuser usermodels.UserModel
	if err := mongodbapi.PassWordResetRequestsCollection.FindOne(ctx, bson.M{"email": requestDoc.Email}).Decode(&dbuser); err != nil {

		fmt.Println("error getting  user in db from email in password request doc")
		return
	}

	if newpass.ConfirmNewPassWord != newpass.NewPassword {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"name":      dbuser.UserName,
			"year":      time.Now().Year(),
			"dontmatch": "Failed: Your new password dont match confirm password",
		})
		return
	}

	newhash, err := helperfunctions.HashPassword(newpass.NewPassword)
	if err != nil {
		fmt.Println("unable to hash the new password ")
	}

	dbuser.PassWord = newhash
	_, err = mongodbapi.UpdateADocInCollection(dbuser.ToMap(), mongodbapi.UsersCollection, dbuser.UserId)

	if err != nil {
		fmt.Println("unable to  update user with the new password")
		return
	}

	_, _ = mongodbapi.PassWordResetRequestsCollection.DeleteOne(context.Background(), map[string]interface{}{"userId": dbuser.UserId})
	c.HTML(http.StatusOK, "success_password_reset.tmpl", gin.H{
		"name": strings.Split(dbuser.Email, "@")[0],
		"year": time.Now().Year(),
	})

}

//social apps

type googleloginpayload struct {
	Email string `json:"email" binding:"required"`
	Token string `json:"token" binding:"required"`
}

func GoogleLoginAUserHandler(c *gin.Context) {
	var payload googleloginpayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "All Login Fields were not provided"})
		return
	}
	//todo validate google token
	if len(payload.Token) < 10 {
		c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Invalid Creditials"})
		return

	}

	dbuserdoc, err := helperfunctions.GetUserModelWithFilterFromDb(map[string]interface{}{"email": payload.Email})
	if err != nil {
		c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "user does not exist"})
		return
	}
	token, err := tokenhelperfunctions.CreateToken(dbuserdoc.UserId)
	refreshToken, _ := tokenhelperfunctions.CreateRefreshToken(dbuserdoc.UserId)
	if err != nil {
		helperfunctions.LogToAFileInServer("required to create Auth Tokens", "ERROR")
		c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Authentication failed, try again later"})
		return
	}
	//get the bonus data
	bounsData, err := mongodbapi.RetrieveOneDocumentInACollection(mongodbapi.ReferBonusesDataCollection, map[string]interface{}{"userId": dbuserdoc.UserId})
	if err != nil {
		helperfunctions.LogToAFileInServer("unable to retrieve bonus data during login", "ERROR")
		c.JSON(appconstants.ErrorStatusCode, map[string]interface{}{"message": "Authentication failed, try again later"})
		return
	}

	c.JSON(200, map[string]interface{}{
		"userData":      dbuserdoc,
		"bonusData":     bounsData,
		"access_token":  token,
		"refresh_token": refreshToken})

}

type NewPassWordModel struct {
	NewPassword        string `form:"newpassword"`
	ConfirmNewPassWord string `form:"confirmnewpassword"`
}
