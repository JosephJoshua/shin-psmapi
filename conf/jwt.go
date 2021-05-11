package conf

import (
	"shin-psmapi/models"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const IdentityKey = "id"

var JWTMiddleware *jwt.GinJWTMiddleware

func InitJWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	JWTMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "shin-psmapi",
		Key:           []byte("very very very secret key"),
		Timeout:       time.Hour * 24,
		MaxRefresh:    time.Hour,
		IdentityKey:   IdentityKey,
		Authenticator: models.AuthenticateUser,
		Authorizator:  authorizator,
		Unauthorized:  unauthorized,
		PayloadFunc:   payloadFunc,
		LoginResponse: loginResponse,
		TimeFunc:      time.Now,
		SendCookie:    true,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
	})

	return JWTMiddleware, err
}

func authorizator(data interface{}, c *gin.Context) bool {
	return true
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if user, ok := data.(models.User); ok {
		return jwt.MapClaims{IdentityKey: user.ID}
	}

	return jwt.MapClaims{}
}

func loginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(code, gin.H{
		"expire": expire,
		"token":  token,
	})
}
