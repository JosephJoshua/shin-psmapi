package models

import (
	"fmt"

	"github.com/JosephJoshua/shin-psmapi/db"
)

type DPType struct {
	ID   int 		`json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name string `json:"name" gorm:"size:256;not null"`
}

func (DPType) TableName() string {
	return "dp_types"
}

type DPTypeModel struct{}

func (DPTypeModel) All() ([]DPType, error) {
	var dpTypeList []DPType

	err := db.GetDB().Find(&dpTypeList).Error
	if err != nil {
		return nil, fmt.Errorf("All: %w", err)
	}
	
	return dpTypeList, nil
}

func (DPTypeModel) ByID(id int) (*DPType, error) {
	var dpType *DPType

	err := db.GetDB().First(&dpType, id).Error
	if err != nil {
		return nil, fmt.Errorf("ByID: %w", err)
	}
	
	return dpType, nil
}
