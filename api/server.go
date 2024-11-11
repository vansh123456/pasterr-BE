package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/vansh123456/pasterr/services"
	"github.com/vansh123456/pasterr/middleware"
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
	snippets := router.Group("/snippets",middleware.AuthMiddleware()){
		router.POST("/createsnippets", func(c *gin.Context) {
			services.CreateSnippetHandler(c, dbConn)
		})
		snippets.GET("", func(c *gin.Context) {
            services.ListSnippetsHandler(c, dbConn)
        })
		outer.GET("/:id", func(c *gin.Context) {
			services.GetSnippetByIDHandler(c, dbConn)
		})
		router.PUT("/:id", func(c *gin.Context) {
			services.UpdateSnippetContent(c, dbConn)
		})
		router.DELETE("/:id", func(c *gin.Context) {
			services.DeleteSnippetHandler(c, dbConn)
		})

	}
	
}
