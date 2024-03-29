package models

import (
	"errors"

	"github.com/JosephJoshua/shin-psmapi/db"
	"github.com/JosephJoshua/shin-psmapi/forms"
	"github.com/JosephJoshua/shin-psmapi/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
  ID       int            `json:"id" gorm:"<-:false;primaryKey;autoIncrement;not null"`
	Email    string         `json:"email" gorm:"size:256;not null"`
	Password string         `json:"password" gorm:"size:128;not null"`
	Username string         `json:"username" gorm:"size:256;not null"`
	Role     utils.UserRole `json:"role" gorm:"type:user_role;not null"`
}

func (User) TableName() string {
	return "users"
}

type UserModel struct{}

func AuthenticateUser(form forms.LoginForm) (interface{}, error) {
	var user User
	res := db.GetDB().Where("email=LOWER(?)", form.Email).Find(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected < 1 {
		return nil, errors.New("email atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		return nil, errors.New("email atau password salah")
	}

	return user, nil
}

func (UserModel) All() ([]User, error) {
	var userList []User
	err := db.GetDB().Find(&userList).Error

	return userList, err
}

func (UserModel) One(id int) (User, error) {
	var user User
	err := db.GetDB().Where("id = ?", id).Select("id", "email", "username", "role").Find(&user).Error

	return user, err
}

func (UserModel) Register(form forms.RegisterForm) (User, error) {
	db := db.GetDB()

	var userCount int64
	if err := db.Model(&User{}).Where("email=LOWER(?)", form.Email).Count(&userCount).Error; err != nil {
		return User{}, err
	}

	if userCount > 0 {
		return User{}, errors.New("sudah ada user dengan email yang sama")
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, errors.New("gagal meng-hash password: " + err.Error())
	}

	user := User{Email: form.Email, Password: string(hashedPwd), Username: form.Username, Role: form.Role}
	if err = db.Create(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}
