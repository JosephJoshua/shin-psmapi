package forms

import "time"

type GetAllSparepartForm struct {
	MinDate time.Time `form:"min_date"`
	MaxDate time.Time `form:"max_date"`
}

type CreateSparepartForm struct {
	NomorNota int     `json:"nomor_nota" binding:"required"`
	Nama      string  `json:"nama" binding:"required,max=256"`
	Harga     float64 `json:"harga" binding:"required"`
}
