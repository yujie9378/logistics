package main

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"log"
)

func init() {
	err := MySQLInit()
	if err != nil {
		log.Println("数据库连接失败", err)
	}

}

var DB *gorm.DB

func MySQLInit() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/logistics?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	//log.Fatal("数据库连接成功")
	return nil
}
func GetMySQLDB() *gorm.DB {
	return DB
}
