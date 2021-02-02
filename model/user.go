package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//Gorm模型定义
type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

//通过ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

//设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

//check密码是否正确
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
