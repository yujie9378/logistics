package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"logistics/service"
	_type "logistics/type"
	"net/http"
	"regexp"
)

func main() {

	r := gin.Default()
	r.POST("/register", HandleRegister)
	r.POST("/sendCode", HandleSendVerifyCode)
	r.POST("/verify", HandleVerify)
	r.Run(":8082")
}
func HandleVerify(context *gin.Context) {
	tel := context.PostForm("tel")
	code := context.PostForm("code")
	srv := loginservice.NewLoginService()
	err := srv.Verify(tel, code)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"验证码校验错误": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "验证码校验成功"})

}
func HandleSendVerifyCode(context *gin.Context) {
	tel := context.PostForm("tel")
	if tel == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "电话号码为空"})
		return
	}
	reg := regexp.MustCompile(`(13|15|18)[0-9]{9}`)
	checktel := reg.Match([]byte(tel))
	fmt.Println(checktel)
	if !checktel {
		context.JSON(http.StatusBadRequest, gin.H{"error": "电话号码错误"})
		return
	}
	srv := loginservice.NewLoginService()
	b, err := srv.CheckRegistered(tel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"查询手机号码失败失败": err})
		return
	}
	if b {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "此电话号码已存在"})
		return
	}
	err = srv.SendVerifyCode(tel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"发送验证码失败": err})
		return
	}
	fmt.Println(tel)
	context.JSON(http.StatusOK, gin.H{"message": "验证码发送成功"})

}

func HandleRegister(context *gin.Context) {
	var user _type.User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"表单解析失败": err.Error()})
		return
	}
	//把注册的信息写入数据库中
	srv := loginservice.NewLoginService()
	err := srv.Register(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"注册失败": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
