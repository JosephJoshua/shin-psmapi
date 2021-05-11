// Migration calls the AutoMigrate function from gorm for all models in the API.
// We can't put this in the db package so as not to introduce an import cycle.

package migration

import (
	"shin-psmapi/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}