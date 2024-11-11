package services

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vansh123456/pasterr/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	fmt.Println("hashedpass:", string(hashedPassword))
	return string(hashedPassword), nil
}

// check and compare the hashed Password
func CheckPassword(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil //default value is nil but if it is true it returns that
}

func SignupHandler(c *gin.Context, dbConn *sql.DB) {
	var params db.CreateAccountParams
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := HashPassword(params.Password) //takes the json password and passed it to the functions above!
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	params.Password = hashedPassword //updates the value going into the DB with the hashed Password
	queries := db.New(dbConn)
	user, err := queries.CreateAccount(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func ListUsersHandler(c *gin.Context, dbConn *sql.DB) {
	queries := db.New(dbConn)
	users, err := queries.ListUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
