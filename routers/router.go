package routers

import (
	"PennyHardway/middleware/jwt"
	"PennyHardway/routers/api"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

	"PennyHardway/pkg/setting"
	_ "PennyHardway/docs"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Any("/wxrobot/", hello)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//	这里的路由需要加token才可以访问
	}


	return r
}

func hello(c *gin.Context) {

	// 配置微信参数
	config := &wechat.Config {
		AppID:			setting.WechatSetting.AppID,
		AppSecret:		setting.WechatSetting.AppSecret,
		Token: 			setting.WechatSetting.Token,
		EncodingAESKey:	setting.WechatSetting.EncodingAESKey,
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

