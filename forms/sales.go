package forms

type CreateSalesForm struct {
	Nama string `json:"nama" binding:"required,max=256"`
}
