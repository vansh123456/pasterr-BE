package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vansh123456/pasterr/services"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("auth middleware started")
		tokenString := c.GetHeader("Authorization")
		fmt.Println("Authorization header:", tokenString)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}
		// Remove the "Bearer " prefix, if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:] // Strip "Bearer " from the token string
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}
		claims, err := services.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		fmt.Println("parsed userid", claims.UserID)
		c.Set("user_id", claims.UserID) //sets the userID
		c.Next()                        //call the next handler so that the next function can easily execute!
	}
}
