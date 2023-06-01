package global

import (
	"github.com/gin-gonic/gin"
	_ "go.uber.org/zap"
	"gorm.io/gorm"
	"test_go_mod/config"
	"test_go_mod/response"
)

var (
	Settings config.ServerConfig
	Response *response.Body
	C *gin.Context
	DB    *gorm.DB

)
