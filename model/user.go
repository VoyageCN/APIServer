package model

import (
	"APISERVER/pkg/auth"
	"APISERVER/pkg/constvar"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserModel struct {
	BaseModel
	Username string         `json:"username" gorm:"column:username;not null;unique" binding:"required" validate:"min=1,max=32"`
	Password string         `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Printers []PrinterModel `json:"printers" gorm:"many2many:tb_user_printers;"`
}

func (c *UserModel) TableName() string {
	return "tb_users"
}

func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	DB.Self.Model(&u).Association("Printers").Find(&u.Printers)
	return u, d.Error
}

func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt the user password.
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
