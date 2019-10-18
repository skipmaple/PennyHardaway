package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	"log"
)

var db *gorm.DB

//func init() {
//	// open a db connection
//	var err error
//	db, err = gorm.Open("mysql", "root:root@/hardaway?charset=utf8&parseTime=True&loc=Local")
//	if err != nil {
//		panic("failed to connect database")
//	}
//	//Migrate the schema
//	//db.AutoMigrate(&todoModel{})
//}

func hello(c *gin.Context) {

	// 配置微信参数
	config := &wechat.Config {
		AppID:			"wxb49b036278e9e065",
		AppSecret:		"e91664672a12386530600cb5bb30ed84",
		Token: 			"skipmapledrewlee",
		EncodingAESKey:	"YCPnsvCg5oihIRrTALYDugXwwPekWcc3cjQXcUXefds",
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(c.Request, c.Writer)
	// 设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		// 回复消息: 演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	// 处理消息接收及回复
	err := server.Serve()
	if err != nil{
		log.Println("handle message receive err: ", err)
		return
	}

	// 发送恢复的消息
	server.Send()
}

func main() {
	router := gin.Default()

	router.Any("/wxrobot/", hello)
	router.Run(":80")
}