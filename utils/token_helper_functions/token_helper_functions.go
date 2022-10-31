package tokenhelperfunctions

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userid string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 60 * 6).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("TOKEN_ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
func CreateRefreshToken(userid string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

//uncommmnet when you install goredis redis v7
// func IntializeRedis() {
// 	//Initializing redis
// 	dsn := os.Getenv("REDIS_DSN")
// 	if len(dsn) == 0 {
// 		dsn = "localhost:6379"
// 	}
// 	client = redis.NewClient(&redis.Options{
// 		Addr: dsn, //redis port

// 	})
// 	_, err := client.Ping().Result()
// 	if err != nil {
// 		LogToAFileInServer("Redis is Not Connected", "ERROR")
// 	}
// }

//this required uuid and redis

// func CreateToken(userid uint64) (*TokenDetails, error) {
// 	td := &TokenDetails{}
// 	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
// 	td.AccessUuid = uuid.NewV4().String()

// 	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
// 	td.RefreshUuid = uuid.NewV4().String()

// 	var err error
// 	//Creating Access Token
// 	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
// 	atClaims := jwt.MapClaims{}
// 	atClaims["authorized"] = true
// 	atClaims["access_uuid"] = td.AccessUuid
// 	atClaims["user_id"] = userid
// 	atClaims["exp"] = td.AtExpires
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	//Creating Refresh Token
// 	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
// 	rtClaims := jwt.MapClaims{}
// 	rtClaims["refresh_uuid"] = td.RefreshUuid
// 	rtClaims["user_id"] = userid
// 	rtClaims["exp"] = td.RtExpires
// 	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
// 	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return td, nil
// }

//protected route functions

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
func VerifyToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func VerifyRefreshToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) (string, error) {
	t := ExtractToken(r)
	token, err := VerifyToken(t)

	if err != nil {
		return "", err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return "", err
	}

	return token.Claims.(jwt.MapClaims)["user_id"].(string), nil
}
func RefreshTokenValid(r *http.Request) (string, error) {
	t := ExtractToken(r)
	token, err := VerifyRefreshToken(t)

	if err != nil {
		return "", err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return "", err
	}

	return token.Claims.(jwt.MapClaims)["user_id"].(string), nil
}

// func DeleteTokenFromRedis(givenUuid string) (int64, error) {
// 	deleted, err := client.Del(givenUuid).Result()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return deleted, nil
// }

// func Logout(c *gin.Context) {
// 	au, err := ExtractTokenMetadata(c.Request)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "unauthorized")
// 		return
// 	}
// 	deleted, delErr := DeleteAuth(au.AccessUuid)
// 	if delErr != nil || deleted == 0 { //if any goes wrong
// 		c.JSON(http.StatusUnauthorized, "unauthorized")
// 		return
// 	}
// 	c.JSON(http.StatusOK, "Successfully logged out")
// }

// func RefreshTokenHandler(c *gin.Context) {
// 	mapToken := map[string]string{}
// 	if err := c.ShouldBindJSON(&mapToken); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, err.Error())
// 		return
// 	}
// 	refreshToken := mapToken["refresh_token"]

// 	//verify the token
// 	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
// 	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
// 		//Make sure that the token method conform to "SigningMethodHMAC"
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("REFRESH_SECRET")), nil
// 	})
// 	//if there is an error, the token must have expired
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "Refresh token expired")
// 		return
// 	}
// 	//is token valid?
// 	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
// 		c.JSON(http.StatusUnauthorized, err)
// 		return
// 	}
// 	//Since token is valid, get the uuid:
// 	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
// 	if ok && token.Valid {
// 		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
// 		if !ok {
// 			c.JSON(http.StatusUnprocessableEntity, err)
// 			return
// 		}
// 		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
// 		if err != nil {
// 			c.JSON(http.StatusUnprocessableEntity, "Error occurred")
// 			return
// 		}
// 		//Delete the previous Refresh Token
// 		deleted, delErr := DeleteAuth(refreshUuid)
// 		if delErr != nil || deleted == 0 { //if any goes wrong
// 			c.JSON(http.StatusUnauthorized, "unauthorized")
// 			return
// 		}
// 		//Create new pairs of refresh and access tokens
// 		ts, createErr := CreateToken(userId)
// 		if createErr != nil {
// 			c.JSON(http.StatusForbidden, createErr.Error())
// 			return
// 		}
// 		//save the tokens metadata to redis
// 		saveErr := CreateAuth(userId, ts)
// 		if saveErr != nil {
// 			c.JSON(http.StatusForbidden, saveErr.Error())
// 			return
// 		}
// 		tokens := map[string]string{
// 			"access_token":  ts.AccessToken,
// 			"refresh_token": ts.RefreshToken,
// 		}
// 		c.JSON(http.StatusCreated, tokens)
// 	} else {
// 		c.JSON(http.StatusUnauthorized, "refresh expired")
// 	}
// }
