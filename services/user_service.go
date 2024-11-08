package services

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vansh123456/pasterr/db/sqlc"
)

func SignupHandler(c *gin.Context, dbConn *sql.DB) {
	var params db.CreateAccountParams
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
