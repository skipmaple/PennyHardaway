package main

import (
	"PennyHardway/models"
	"PennyHardway/pkg/logging"
	"PennyHardway/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	//"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/silenceper/wechat"
	//"github.com/silenceper/wechat/message"
	//"log"
	"net/http"

	"PennyHardway/routers"
	"PennyHardway/pkg/setting"
)

var db *gorm.DB

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	util.Setup()
}

// @title PennyHardaway API
// @version 1.0

// @contact.name Drew Lee
// @contact.email skipmaple@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	router := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20


	server := &http.Server{
		Addr:              endPoint,
		Handler:           router,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	logging.Error(endPoint, " is running...")
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}