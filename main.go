package main

import (
	"fmt"

	//"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	//"github.com/rs/zerolog"
	//"github.com/rs/zerolog/log"
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
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	//log.Printf("[info] start http server listening %s", endPoint)
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC850}
	//output.FormatLevel = func(i interface{}) string {
	//	return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	//}
	//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC850})
	//log := zerolog.New(output).With().Caller().Timestamp().Logger()

	//log.Info().Msgf("[info] start http server listening %s", endPoint)
	//log.Debug().Msgf("[info] start http server listening %s", endPoint)
	//log.Fatal().Msgf("[info] start http server listening %s", endPoint)
	//log.Error().Msgf("type of log is %T", log)
	logging.Debug("hello world %s", endPoint)

	server.ListenAndServe()
}
