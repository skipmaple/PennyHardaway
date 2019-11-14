package routers

import (
	"PennyHardway/middleware/jwt"
	"PennyHardway/routers/api"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	_ "github.com/silenceper/wechat/server"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

	_ "PennyHardway/docs"
	"PennyHardway/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/wxrobot/", hello)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//	这里的路由需要加token才可以访问
	}

	return r
}

func hello(c *gin.Context) {
	// 配置微信参数
	config := &wechat.Config{
		AppID:          setting.WechatSetting.AppID,
		AppSecret:      setting.WechatSetting.AppSecret,
		Token:          setting.WechatSetting.Token,
		EncodingAESKey: setting.WechatSetting.EncodingAESKey,
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(c.Request, c.Writer)

	// 设置接收消息的处理方法
	server.SetMessageHandler(func(v message.MixMessage) *message.Reply {

		switch v.MsgType {
		// 文本消息
		case message.MsgTypeText:
			//	do something
			// 回复消息: 演示回复用户发送的消息
			text := message.NewText("您输入了: " + v.Content)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 图片消息
		case message.MsgTypeImage:
			//	do something
			text := message.NewText("您发送了图片")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 语音消息
		case message.MsgTypeVoice:
			//	do something
			text := message.NewText("您发送了语音")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 视频消息
		case message.MsgTypeVideo:
			//	do something

		// 小视频消息
		case message.MsgTypeShortVideo:
			//	do something

		// 地理位置消息
		case message.MsgTypeLocation:
			//	do something

		// 链接消息
		case message.MsgTypeLink:
			//	do something

		// 事件推送消息
		case message.MsgTypeEvent:
			//	do something
			switch v.Event {
				// 订阅
				case message.EventSubscribe:
					// do something
					text := message.NewText("欢迎入队～")
					return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

				// 取消订阅
				case message.EventUnsubscribe:
					// do something
					text := message.NewText("拜拜～")
					return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
					
				// 用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
				case message.EventScan:
					// do something

				// 上报地理位置事件
				case message.EventLocation:
					// do something

				// 点击菜单拉取消息时的事件推送
				case message.EventClick:
					// do something

				// 点击菜单跳转链接时的事件推送
				case message.EventView:
					// do something

				// 扫码推事件的事件推送
				case message.EventScancodePush:
					// do something

				// 扫码推事件且弹出"消息接收中"的事件推送
				case message.EventScancodeWaitmsg:
					// do something

				// 弹出系统拍照发图的事件推送
				case message.EventPicSysphoto:
					// do something

				// 弹出拍照或者相册发图的事件推送
				case message.EventPicPhotoOrAlbum:
					// do something

				// 弹出微信相册发图器的事件推送
				case message.EventPicWeixin:
					// do something

				// 弹出地理位置选择器的事件推送
				case message.EventLocationSelect:
					// do something
			}
		}
		return nil
	})

	// 处理消息接收及回复
	err := server.Serve()
	if err != nil {
		log.Println("handle message receive err: ", err)
		return
	}

	// 发送恢复的消息
	server.Send()
}
