package paymentsroutehandlers

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/websocket"
// 	"github.com/jwambugu/mpesa-golang-sdk"
// 	"github.com/jwambugu/mpesa-golang-sdk/pkg/config"
// 	onlineusersmap "github.com/kennedy-muthaura/anzaapi/controllers/online_users_map"
// 	paymentmodels "github.com/kennedy-muthaura/anzaapi/models/payment_models"
// 	referearnmodel "github.com/kennedy-muthaura/anzaapi/models/refer_earn_model"
// 	transactionsmodel "github.com/kennedy-muthaura/anzaapi/models/transactions_model"
// 	usermodels "github.com/kennedy-muthaura/anzaapi/models/user_models"
// 	"github.com/kennedy-muthaura/anzaapi/services/mongodbapi"
// 	"github.com/kennedy-muthaura/anzaapi/utils/appconstants"
// 	helperfunctions "github.com/kennedy-muthaura/anzaapi/utils/helper_functions"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// } // use default options

// func getBaseCallbackUrl() string {
// 	var baseCallbackUrl string
// 	if appconstants.IsDebug {
// 		// baseCallbackUrl = "ngrok route here"
// 		baseCallbackUrl = "https://api.darasa.co.ke"

// 	} else {
// 		baseCallbackUrl = "https://api.darasa.co.ke"
// 	}

// 	return baseCallbackUrl
// }

// //notify users with another go routing

// var socketConnections = map[string]*websocket.Conn{}
// var paymentsRequestsConnections = map[string]paymentmodels.PaymentRequestPayload{}
// var withdrawRequestConnections = map[string]paymentmodels.WithdrawRequestPayload{}

// func ProcessMpesPaymentRequestHandler(c *gin.Context) {

// 	amount, _ := strconv.ParseUint(c.Query("amount"), 10, 64)
// 	phoneNumber, _ := strconv.ParseUint(c.Query("phoneNumber"), 10, 64)
// 	propertyId := c.Query("propertyId")
// 	userId := c.Query("userId")
// 	inviteCode := c.Query("inviteCode")
// 	plan := c.Query("plan")
// 	months := c.Query("months")
// 	time := c.Query("time")

// 	var payload paymentmodels.PaymentRequestPayload = paymentmodels.PaymentRequestPayload{PhoneNumber: phoneNumber, Amount: amount, PropertyId: propertyId, InviteCode: inviteCode, Months: months, Plan: plan, UserId: userId, Time: time}

// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Print("upgrade:", err)
// 		return
// 	}

// 	// conf, err := config.Get()
// 	if err != nil {
// 		log.Fatalln(err)
// 		return
// 	}

// 	conn.WriteMessage(1, []byte("Connection Established"))
// 	mpesaApp := mpesa.Init(&config.Credentials{ConsumerKey: os.Getenv("MPESA_C2B_CONSUMERKEY"), ConsumerSecret: os.Getenv("MPESA_C2B_CONSUMERSECRET")}, true)
// 	mpesaApp.B2CPayment(&mpesa.B2CPaymentRequest{})
// 	response, err := mpesaApp.LipaNaMpesaOnline(&mpesa.STKPushRequest{
// 		Shortcode: 470470,
// 		PartyB:    470470,
// 		Passkey:   os.Getenv("MPESA_C2B_PASSKEY"),

// 		Amount:                 payload.Amount,
// 		PhoneNumber:            payload.PhoneNumber,
// 		ReferenceCode:          "Rmanager",
// 		TransactionDescription: "Rmanager",
// 		CallbackURL:            fmt.Sprintf("%v/anzaapi/payments/mpesa/payment_callback", getBaseCallbackUrl()), // Add your callback URL here
// 		TransactionType:        "CustomerPayBillOnline",                                                         // CustomerPayBillOnline or CustomerBuyGoodsOnline
// 	})
// 	if err != nil {

// 		helperfunctions.LogToAFileInServer("Failed To Make StkPush", "ERROR")
// 		conn.WriteMessage(1, []byte("failed to  make stk push error"))
// 		conn.WriteMessage(1, []byte("payment failed"))

// 		if appconstants.IsDebug {
// 			fmt.Println(err)
// 		}

// 		return
// 	}

// 	// Check if the request was successful
// 	if response.IsSuccessful {
// 		// Handle your successful logic here
// 		conn.WriteMessage(1, []byte("Stk push Sucess"))
// 		conn.WriteMessage(1, []byte("Check you phone to Complete Payment"))

// 		socketConnections[response.CheckoutRequestID] = conn
// 		paymentsRequestsConnections[response.CheckoutRequestID] = payload
// 	} else {
// 		conn.WriteMessage(1, []byte("failed to  make stk push-- safaricom"))
// 		conn.WriteMessage(1, []byte(response.ErrorMessage))

// 		conn.WriteMessage(1, []byte("payment failed"))

// 	}

// 	for {
// 		_, message, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("read:", err)
// 			break
// 		}
// 		log.Printf("recv: %s", message)

// 		err = conn.WriteMessage(1, message)
// 		if string(message) == "close" {
// 			conn.Close()
// 			delete(socketConnections, response.CheckoutRequestID)
// 			delete(paymentsRequestsConnections, response.CheckoutRequestID)

// 		}
// 		if err != nil {
// 			log.Println("write:", err)
// 			break
// 		}
// 	}
// }

// func MpesaPaymentrequestSafCallbackHandler(c *gin.Context) {

// 	var res mpesa.LipaNaMpesaOnlineCallback

// 	err := c.ShouldBind(&res)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var conn = socketConnections[res.Body.StkCallback.CheckoutRequestID]
// 	var originalRequestPayload = paymentsRequestsConnections[res.Body.StkCallback.CheckoutRequestID]

// 	if res.Body.StkCallback.ResultCode == 0 {

// 		conn.WriteMessage(1, []byte(res.Body.StkCallback.ResultDesc))
// 		conn.WriteMessage(1, []byte("updating .."))

// 		property, err := helperfunctions.GetPropertyModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(originalRequestPayload.PropertyId)})
// 		if err != nil {
// 			helperfunctions.LogToAFileInServer("Error getting property  from db during saf callback paymnet", "ERROR")
// 			if appconstants.IsDebug {
// 				fmt.Println("error getting  property doc")
// 			}
// 		} else {

// 			intmonths, _ := strconv.Atoi(originalRequestPayload.Months)

// 			var subscription = usermodels.AccountSubscription{
// 				SubscriptionStartDate: strings.Split(time.Now().String(), " ")[0], AccountLockDate: strings.Split(time.Now().Add(time.Hour*24*31*time.Duration(intmonths)+(time.Hour*24*10)).String(), " ")[0], SubscriptionEndDate: strings.Split(time.Now().Add(time.Hour*24*31*time.Duration(intmonths)).String(), " ")[0], CurrentSubscriptionPlan: fmt.Sprintf("%v Months Plan", originalRequestPayload.Months), Resubscriptions: property.AccountSubscription.Resubscriptions + 1}

// 			property.AccountSubscription = subscription
// 			_, err := mongodbapi.UpdateADocInCollection(property.ToMap(), mongodbapi.AllPropertiesCollection, property.PropertyId)
// 			if err != nil {
// 				helperfunctions.LogToAFileInServer("Unable to update document subscription after payment", "ERROR")
// 			}

// 			helperfunctions.CreateNewNotificationAndSaveIt(property.OwnerId, fmt.Sprintf("You Subscription Renewal was successful, next payment with  be %v", subscription.SubscriptionEndDate), "Subscriptions", originalRequestPayload.Time)
// 			if len(property.InvitedUsers) > 0 {
// 				for _, v := range property.InvitedUsers {
// 					//send
// 					helperfunctions.CreateNewNotificationAndSaveItUsingEmail(v.UserEmail, fmt.Sprintf("You Subscription Renewal was successful, next payment with  be %v", subscription.SubscriptionEndDate), "Subscriptions", originalRequestPayload.Time)

// 					onlineusersmap.SendRefreshMessageToUserWithEmail(v.UserEmail)

// 				}
// 			}

// 		}

// 		userDoc, _ := helperfunctions.GetUserModelWithFilterFromDb(map[string]interface{}{"_id": helperfunctions.GetMongoidFromString(originalRequestPayload.UserId)})

// 		//pay the market if not payed
// 		if !userDoc.IsRefererPaid {
// 			if len(originalRequestPayload.InviteCode) > 1 {
// 				ctx := context.Background()

// 				//get the document, updates the client ids with this, save the document
// 				var bonusDoc referearnmodel.ReferalBonus
// 				if err = mongodbapi.AllReferBonusesDataCollection.FindOne(ctx, bson.M{"bonusCode": originalRequestPayload.InviteCode}).Decode(&bonusDoc); err != nil {
// 					fmt.Println("error getting  bonus doc")
// 				} else {
// 					bonusDoc.PaidAmount += 100

// 					mongodbapi.UpdateADocInCollection(bonusDoc.ToMap(), mongodbapi.AllReferBonusesDataCollection, bonusDoc.Id)
// 					if appconstants.IsDebug {
// 						fmt.Println("referer paid amount added with one hundred")
// 					}
// 					userDoc.IsRefererPaid = true
// 				}

// 			}

// 		}
// 		newid, err := mongodbapi.UpdateADocInCollection(userDoc.ToMap(), mongodbapi.AllUsersCollection, userDoc.UserId)
// 		if err != nil {
// 			if appconstants.IsDebug {
// 				fmt.Println("unable to update the user")
// 			}
// 			if appconstants.IsDebug {
// 				fmt.Println(newid)
// 			}
// 			return
// 		}
// 		if appconstants.IsDebug {
// 			fmt.Println("user updated after a successful payment")
// 		}
// 		conn.WriteMessage(1, []byte("payment success"))

// 	} else {

// 		conn.WriteMessage(1, []byte(res.Body.StkCallback.ResultDesc))
// 		// time.Sleep(time.Millisecond * 600)
// 		conn.WriteMessage(1, []byte("payment failed"))
// 	}

// }

// func ProcessMpesawWthdrawRequestHandler(c *gin.Context) {

// 	amount, _ := strconv.ParseUint(c.Query("amount"), 10, 64)
// 	phoneNumber, _ := strconv.ParseUint(c.Query("phoneNumber"), 10, 64)
// 	userId := c.Query("userId")
// 	time := c.Query("time")
// 	fromWallet, _ := strconv.ParseBool(c.Query("fromWallet"))

// 	var originalWithdrawRequest paymentmodels.WithdrawRequestPayload = paymentmodels.WithdrawRequestPayload{PhoneNumber: phoneNumber, FromWallet: fromWallet, Amount: amount, UserId: userId, Time: time}
// 	bonusDoc, _ := helperfunctions.GetBonusModelWithFilterFromDb(map[string]interface{}{"userId": originalWithdrawRequest.UserId})

// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Print("upgrade:", err)
// 		return
// 	}

// 	// conf, err := config.Get()
// 	if err != nil {
// 		log.Fatalln(err)
// 		return
// 	}
// 	conn.WriteMessage(1, []byte("Connection Established"))
// 	if fromWallet {
// 		wallet, _ := helperfunctions.GetWalletModelWithFilterFromDb(map[string]interface{}{"propertyId": originalWithdrawRequest.PropertyId})
// 		if wallet.AvailableAmount < int64(originalWithdrawRequest.Amount) {

// 			conn.WriteMessage(1, []byte("failed: amount requests exceeds your wallet available amount"))

// 			conn.WriteMessage(1, []byte("payment failed"))
// 			return
// 		}

// 	} else {
// 		if bonusDoc.PaidAmount <= int64(amount) {

// 			conn.WriteMessage(1, []byte("failed: amount requests exceeds your available amount"))

// 			conn.WriteMessage(1, []byte("payment failed"))
// 			return

// 		}
// 	}
// 	mpesaApp := mpesa.Init(&config.Credentials{ConsumerKey: os.Getenv("MPESA_C2B_CONSUMERKEY"), ConsumerSecret: os.Getenv("MPESA_C2B_CONSUMERSECRET")}, true)
// 	//make sure this information is correct
// 	res, err := mpesaApp.B2CPayment(&mpesa.B2CPaymentRequest{InitiatorName: os.Getenv("MPESA_B2C_INTIATORNAME"), InitiatorPassword: os.Getenv("MPESA_B2C_INTIATORPASSWORD"),
// 		CommandID: os.Getenv("MPESA_B2C_COMMANDID"), Amount: originalWithdrawRequest.Amount, Shortcode: 9215085, PhoneNumber: originalWithdrawRequest.PhoneNumber, ResultURL: fmt.Sprintf("%v/anzaapi/payments/withdraw_callback", getBaseCallbackUrl()), QueueTimeOutURL: fmt.Sprintf("%v/anzaapi/payments/withdraw_callback", getBaseCallbackUrl()), Occasion: "Bonus", Remarks: "Congrats"})
// 	if err != nil {

// 		helperfunctions.LogToAFileInServer("Failed To Withdraw Request", "ERROR")
// 		conn.WriteMessage(1, []byte("failed to  make widthdraw request"))
// 		conn.WriteMessage(1, []byte("payment failed"))

// 		if appconstants.IsDebug {
// 			fmt.Println(err)
// 		}

// 		return
// 	}

// 	if res.IsSuccessful {
// 		// Handle your successful logic here
// 		conn.WriteMessage(1, []byte("Request Received"))
// 		conn.WriteMessage(1, []byte("Sending Money to your phone"))

// 		socketConnections[res.ConversationId] = conn
// 		withdrawRequestConnections[res.ConversationId] = originalWithdrawRequest
// 	} else {
// 		conn.WriteMessage(1, []byte("failed to start transaction"))
// 		conn.WriteMessage(1, []byte(res.ErrorMessage))

// 		conn.WriteMessage(1, []byte("payment failed"))

// 	}

// 	for {
// 		_, message, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("read:", err)
// 			break
// 		}
// 		log.Printf("recv: %s", message)

// 		err = conn.WriteMessage(1, message)
// 		if string(message) == "close" {
// 			conn.Close()
// 			delete(socketConnections, res.ConversationId)
// 			delete(paymentsRequestsConnections, res.ConversationId)

// 		}
// 		if err != nil {
// 			log.Println("write:", err)
// 			break
// 		}
// 	}

// }

// func MpesaWithdrawRequestSafCallbackHandler(c *gin.Context) {

// 	var res mpesa.B2CPaymentRequestCallback

// 	err := c.ShouldBind(&res)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var conn = socketConnections[res.Result.ConversationID]

// 	if res.Result.ResultCode == 0 {
// 		if originalRequestPayload, ok := withdrawRequestConnections[res.Result.ConversationID]; ok {

// 			conn.WriteMessage(1, []byte(res.Result.ResultDesc))
// 			conn.WriteMessage(1, []byte("updating .."))
// 			var balance int64 = 0

// 			if originalRequestPayload.FromWallet {

// 				wallet, _ := helperfunctions.GetWalletModelWithFilterFromDb(map[string]interface{}{"propertyId": originalRequestPayload.PropertyId})
// 				wallet.WithdrawnAmount += int64(originalRequestPayload.Amount)
// 				wallet.AvailableAmount -= int64(originalRequestPayload.Amount)
// 				mongodbapi.UpdateADocInCollection(wallet.ToMap(), mongodbapi.MpesaWalletsCollection, wallet.WalletId)
// 			} else {

// 				bonusDoc, err := helperfunctions.GetBonusModelWithFilterFromDb(map[string]interface{}{"userId": originalRequestPayload.UserId})
// 				if err != nil {
// 					if appconstants.IsDebug {
// 						fmt.Println("cannot find the bonus data so no updates were made")
// 					}

// 				} else {
// 					bonusDoc.WidthdrawnAmount = bonusDoc.WidthdrawnAmount + int64(originalRequestPayload.Amount)
// 					bonusDoc.PaidAmount = bonusDoc.PaidAmount - int64(originalRequestPayload.Amount)
// 					balance = bonusDoc.PaidAmount

// 				}

// 				newid, err := mongodbapi.UpdateADocInCollection(bonusDoc.ToMap(), mongodbapi.AllReferBonusesDataCollection, bonusDoc.Id)
// 				if err != nil {
// 					if appconstants.IsDebug {
// 						fmt.Println("unable to update the bonus doc")
// 					}
// 					if appconstants.IsDebug {
// 						fmt.Println(newid)
// 					}
// 					return
// 				}
// 				if appconstants.IsDebug {
// 					fmt.Println("bonus updated after a successful withdrawal")
// 				}
// 			}
// 			conn.WriteMessage(1, []byte("payment success"))
// 			helperfunctions.CreateNewNotificationAndSaveIt(originalRequestPayload.UserId, fmt.Sprintf("You withdraw from %v was successful, current balance %v", originalRequestPayload.Amount, balance), "withdrawal", originalRequestPayload.Time)

// 			onlineusersmap.SendRefreshMessageToUserWithId(originalRequestPayload.UserId)
// 		} else {
// 			go helperfunctions.LogToAFileInServer("Payment Withdraw Connection Not Found So Nothing was updated", "ERROR")
// 			bonusDoc, err := helperfunctions.GetBonusModelWithFilterFromDb(map[string]interface{}{"userId": originalRequestPayload.UserId})
// 			if err != nil {
// 				if appconstants.IsDebug {
// 					fmt.Println("cannot find the bonus data so no updates were made")
// 				}

// 			} else {
// 				bonusDoc.WidthdrawnAmount = bonusDoc.WidthdrawnAmount + int64(originalRequestPayload.Amount)
// 				bonusDoc.PaidAmount = bonusDoc.PaidAmount - int64(originalRequestPayload.Amount)

// 				newid, err := mongodbapi.UpdateADocInCollection(bonusDoc.ToMap(), mongodbapi.AllReferBonusesDataCollection, bonusDoc.Id)
// 				if err != nil {
// 					if appconstants.IsDebug {
// 						fmt.Println("unable to update the bonus doc")
// 					}
// 					if appconstants.IsDebug {
// 						fmt.Println(newid)
// 					}
// 					return
// 				}
// 			}

// 		}

// 	} else {

// 		conn.WriteMessage(1, []byte(res.Result.ResultDesc))
// 		// time.Sleep(time.Millisecond * 600)
// 		conn.WriteMessage(1, []byte("payment failed"))
// 	}

// 	//transaction id is the code sent to customer for verification

// 	//update the user bonus doc when success and  save it to db
// 	//return mpesa code

// 	// 	res := map[string]interface{}{
// 	//    "Result": {
// 	//       "ResultType": 0,
// 	//       "ResultCode": 0,
// 	//       "ResultDesc": "The service request is processed successfully.",
// 	//       "OriginatorConversationID": "10571-7910404-1",
// 	//       "ConversationID": "AG_20191219_00004e48cf7e3533f581",
// 	//       "TransactionID": "NLJ41HAY6Q",
// 	//       "ResultParameters": {
// 	//          "ResultParameter": [
// 	//           {
// 	//              "Key": "TransactionAmount",
// 	//              "Value": 10
// 	//           },
// 	//           {
// 	//              "Key": "TransactionReceipt",
// 	//              "Value": "NLJ41HAY6Q"
// 	//           },
// 	//           {
// 	//              "Key": "B2CRecipientIsRegisteredCustomer",
// 	//              "Value": "Y"
// 	//           },
// 	//           {
// 	//              "Key": "B2CChargesPaidAccountAvailableFunds",
// 	//              "Value": -4510.00
// 	//           },
// 	//           {
// 	//              "Key": "ReceiverPartyPublicName",
// 	//              "Value": "254708374149 - John Doe"
// 	//           },
// 	//           {
// 	//              "Key": "TransactionCompletedDateTime",
// 	//              "Value": "19.12.2019 11:45:50"
// 	//           },
// 	//           {
// 	//              "Key": "B2CUtilityAccountAvailableFunds",
// 	//              "Value": 10116.00
// 	//           },
// 	//           {
// 	//              "Key": "B2CWorkingAccountAvailableFunds",
// 	//              "Value": 900000.00
// 	//           }
// 	//         ]
// 	//       },
// 	//       "ReferenceData": {
// 	//          "ReferenceItem": {
// 	//             "Key": "QueueTimeoutURL",
// 	//             "Value": "https:\/\/internalsandbox.safaricom.co.ke\/mpesa\/b2cresults\/v1\/submit"
// 	//           }
// 	//   }
// 	//    }
// 	// }

// 	c.String(200, "thisis the route to handle saf callback for mpesa withdraw")
// }

// func RemoveAdsPaymentHandler(c *gin.Context) {
// 	//sets the isproaccount to true when payment is complete charges 200 per year

// 	c.String(200, "this is the route to process the mpesa widthdraw requests")

// }

// func RemoveAdsSafCallback(c *gin.Context) {
// 	c.String(200, "thisis the route to process the mpesa widthdraw requests")

// }

// func ProcessPaypalPaymentRequest(c *gin.Context) {
// 	c.String(200, "this is the route for handling paypal payments")
// }

// // this
// // is
// // for
// // the
// // tenant
// // payments
// // routes
// // with
// // mpesa
// // online

// //tenant payment handlers
// var tenantSocketConnections = map[string]*websocket.Conn{}
// var tenantPaymentsRequestsConnections = map[string]paymentmodels.TenantMpesaPaymentRequestPayload{}

// func ProcessTenantMpesPaymentRequestHandler(c *gin.Context) {
// 	// check if the property uses internal wallet or has mpesa key
// 	// pay referrer
// 	//
// 	amount, err := strconv.ParseUint(c.Query("amount"), 10, 64)
// 	if err != nil {
// 		if appconstants.IsDebug {
// 			fmt.Println(err)
// 		}
// 		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": " failed to get amount in query"})
// 		return
// 	}
// 	phoneNumber, err := strconv.ParseUint(c.Query("phoneNumber"), 10, 64)
// 	if err != nil {
// 		if appconstants.IsDebug {
// 			fmt.Println(err)
// 		}
// 		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": " failed to get phoneNumber query "})
// 		return
// 	}
// 	propertyId := c.Query("propertyId")
// 	houseNumber := c.Query("houseNumber")
// 	tenantid := c.Query("tenantId")
// 	payerName := c.Query("payerName")
// 	payTime := c.Query("payTime")
// 	isFullPayment, err := strconv.ParseBool(c.Query("isFullPayment"))
// 	if err != nil {
// 		if appconstants.IsDebug {
// 			fmt.Println(err)
// 		}
// 		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": " failed to get isFullPayment query "})
// 		return
// 	}
// 	isCleared, err := strconv.ParseBool(c.Query("isCleared"))
// 	if err != nil {
// 		if appconstants.IsDebug {
// 			fmt.Println(err)
// 		}
// 		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": " failed to get isCleared query "})
// 		return
// 	}
// 	billName := c.Query("billName")
// 	remainingAmount, _ := strconv.ParseUint(c.Query("remainingAmount"), 10, 64)
// 	requiredAmount, err := strconv.ParseUint(c.Query("requiredAmount"), 10, 64)

// 	if err != nil {
// 		if appconstants.IsDebug {
// 			fmt.Println(err)
// 		}
// 		c.JSON(appconstants.ErrorStatusCode, bson.M{"message": " failed to get all query parameters"})
// 		return
// 	}

// 	var payload paymentmodels.TenantMpesaPaymentRequestPayload = paymentmodels.TenantMpesaPaymentRequestPayload{PhoneNumber: phoneNumber, TenantId: tenantid, Amount: amount, PropertyId: propertyId, HouseNumber: houseNumber,
// 		BillName: billName, RemainingAmount: remainingAmount, RequiredAmount: requiredAmount, IsFullPayment: isFullPayment, IsCleared: isCleared, PayerName: payerName, PayerTime: payTime,
// 	}

// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Print("upgrade:", err)
// 		return
// 	}

// 	// conf, err := config.Get()
// 	if err != nil {
// 		log.Fatalln(err)
// 		return
// 	}

// 	conn.WriteMessage(1, []byte("Connection Established"))
// 	rental, err := helperfunctions.GetPropertyModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(payload.PropertyId)})
// 	if err != nil {
// 		fmt.Println("property  is not found")
// 	}
// 	//validate if property is enabled for mpesa payment

// 	//get apartment and get mpesa details  should be encrpted
// 	//mpesa details
// 	var consumerkey string
// 	var consumersecret string
// 	var mpesapasskey string
// 	var shortCode uint
// 	var transactionType string
// 	var description string

// 	if rental.IsInternalMpesaWallet {

// 		transactionType = "CustomerPayBillOnline"
// 		shortCode = 470470
// 		consumerkey = os.Getenv("MPESA_C2B_CONSUMERKEY")
// 		consumersecret = os.Getenv("MPESA_C2B_CONSUMERSECRET")
// 		mpesapasskey = os.Getenv("MPESA_C2B_PASSKEY")
// 		description = "Rmanager"

// 	} else {
// 		if rental.MpesaDetails.UsesPaybill {
// 			transactionType = "CustomerPayBillOnline"
// 		} else {
// 			transactionType = "CustomerBuyGoodsOnline"

// 		}
// 		consumerkey = rental.MpesaDetails.ConsumerKey
// 		consumersecret = rental.MpesaDetails.ConsumerSecret
// 		mpesapasskey = rental.MpesaDetails.PassKey
// 		if rental.MpesaDetails.UsesPaybill {
// 			shortCode = rental.MpesaDetails.Paybill
// 		} else {
// 			shortCode = rental.MpesaDetails.TillNumber

// 		}
// 		description = rental.MpesaDetails.StkDescription
// 	}

// 	mpesaApp := mpesa.Init(&config.Credentials{ConsumerKey: consumerkey, ConsumerSecret: consumersecret}, true)

// 	response, err := mpesaApp.LipaNaMpesaOnline(&mpesa.STKPushRequest{
// 		Shortcode: shortCode,
// 		PartyB:    shortCode,
// 		Passkey:   mpesapasskey,

// 		Amount:                 payload.Amount,
// 		PhoneNumber:            payload.PhoneNumber,
// 		ReferenceCode:          description, //make sure its 13 characters by triming
// 		TransactionDescription: description,
// 		CallbackURL:            fmt.Sprintf("%v/anzaapi/payments/tenant/mpesa/payment_callback", getBaseCallbackUrl()), // Add your callback URL here
// 		TransactionType:        transactionType,                                                                        // CustomerPayBillOnline or CustomerBuyGoodsOnline
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println(shortCode)
// 		helperfunctions.LogToAFileInServer("Failed To Make StkPush", "ERROR")
// 		conn.WriteMessage(1, []byte("failed to  make stk push error"))
// 		conn.WriteMessage(1, []byte("payment failed"))

// 		if appconstants.IsDebug {
// 			fmt.Println(err)
// 		}

// 		return
// 	}

// 	// Check if the request was successful
// 	if response.IsSuccessful {
// 		// Handle your successful logic here
// 		conn.WriteMessage(1, []byte("Stk push Sucess"))
// 		conn.WriteMessage(1, []byte("Check you phone to Complete Payment"))

// 		tenantSocketConnections[response.CheckoutRequestID] = conn
// 		tenantPaymentsRequestsConnections[response.CheckoutRequestID] = payload
// 	} else {
// 		conn.WriteMessage(1, []byte("failed to  make stk push-- safaricom"))
// 		conn.WriteMessage(1, []byte(response.ErrorMessage))

// 		conn.WriteMessage(1, []byte("payment failed"))

// 	}

// 	for {
// 		_, message, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("read:", err)
// 			break
// 		}
// 		log.Printf("recv: %s", message)

// 		err = conn.WriteMessage(1, message)
// 		if string(message) == "close" {
// 			conn.Close()
// 			delete(tenantSocketConnections, response.CheckoutRequestID)
// 			delete(tenantPaymentsRequestsConnections, response.CheckoutRequestID)

// 		}
// 		if err != nil {
// 			log.Println("write:", err)
// 			break
// 		}
// 	}
// }

// func TenantPaymentrequestSafCallbackHandler(c *gin.Context) {

// 	var res mpesa.LipaNaMpesaOnlineCallback

// 	err := c.ShouldBind(&res)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var conn = tenantSocketConnections[res.Body.StkCallback.CheckoutRequestID]
// 	var originalRequestPayload = tenantPaymentsRequestsConnections[res.Body.StkCallback.CheckoutRequestID]

// 	if res.Body.StkCallback.ResultCode == 0 {

// 		conn.WriteMessage(1, []byte(res.Body.StkCallback.ResultDesc))
// 		conn.WriteMessage(1, []byte("updating .."))

// 		propid, _ := primitive.ObjectIDFromHex(originalRequestPayload.PropertyId)
// 		//get the document, updates the client ids with this, save the document
// 		rental, err := helperfunctions.GetPropertyModelWithFilterFromDb(map[string]interface{}{"_id": propid})

// 		if err != nil {
// 			helperfunctions.LogToAFileInServer("Error getting property from db during tenant saf callback paymnet", "ERROR")
// 			if appconstants.IsDebug {
// 				fmt.Println("error getting  property  doc--tenant saf callback")
// 			}
// 		}
// 		//get wallet details
// 		wallet, err := helperfunctions.GetWalletModelWithFilterFromDb(map[string]interface{}{"propertyId": originalRequestPayload.PropertyId})

// 		if err != nil {
// 			helperfunctions.LogToAFileInServer("Error getting property  wallet from db during tenant saf callback paymnet", "ERROR")
// 			if appconstants.IsDebug {
// 				fmt.Println("error getting  property wallet  doc--tenant saf callback")
// 			}
// 		}

// 		for i, house := range rental.Houses {
// 			if house.HouseNumber == originalRequestPayload.HouseNumber {
// 				if originalRequestPayload.BillName == "All Bills" {

// 					for j, _ := range house.HouseExpenses {

// 						rental.Houses[i].HouseExpenses[j].RemainingAmount = 0
// 						rental.Houses[i].HouseExpenses[j].IsPaid = true
// 						rental.Houses[i].HouseExpenses[j].UpdatedAt = helperfunctions.GetCurrrentDate()

// 					}

// 				} else {

// 					if originalRequestPayload.IsCleared {

// 						for j, expense := range house.HouseExpenses {
// 							// check expense name if matches the original bill name
// 							if expense.Name == originalRequestPayload.BillName {
// 								expense.IsPaid = true
// 								expense.RemainingAmount = 0
// 								expense.UpdatedAt = helperfunctions.GetCurrrentDate()

// 								rental.Houses[i].HouseExpenses[j] = expense
// 							}

// 						}
// 					} else {

// 						for j, expense := range house.HouseExpenses {

// 							if expense.Name == originalRequestPayload.BillName {
// 								expense.RemainingAmount = expense.RemainingAmount - float64(originalRequestPayload.Amount)

// 								expense.UpdatedAt = helperfunctions.GetCurrrentDate()
// 								rental.Houses[i].HouseExpenses[j] = expense
// 							}
// 						}

// 					}

// 				}

// 			}

// 		}

// 		//save transaction
// 		var paymentcategory = "Partial Payment"
// 		if originalRequestPayload.IsFullPayment {
// 			paymentcategory = "Full Payment"
// 		}
// 		var newTransaction transactionsmodel.TransactionsModel = transactionsmodel.TransactionsModel{Amount: int64(originalRequestPayload.Amount), PropertyId: originalRequestPayload.PropertyId,

// 			PropertyName: rental.PropertyName, HouseNumber: originalRequestPayload.HouseNumber, BillName: originalRequestPayload.BillName, PayerName: originalRequestPayload.PayerName, PaymentMethod: "Mpesa Online", CreatedAt: helperfunctions.GetCurrrentDate(), PaymentTime: originalRequestPayload.PayerTime,
// 			UpdatedAt: helperfunctions.GetCurrrentDate(), TotalRequiredAmount: int64(originalRequestPayload.RequiredAmount), RemainingAmount: int64(originalRequestPayload.RemainingAmount), MpesaConfirmationCode: res.Body.StkCallback.CheckoutRequestID, Category: paymentcategory,
// 		}

// 		newid, err := mongodbapi.UpdateADocInCollection(rental.ToMap(), mongodbapi.AllPropertiesCollection, rental.PropertyId)
// 		if err != nil {
// 			if appconstants.IsDebug {
// 				fmt.Println("unable to update the property")
// 			}
// 			if appconstants.IsDebug {
// 				fmt.Println(newid)
// 			}
// 			return
// 		}
// 		if appconstants.IsDebug {
// 			fmt.Println("property updated after a successful payment")
// 		}

// 		_, err = mongodbapi.AddADocumentToCollection(newTransaction.ToMap(), mongodbapi.AllTransactionsCollection)
// 		if err != nil {
// 			helperfunctions.LogToAFileInServer("unable to save transaction", "ERROR")
// 		}

// 		conn.WriteMessage(1, []byte("payment success"))
// 		helperfunctions.CreateNewNotificationAndSaveIt(rental.OwnerId, fmt.Sprintf("%v payed %v  for %v in  house number %v ", originalRequestPayload.PayerName, originalRequestPayload.Amount, originalRequestPayload.BillName, originalRequestPayload.HouseNumber), "Payments", originalRequestPayload.PayerTime)
// 		//todo send notification also to the user
// 		if rental.IsInternalMpesaWallet {
// 			wallet.AvailableAmount += int64(originalRequestPayload.Amount)
// 			wallet.TotalReceivedAmount += int64(originalRequestPayload.Amount)
// 			mongodbapi.UpdateADocInCollection(wallet.ToMap(), mongodbapi.MpesaWalletsCollection, wallet.WalletId)
// 		}

// 		onlineusersmap.SendRefreshMessageToUserWithId(rental.OwnerId)
// 		if len(rental.InvitedUsers) > 0 {
// 			for _, v := range rental.InvitedUsers {
// 				onlineusersmap.SendRefreshMessageToUserWithEmail(v.UserEmail)

// 			}

// 		}

// 		tenantuser, err := helperfunctions.GetUserModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(originalRequestPayload.TenantId)})
// 		if err == nil {
// 			if !tenantuser.IsRefererPaid {
// 				if len(originalRequestPayload.InvitedBy) > 1 {
// 					ctx := context.Background()

// 					//get the document, updates the client ids with this, save the document
// 					var bonusDoc referearnmodel.ReferalBonus
// 					if err = mongodbapi.AllReferBonusesDataCollection.FindOne(ctx, bson.M{"bonusCode": originalRequestPayload.InvitedBy}).Decode(&bonusDoc); err != nil {
// 						fmt.Println("error getting  bonus doc")
// 					} else {
// 						bonusDoc.PaidAmount += 30

// 						mongodbapi.UpdateADocInCollection(bonusDoc.ToMap(), mongodbapi.AllReferBonusesDataCollection, bonusDoc.Id)
// 						if appconstants.IsDebug {
// 							fmt.Println("referer paid amount added with one hundred")
// 						}
// 						tenantuser.IsRefererPaid = true

// 						//update tenat
// 						//add notificaton for refrerid  and refresh him with id
// 						//todo check if this can run on go routine
// 						mongodbapi.UpdateADocInCollection(tenantuser.ToMap(), mongodbapi.AllUsersCollection, tenantuser.UserId)
// 						helperfunctions.CreateNewNotificationAndSaveIt(bonusDoc.UserId, fmt.Sprintf("You Have Earned ksh  30  from your referal of %v", tenantuser.Email), "bonus", originalRequestPayload.PayerTime)
// 						helperfunctions.CreateNewNotificationAndSaveIt(tenantuser.UserId, fmt.Sprintf("Your referer received a bonus of ksh 30.. his/her id=%v", originalRequestPayload.InvitedBy), "bonus", originalRequestPayload.PayerTime)
// 						onlineusersmap.SendRefreshMessageToUserWithId(bonusDoc.UserId)
// 						onlineusersmap.SendRefreshMessageToUserWithId(tenantuser.UserId)

// 					}

// 				}

// 			}
// 		}

// 	} else {

// 		conn.WriteMessage(1, []byte(res.Body.StkCallback.ResultDesc))
// 		// time.Sleep(time.Millisecond * 600)
// 		conn.WriteMessage(1, []byte("payment failed"))
// 	}

// }

// var testPaymentsRequestsConnections = map[string]paymentmodels.TestMpesaRequestPayload{}

// func TestMpesPaymentCreditialsHandler(c *gin.Context) {

// 	phoneNumber, _ := strconv.ParseUint(c.Query("phoneNumber"), 10, 64)
// 	shortcode, _ := strconv.ParseUint(c.Query("shortcode"), 10, 64)
// 	consumerkey := c.Query("consumerKey")
// 	consumerSecret := c.Query("consumerSecret")
// 	passkey := c.Query("passKey")
// 	userId := c.Query("userId")
// 	time := c.Query("time")
// 	ispaybill, _ := strconv.ParseBool(c.Query("isPaybill"))
// 	description := c.Query("description")

// 	fmt.Println(consumerSecret)
// 	fmt.Println(consumerkey)
// 	fmt.Println(passkey)
// 	fmt.Println(ispaybill)
// 	fmt.Println(shortcode)
// 	fmt.Println(description)

// 	var payload paymentmodels.TestMpesaRequestPayload = paymentmodels.TestMpesaRequestPayload{PhoneNumber: phoneNumber, ConsumerKey: consumerkey, ConsumerSecret: consumerSecret, UserId: userId, Time: time, PassKey: passkey, ShortCode: uint(shortcode), IsPaybill: ispaybill}

// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Print("upgrade:", err)
// 		return
// 	}

// 	// conf, err := config.Get()
// 	if err != nil {
// 		log.Fatalln(err)
// 		return
// 	}

// 	var transactiontype string

// 	if ispaybill {
// 		transactiontype = "CustomerPayBillOnline"
// 	} else {
// 		transactiontype = "BuyGoodsOnline"

// 	}
// 	conn.WriteMessage(1, []byte("Connection Established"))
// 	mpesaApp := mpesa.Init(&config.Credentials{ConsumerKey: os.Getenv("MPESA_C2B_CONSUMERKEY"), ConsumerSecret: os.Getenv("MPESA_C2B_CONSUMERSECRET")}, true)
// 	mpesaApp.B2CPayment(&mpesa.B2CPaymentRequest{})
// 	response, err := mpesaApp.LipaNaMpesaOnline(&mpesa.STKPushRequest{
// 		Shortcode: payload.ShortCode,
// 		PartyB:    payload.ShortCode,
// 		Passkey:   os.Getenv("MPESA_C2B_PASSKEY"),
// 		Amount:    1,

// 		PhoneNumber:            payload.PhoneNumber,
// 		ReferenceCode:          description,
// 		TransactionDescription: description,
// 		CallbackURL:            fmt.Sprintf("%v/anzaapi/payments/owner/mpesa/test_creditials_callback", getBaseCallbackUrl()), // Add your callback URL here
// 		TransactionType:        transactiontype,                                                                               // CustomerPayBillOnline or CustomerBuyGoodsOnline
// 	})
// 	if err != nil {

// 		helperfunctions.LogToAFileInServer("Failed To Make StkPush", "ERROR")
// 		conn.WriteMessage(1, []byte("failed to  make stk push error"))
// 		conn.WriteMessage(1, []byte("payment failed"))

// 		if appconstants.IsDebug {
// 			fmt.Println(err)
// 		}

// 		return
// 	}

// 	// Check if the request was successful
// 	if response.IsSuccessful {
// 		// Handle your successful logic here
// 		conn.WriteMessage(1, []byte("Stk push Sucess"))
// 		conn.WriteMessage(1, []byte("Check you phone to Complete Payment"))

// 		socketConnections[response.CheckoutRequestID] = conn
// 		testPaymentsRequestsConnections[response.CheckoutRequestID] = payload
// 	} else {

// 		conn.WriteMessage(1, []byte("failed to  make stk push-- safaricom"))
// 		conn.WriteMessage(1, []byte(response.ErrorMessage))

// 		conn.WriteMessage(1, []byte("payment failed"))

// 	}

// 	for {
// 		_, message, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("read:", err)
// 			break
// 		}
// 		log.Printf("recv: %s", message)

// 		err = conn.WriteMessage(1, message)
// 		if string(message) == "close" {
// 			conn.Close()
// 			delete(socketConnections, response.CheckoutRequestID)
// 			delete(paymentsRequestsConnections, response.CheckoutRequestID)

// 		}
// 		if err != nil {
// 			log.Println("write:", err)
// 			break
// 		}
// 	}
// }

// func TestPaymentrequestSafCallbackHandler(c *gin.Context) {

// 	var res mpesa.LipaNaMpesaOnlineCallback
// 	err := c.ShouldBind(&res)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var conn = socketConnections[res.Body.StkCallback.CheckoutRequestID]
// 	var originalRequestPayload = testPaymentsRequestsConnections[res.Body.StkCallback.CheckoutRequestID]

// 	if res.Body.StkCallback.ResultCode == 0 {

// 		conn.WriteMessage(1, []byte(res.Body.StkCallback.ResultDesc))
// 		conn.WriteMessage(1, []byte("updating .."))

// 		helperfunctions.CreateNewNotificationAndSaveIt(originalRequestPayload.UserId, "ksh 1  was paid to test lipa na mpesa online", "Testing payments", originalRequestPayload.Time)

// 		userDoc, _ := helperfunctions.GetUserModelWithFilterFromDb(map[string]interface{}{"_id": helperfunctions.GetMongoidFromString(originalRequestPayload.UserId)})

// 		newid, err := mongodbapi.UpdateADocInCollection(userDoc.ToMap(), mongodbapi.AllUsersCollection, userDoc.UserId)
// 		if err != nil {
// 			if appconstants.IsDebug {
// 				fmt.Println("unable to update the user")
// 			}
// 			if appconstants.IsDebug {
// 				fmt.Println(newid)
// 			}
// 			return
// 		}
// 		if appconstants.IsDebug {
// 			fmt.Println("user updated after a successful payment")
// 		}
// 		conn.WriteMessage(1, []byte("payment success"))
// 	} else {
// 		conn.WriteMessage(1, []byte(res.Body.StkCallback.ResultDesc))
// 		// time.Sleep(time.Millisecond * 600)
// 		conn.WriteMessage(1, []byte("payment failed"))
// 	}

// }
