package main

import (
	"log"
	"os"
	"shin-psmapi/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.Init()
	db.Migrate()

	r.Static("/public", "./public")

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3030"
	}

	if err := r.Run(":" + port); err != nil {
		log.Panic(err.Error())
	}
}
