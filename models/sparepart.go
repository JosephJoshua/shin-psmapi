package models

import (
	"fmt"
	"time"

	"github.com/JosephJoshua/shin-psmapi/db"
	"github.com/JosephJoshua/shin-psmapi/forms"
	"github.com/JosephJoshua/shin-psmapi/utils"
)

type Sparepart struct {
	ID               int       `json:"id" gorm:"primaryKey;type:integer GENERATED ALWAYS AS IDENTITY (INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1);not null"`
	NomorNota        int       `json:"nomor_nota" gorm:"type:integer;not null"`
	Servisan         Servisan  `json:"-" gorm:"foreignKey:NomorNota;constraint:OnDelete:RESTRICT;"`
	Nama             string    `json:"nama" gorm:"size:256;not null"`
	Harga            float64   `json:"harga" gorm:"type:double precision;not null;default:0"`
	TanggalPembelian time.Time `json:"tanggal_pembelian" gorm:"->;autoCreateTime;not null;default:CURRENT_TIMESTAMP"`
}

func (Sparepart) TableName() string {
	return "sparepart"
}

type SparepartModel struct{}

func (SparepartModel) All(form forms.GetAllSparepartForm) ([]Sparepart, error) {
	var (
		sparepartList []Sparepart
		query         string
		params        []interface{}
	)

	if !form.MinDate.IsZero() {
		query += "DATE(tanggal_pembelian) >= ?"

		params = append(params, form.MinDate.Format(utils.DateOnly))
	}

	if !form.MaxDate.IsZero() {
		if len(params) > 0 {
			query += " AND "
		}

		query += "DATE(tanggal_pembelian) <= ?"

		params = append(params, form.MaxDate.Format(utils.DateOnly))
	}

	if err := db.GetDB().Where(query, params...).Find(&sparepartList).Error; err != nil {
		return sparepartList, err
	}

	return sparepartList, nil
}

func (SparepartModel) ByNomorNota(nomorNota int) ([]Sparepart, error) {
	var sparepartList []Sparepart
	res := db.GetDB().Where("nomor_nota = ?", nomorNota).Find(&sparepartList)

	return sparepartList, res.Error
}

func (SparepartModel) Create(form forms.CreateSparepartForm) (id int, err error) {
	servisan := Servisan{NomorNota: form.NomorNota}

	err = db.GetDB().First(&servisan).Error
	if err != nil {
		return 0, fmt.Errorf("gagal mengambil servisan: %v", err)
	}

	sparepart := Sparepart{NomorNota: form.NomorNota, Harga: form.Harga, Nama: form.Nama}

	err = db.GetDB().Create(&sparepart).Error
	if err != nil {
		return 0, fmt.Errorf("gagal membuat sparepart: %v", err)
	}

	servisan.HargaSparepart += form.Harga

	err = db.GetDB().Save(&servisan).Error
	if err != nil {
		return 0, fmt.Errorf("gagal mengubah servisan: %v", err)
	}

	return sparepart.ID, nil
}

func (SparepartModel) Delete(id int) error {
	sparepart := Sparepart{ID: id}

	err := db.GetDB().First(&sparepart, id).Error
	if err != nil {
		return fmt.Errorf("gagal mengambil sparepart: %v", err)
	}	

	servisan := Servisan{NomorNota: sparepart.NomorNota}

	err = db.GetDB().First(&servisan).Error
	if err != nil {
		return fmt.Errorf("gagal mengambil servisan: %v", err)
	}
	
	servisan.HargaSparepart -= sparepart.Harga

	err = db.GetDB().Save(&servisan).Error
	if err != nil {
		return fmt.Errorf("gagal merubah servisan: %v", err)
	}

	err = db.GetDB().Delete(&sparepart).Error
	if err != nil {
		return fmt.Errorf("gagal menghapus sparepart: %v", err)
	}

	return nil
}
