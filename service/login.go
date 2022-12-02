package loginservice

import (
	"fmt"
	"logistics/global"
	_type "logistics/type"
	"logistics/util"
	"math/rand"
	"time"
)

type Login struct {
}

// 登录的时候如果是采用验证码登录的话，（设置type1，type2），则无需密码，需要搜索数据库是否有该电话号码再发送；否则密码登录。
var loginService = &Login{}

func NewLoginService() ILogin {
	return loginService
}
func (l *Login) Register(user *_type.User) error {

	user.Password = util.Encode(user.Password)
	result := global.DB.Create(&user)
	//密码需要加密存入数据库。
	fmt.Println(user.ID, result.Error, result.RowsAffected)
	return nil
}

// 可以再加一个添加手机号码验证查询，一个手机号码只能注册一次，发送验证码之前先校验手机号是否注册过，如果注册过就不用发送验证码了..方法，独立开来，在controller层调用，这样login的时候也可以使用这个验证码功能。
func (l *Login) CheckRegistered(tel string) (bool, error) {
	db := global.DB
	result := map[string]interface{}{}
	err := db.Debug().Table("users").Select("id").Where("tel=?", tel).Find(&result).Error
	if err != nil {
		return false, fmt.Errorf("查询数据库失败：%w", err)
	}
	if _, ok := result["id"]; !ok {
		return false, nil
	}

	return true, nil
}

// SendVerifyCode 发送验证码功能
func (l *Login) SendVerifyCode(tel string) error {

	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(1000000)
	fmt.Println(code)
	rd := global.RD
	v, err := rd.Get(tel).Result()

	if v != "" {
		return fmt.Errorf("此手机号码已经发送过验证码")
	} else {
		err = rd.Set(tel, code, 160*time.Second).Err()
		if err != nil {
			return fmt.Errorf("验证码设置失败有误")
		}
	}
	fmt.Println("验证码发送成功")
	return nil
}
func (l *Login) Verify(tel string, code string) error {
	rd := global.RD
	c, err := rd.Get(tel).Result()
	if err != nil {
		return fmt.Errorf("该手机号未发送验证码")
	}
	if c != code {
		return fmt.Errorf("验证码错误！")
	}
	return nil
}
