package forms

type CreateTeknisiForm struct {
	Nama string `json:"nama" binding:"required,max=256"`
}
