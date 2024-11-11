package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/vansh123456/pasterr/services"
)

// what we were doing in store.db,usko abstraction layer banake we used to connect with that,but we havent used that for simple CRUD functions so we are defining it here
func InitializeRouter(router *gin.Engine, dbConn *sql.DB) {

	router.POST("/signup", func(c *gin.Context) {
		services.SignupHandler(c, dbConn)
	})
	router.POST("/signin", func(c *gin.Context) {
		services.SigninHandler(c, dbConn)
	})
	router.GET("/users", func(c *gin.Context) {
		services.ListUsersHandler(c, dbConn)
	})
}
