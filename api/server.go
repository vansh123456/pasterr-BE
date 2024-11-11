package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/vansh123456/pasterr/middleware"
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
	//url like localhost/snippets/createsnippets
	//localhost/snippets
	//localhost/snippets/:id
	snippets := router.Group("/snippets", middleware.AuthMiddleware())
	{
		snippets.POST("", func(c *gin.Context) { // POST /snippets
			services.CreateSnippetHandler(c, dbConn)
		})
		snippets.GET("", func(c *gin.Context) { // GET /snippets
			services.ListSnippetsHandler(c, dbConn)
		})
		snippets.GET("/:id", func(c *gin.Context) { // GET /snippets/:id
			services.GetSnippetByIDHandler(c, dbConn)
		})
		snippets.PUT("/:id", func(c *gin.Context) { // PUT /snippets/:id
			services.UpdateSnippetContent(c, dbConn)
		})
		snippets.DELETE("/:id", func(c *gin.Context) { // DELETE /snippets/:id
			services.DeleteSnippetHandler(c, dbConn)
		})
	}
}
