package forms

type GetAllTeknisiForm struct {
	SearchQuery string `form:"q"`
}

type CreateTeknisiForm struct {
	Nama string `json:"nama" binding:"required,max=256"`
}
