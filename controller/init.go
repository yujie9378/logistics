package main

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"log"
	"logistics/global"
)

func init() {
	err := MySQLInit()
	if err != nil {
		log.Panic("数据库连接失败", err)
	}
	err = RedisInit()
	if err != nil {
		log.Panic("Redis连接失败", err)
	}
}

var DB *gorm.DB
var RD *redis.Client

func MySQLInit() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/logistics?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	global.DB = db
	//log.Fatal("数据库连接成功")
	return nil
}
func RedisInit() error {
	Addr := "127.0.0.1:6379"
	rd := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: "",
		DB:       0,
	})
	_, err := rd.Ping().Result()
	if err != nil {
		return err
	}
	global.RD = rd
	return nil
}
