package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func AuthMiddelware() gin.HandlerFunc {
	return func(request *gin.Context) {

		var authorization_header string = request.GetHeader("Authorization")
		if authorization_header == "" {
			request.JSON(http.StatusUnauthorized, gin.H{"Message": "Invalid Token"})
			request.Abort()
			return
		}

		tokenString := strings.Split(authorization_header, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil {
			request.JSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
			request.Abort()
			return

		}

		if !token.Valid {
			request.JSON(http.StatusUnauthorized, gin.H{"Message": "Invalid Token"})
			request.Abort()
			return
		}
		request.Next()
	}
}
