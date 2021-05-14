package forms

import "github.com/JosephJoshua/shin-psmapi/utils"

type LoginForm struct {
	Email    string `json:"email" binding:"required,email,max=256"`
	Password string `json:"password" binding:"required,max=64"`
}

type RegisterForm struct {
	Email string `json:"email" binding:"required,email,max=256"`

	// Password must not be greater than 64 characters in length since bcrypt
	// has a max input length of 72 bytes.
	Password string         `json:"password" binding:"required,max=64"`
	Username string         `json:"username" binding:"required,max=256"`
	Role     utils.UserRole `json:"role" binding:"required,user_role"`
}
