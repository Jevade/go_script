package model

import (
	"fmt"

	"../pkg/auth"
	"../pkg/constvar"
	validator "gopkg.in/go-playground/validator.v9"
)

//UserModel is to store userinfo
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

//TableName return table name
func (u *UserModel) TableName() string {
	return "tb_users"
}

// Create is to create user in db
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

//DeleteUser is to delete user from db
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.ID = id
	return DB.Self.Delete(&user).Error
}

//Update is to update userinfo in db
func (u *UserModel) Update() error {
	return DB.Self.Model(&u).Updates(
		UserModel{
			Username: u.Username,
			Password: u.Password,
		}).Error
	// return DB.Self.Save(u).Error
}

//GetUser will return userinfo by username
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username=?", username).First(&u)
	return u, d.Error
}

//ListUser list all users
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
	if err := DB.Self.Where(where).Limit(limit).Offset(offset).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}
	return users, count, nil
}

//Compare check password
func (u *UserModel) Compare(pwd string) error {
	return auth.Compare(u.Password, pwd)
}

//Encrypt will crypt user password
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

//Validate will valite the instence and return error
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
