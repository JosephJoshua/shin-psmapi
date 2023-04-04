package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
  "path/filepath"

	"github.com/JosephJoshua/shin-psmapi/conf"
	"github.com/JosephJoshua/shin-psmapi/db"
	"github.com/JosephJoshua/shin-psmapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	loadDotEnvFile()

	logFile := createLogFile()
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

func loadDotEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
		return
	}
}

func createLogFile() *os.File {
  const logDir string = "logs/"

  err := os.MkdirAll(logDir, os.ModePerm)
  if err != nil {
    log.Fatal("Unable to create logs/ folder:\n" + err.Error())
    return nil
  }

	dateStr := time.Now().Format("2006-01-02")
  filePath := filepath.Join(logDir, fmt.Sprintf("%v.log", dateStr))

	// Open a file with append mode if it exists; if it doesn't, create it.
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Unable to create log file:\n" + err.Error())
		return nil
	}

	if _, err := logFile.WriteString("\n"); err != nil {
		log.Fatal("Unable to write to log file:\n" + err.Error())
		return nil
	}

	return logFile
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
