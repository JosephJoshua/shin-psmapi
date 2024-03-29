package conf

import (
	"os"
	"time"

	"github.com/JosephJoshua/shin-psmapi/forms"
	"github.com/JosephJoshua/shin-psmapi/models"
	"github.com/JosephJoshua/shin-psmapi/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var JWTMiddleware *jwt.GinJWTMiddleware

func InitJWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	JWTMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         utils.JWTRealm,
		Key:           []byte(os.Getenv("JWT_KEY")),
		Timeout:       time.Hour * 24,
		MaxRefresh:    time.Hour,
		IdentityKey:   utils.JWTIdentityKey,
		Authenticator: authenticator,
		Authorizator:  authorizator,
		Unauthorized:  unauthorized,
		PayloadFunc:   payloadFunc,
		LoginResponse: loginResponse,
		TimeFunc:      time.Now,
		SendCookie:    true,
		TokenLookup:   utils.JWTTokenLookup,
		TokenHeadName: utils.JWTTokenHeadName,
	})

	return JWTMiddleware, err
}

func authenticator(c *gin.Context) (interface{}, error) {
	var form forms.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		return nil, err
	}

	return models.AuthenticateUser(form)
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
		return jwt.MapClaims{utils.JWTIdentityKey: user.ID}
	}

	return jwt.MapClaims{}
}

func loginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(code, gin.H{
		"expire": expire,
		"token":  token,
	})
}
