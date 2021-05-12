package models

import (
	"shin-psmapi/db"
	"shin-psmapi/forms"
)

type Sales struct {
	ID   int    `json:"id"`
	Nama string `json:"nama" gorm:"size:256,not null"`
}

type SalesModel struct{}

func (SalesModel) All(searchQuery string) ([]Sales, error) {
	var salesList []Sales

	err := db.GetDB().Where("LOWER(nama) LIKE '%' || LOWER(?) || '%'", searchQuery).Find(&salesList).Error
	if err != nil {
		return salesList, err
	}

	return salesList, nil
}

func (SalesModel) ByID(id int) (Sales, error) {
	var sales Sales

	err := db.GetDB().Where("id = ?", id).Find(&sales).Error
	if err != nil {
		return sales, err
	}

	return sales, nil
}

func (SalesModel) Create(form forms.CreateSalesForm) error {
	sales := Sales{Nama: form.Nama}
	err := db.GetDB().Create(&sales).Error

	return err
}

func (SalesModel) Delete(id int) error {
	err := db.GetDB().Delete(&Sales{}, id).Error
	return err
}
