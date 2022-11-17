package middleware

import (
	helpers "mongo-api/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusUnauthorized})
			c.Abort()
			return
		}
		arry := strings.Split(token, " ")
		token = arry[1]
		claims, err := helpers.Validatetoken(token)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err, "tkn": token})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("user_type", claims.User_type)
		c.Next()
	}

}
