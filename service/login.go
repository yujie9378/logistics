package service

import (
	"fmt"
	"gorm.io/gorm"
	_type "logistics/type"
)

type Login struct {
}

var loginService = &Login{}

func NewLoginService() IService {
	return loginService
}
func (l *Login) Register(db *gorm.DB, user *_type.User) error {
	result := db.Create(&user)
	fmt.Println(user.ID, result.Error, result.RowsAffected)
	return nil
}
