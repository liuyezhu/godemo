package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test_go_mod/global"
	"test_go_mod/initialize"
	"test_go_mod/router"
)

func main() {
	//初始化yaml配置
	initialize.InitConfig()
	initialize.InitMysqlDB()
	r := gin.Default()
	router.Http(r)
	r.Run(fmt.Sprintf(":%d", global.Settings.Port))
}
