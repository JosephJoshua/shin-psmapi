// Migration calls the AutoMigrate function from gorm for all models in the API.
// We can't put this in the db package so as not to introduce an import cycle.

package conf

import (
	"github.com/JosephJoshua/shin-psmapi/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.Exec(`
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
			CREATE TYPE user_role AS ENUM ('admin', 'customer_service', 'buyer');
		END IF;
	END$$;
	`)
	
	db.Exec(`
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_servisan') THEN
			CREATE TYPE status_servisan AS ENUM ('Sedang dikerjakan', 'Jadi (Belum diambil)', 'Jadi (Sudah diambil)', 'Tidak jadi (Belum diambil)', 'Tidak jadi (Sudah diambil)', 'Pending');
		END IF;
	END$$;
	`)

	db.AutoMigrate(&models.DPType{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Sales{})
	db.AutoMigrate(&models.Teknisi{})
	db.AutoMigrate(&models.Servisan{})
	db.AutoMigrate(&models.Sparepart{})
}
