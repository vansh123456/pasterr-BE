package services

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/vansh123456/pasterr/db/sqlc"
)

func CreateSnippetHandler(c *gin.Context, dbConn *sql.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var params db.CreateSnippetParams
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	params.UserID = int32(userID.(uint)) //UserID in the sqlc params is in int32 so we are converting it to int32 explicitely
	queries := db.New(dbConn)
	snippet, err := queries.CreateSnippet(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:failed to create snippet": err.Error()})
		return
	}
	c.JSON(http.StatusOK, snippet)
}

func ListSnippetsHandler(c *gin.Context, dbConn *sql.DB) {
	queries := db.New(dbConn)
	snippets, err := queries.ListSnippets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve snippets"})
		return
	}

	c.JSON(http.StatusOK, snippets)
}

func GetSnippetByIDHandler(c *gin.Context, dbConn *sql.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error invalid ID": err.Error()})
		return
	}
	queries := db.New(dbConn)
	snippet, err := queries.GetSnippetByID(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:failed to list snippets": err.Error()})
		return
	}
	c.JSON(http.StatusOK, snippet)
}
func UpdateSnippetContent(c *gin.Context, dbConn *sql.DB) {
	var params db.UpdateSnippetContentParams
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	queries := db.New(dbConn)
	updatedsnippet, err := queries.UpdateSnippetContent(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update snippet"})
		return
	}
	c.JSON(http.StatusOK, updatedsnippet)
}
func DeleteSnippetHandler(c *gin.Context, dbConn *sql.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error invalid ID to delete": err.Error()})
		return
	}
	queries := db.New(dbConn)
	if err := queries.DeleteSnippet(c.Request.Context(), int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete snippet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Snippet deleted"})
}
