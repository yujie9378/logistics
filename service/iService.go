package service

import (
	"gorm.io/gorm"
	_type "logistics/type"
)

type IService interface {
	Register(db *gorm.DB, user *_type.User) error
}
