package main

import (
	"log"
	"os"
	"shin-psmapi/db"
	"shin-psmapi/migration"
	"shin-psmapi/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()

	db.Init()
	migration.Migrate(db.GetDB())

	r.Static("/public", "./public")
	setupValidators()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3030"
	}

	if err := r.Run(":" + port); err != nil {
		log.Panic(err.Error())
	}
}

func setupValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("user_role", func(fl validator.FieldLevel) bool {
			return utils.IsValidUserRole(fl.Field().String())
		})
	}
}
