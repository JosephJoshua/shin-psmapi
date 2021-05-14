package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/JosephJoshua/shin-psmapi/conf"
	"github.com/JosephJoshua/shin-psmapi/db"
	"github.com/JosephJoshua/shin-psmapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
		return
	}

	dateStr := time.Now().Format("2006-01-02")
	logFile, err := os.Create(fmt.Sprintf("logs/%v.log", dateStr))

	if err != nil {
		log.Fatal("Unable to create log file:\n" + err.Error())
		return
	}

	defer logFile.Close()
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	r := gin.Default()

	db.Init()

	authMiddleware, err := conf.InitJWTMiddleware()
	if err != nil {
		log.Fatal("JWT initialization error: " + err.Error())
	}

	conf.MigrateDB(db.GetDB())
	conf.SetupRoutes(r, authMiddleware)

	setupValidators()

	if err := r.Run(); err != nil {
		log.Panic(err.Error())
	}
}

func setupValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("user_role", func(fl validator.FieldLevel) bool {
			return utils.IsValidUserRole(fl.Field().String())
		})

		v.RegisterValidation("servisan_search_by_col", func(fl validator.FieldLevel) bool {
			return utils.IsValidServisanSearchByColumn(fl.Field().String())
		})

		v.RegisterValidation("status_servisan", func(fl validator.FieldLevel) bool {
			return utils.IsValidServisanStatus(fl.Field().String())
		})
	}
}
