package main

import (
	"fmt"
	loginservice "logistics/service"
	"logistics/util"
	"testing"
)

func TestRedisconn(t *testing.T) {
	//tel := "15925131827"
	//reg := regexp.MustCompile(`(13|15|18)[0-9]{9}`)
	//checktel := reg.Match([]byte(tel))
	//err := RD.Set(tel, reg, 60*time.Second).Err()
	//if err != nil {
	//
	//}
	//fmt.Println(checktel)
	//rd := redis.NewClient(&redis.Options{
	//	DB:       0,
	//	Addr:     "127.0.0.1:6379",
	//	Password: "",
	//})
	//a, err := rd.Get("a").Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(a)
	fmt.Println(util.Encode("sddfslkjds"))
	//48175b2815c0807f5cea2a1d2ab70233
	fmt.Println(len("48175b2815c0807f5cea2a1d2ab70233"))
	fmt.Println(util.Check("sddfslkjds", "48175b2815c0807f5cea2a1d2ab70233"))
	srv := loginservice.NewLoginService()
	b, err := srv.CheckRegistered("13025131827")
	fmt.Println(b, err)
}
