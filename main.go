package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3030"
	}

	r.Run(":" + port)
}
