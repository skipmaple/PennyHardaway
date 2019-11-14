package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"PennyHardway/models"
	"PennyHardway/pkg/logging"
	"PennyHardway/pkg/setting"
	"PennyHardway/pkg/util"
	"PennyHardway/routers"
)

var db *gorm.DB

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	util.Setup()
	//wechat.Setup()
}

// @title PennyHardaway API
// @version 1.0
// @contact.name Drew Lee
// @contact.email skipmaple@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	//router := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	maxHeaderBytes := 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		logging.Error("Server err: %v", err)
	}
}
