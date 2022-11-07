package middlewares

import (
	"net/http"
	"time"

	helperfunctions "github.com/Anza2022/Anza_Backend_API/utils/helper_functions"
	tokenhelperfunctions "github.com/Anza2022/Anza_Backend_API/utils/token_helper_functions"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
)

var redisStore = persist.NewRedisStore(redis.NewClient(&redis.Options{
	Network: "tcp",
	Addr:    "127.0.0.1:6379",
}))

func CacheRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cache.CacheByRequestURI(redisStore, 200*time.Second)
		c.Next()
	}
}
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, err := tokenhelperfunctions.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, bson.M{"message": "Unauthorised Request"})
			c.Abort()
			return
		}

		c.Set("userId", userid)
		c.Next()
	}
}
func RefreshTokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, err := tokenhelperfunctions.RefreshTokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, bson.M{"message": "Unauthorised Request"})
			c.Abort()
			return
		}

		c.Set("userId", userid)
		c.Next()
	}
}

func CheckIfAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid := c.GetString("userId")

		usermodel, _ := helperfunctions.GetUserModelWithFilterFromDb(bson.M{"_id": helperfunctions.GetMongoidFromString(userid)})

		if usermodel.IsAdmin {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, bson.M{"message": "Permision Denied"})
			c.Abort()
			return

		}

	}
}
