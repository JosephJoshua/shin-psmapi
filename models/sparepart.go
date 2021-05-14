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
		query += "tanggal_pembelian >= ?"

		// MUST convert to ISO8601/RFC3339 format first before sending it to the postgres db
		params = append(params, utils.ToRFC3339TimeString(form.MinDate))
	}

	if !form.MaxDate.IsZero() {
		if len(params) > 0 {
			query += " AND "
		}

		query += "tanggal_pembelian <= ?"

		// MUST convert to ISO8601/RFC3339 format first before sending it to the postgres db
		params = append(params, utils.ToRFC3339TimeString(form.MaxDate))
	}

	if err := db.GetDB().Where(query, params...).Find(&sparepartList).Error; err != nil {
		return sparepartList, err
	}

	return sparepartList, nil
}

func (SparepartModel) ByNomorNota(nomorNota int) ([]Sparepart, error) {
	var sparepartList []Sparepart
	res := db.GetDB().Where("nomor_nota = ?", nomorNota).Find(&sparepartList)

	if res.Error != nil {
		return sparepartList, res.Error
	}

	if res.RowsAffected < 1 {
		return sparepartList, fmt.Errorf("tidak ada sparepart yang dibeli untuk servisan dengan nomor nota %d", nomorNota)
	}

	return sparepartList, nil
}

func (SparepartModel) Create(form forms.CreateSparepartForm) (id int, err error) {
	sparepart := Sparepart{NomorNota: form.NomorNota, Harga: form.Harga, Nama: form.Nama}
	err = db.GetDB().Create(&sparepart).Error

	return sparepart.ID, err
}

func (SparepartModel) Delete(id int) error {
	res := db.GetDB().Delete(&Sparepart{}, id)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 {
		return fmt.Errorf("tidak ada sparepart dengan ID %d", id)
	}

	return nil
}
