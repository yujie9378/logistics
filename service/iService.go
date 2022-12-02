package loginservice

import (
	_type "logistics/type"
)

type ILogin interface {
	Register(user *_type.User) error
	SendVerifyCode(tel string) error
	Verify(tel string, code string) error
	CheckRegistered(tel string) (bool, error)
}
