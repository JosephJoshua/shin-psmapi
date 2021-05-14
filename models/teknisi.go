package models

import (
	"fmt"
	"shin-psmapi/db"
	"shin-psmapi/forms"
)

type Teknisi struct {
	ID   int    `json:"id"`
	Nama string `json:"nama" gorm:"size:256,not null"`
}

func (Teknisi) TableName() string {
	return "teknisi"
}

type TeknisiModel struct{}

func (TeknisiModel) All(searchQuery string) ([]Teknisi, error) {
	var teknisiList []Teknisi

	err := db.GetDB().Where("LOWER(nama) LIKE '%' || LOWER(?) || '%'", searchQuery).Find(&teknisiList).Error
	if err != nil {
		return teknisiList, err
	}

	return teknisiList, nil
}

func (TeknisiModel) ByID(id int) (*Teknisi, error) {
	var teknisi *Teknisi

	res := db.GetDB().Where("id = ?", id).Find(&teknisi)

	if res.Error != nil {
		return teknisi, res.Error
	}

	if res.RowsAffected < 1 {
		return &Teknisi{}, fmt.Errorf("tidak ada teknisi dengan ID %d", id)
	}

	return teknisi, nil
}

func (TeknisiModel) Create(form forms.CreateTeknisiForm) (id int, err error) {
	teknisi := Teknisi{Nama: form.Nama}
	err = db.GetDB().Create(&teknisi).Error

	return teknisi.ID, err
}

func (TeknisiModel) Delete(id int) error {
	res := db.GetDB().Delete(&Teknisi{}, id)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 {
		return fmt.Errorf("tidak ada teknisi dengan ID %d", id)
	}

	return nil
}
