package main

import (
	"fmt"
	"test_go_mod/global"
	"test_go_mod/initialize"
	"test_go_mod/models"
)

func main() {
	initialize.InitConfig()
	initialize.InitMysqlDB()
	err := global.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("失败")
	} else {
		fmt.Println("成功")
	}
}
