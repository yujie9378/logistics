package main

import (
	"github.com/gin-gonic/gin"
	"logistics/service"
	_type "logistics/type"
	"net/http"
)

func main() {

	r := gin.Default()
	r.POST("/register", HandleRegister)
	r.Run(":8082")
}

func HandleRegister(context *gin.Context) {
	var user _type.User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"表单解析失败": err.Error()})
	}
	//把注册的信息写入数据库中
	srv := service.NewLoginService()
	err := srv.Register(DB, &user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"注册失败": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
