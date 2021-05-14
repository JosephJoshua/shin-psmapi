package models

import (
	"fmt"
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

func (SalesModel) ByID(id int) (*Sales, error) {
	var sales *Sales

	res := db.GetDB().Where("id = ?", id).Find(&sales)

	if res.Error != nil {
		return sales, res.Error
	}

	if res.RowsAffected < 1 {
		return &Sales{}, fmt.Errorf("tidak ada sales dengan ID %d", id)
	}

	return sales, nil
}

func (SalesModel) Create(form forms.CreateSalesForm) (id int, err error) {
	sales := Sales{Nama: form.Nama}
	err = db.GetDB().Create(&sales).Error

	return sales.ID, err
}

func (SalesModel) Delete(id int) error {
	res := db.GetDB().Delete(&Sales{}, id)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 {
		return fmt.Errorf("tidak ada sales dengan ID %d", id)
	}

	return nil
}
