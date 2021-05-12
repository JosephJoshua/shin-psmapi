package forms

type GetAllSalesForm struct {
	SearchQuery string `form:"q"`
}

type CreateSalesForm struct {
	Nama string `json:"nama" binding:"required,max=256"`
}
