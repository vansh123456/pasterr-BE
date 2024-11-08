package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/vansh123456/pasterr/api"
)

const (
	dbDriver = "postgres"
	dbSource = "DBurl"
	//serverAddress = "0.0.0.0:8080"
)

func main() {
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("failed to start the server", err)
	}
	defer dbConn.Close()

	router := gin.Default()
	api.InitializeRouter(router, dbConn)
	log.Fatal(router.Run(":8080"))
}
