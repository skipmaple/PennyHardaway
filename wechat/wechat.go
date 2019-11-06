package wechat

import (
	"PennyHardway/pkg/setting"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/menu"
	"log"
)

func Setup() {
	// 配置微信参数
	config := &wechat.Config{
	AppID:          setting.WechatSetting.AppID,
	AppSecret:      setting.WechatSetting.AppSecret,
	Token:          setting.WechatSetting.Token,
	EncodingAESKey: setting.WechatSetting.EncodingAESKey,
	}
	wc := wechat.NewWechat(config)


	mu := wc.GetMenu()

	buttons := make([]*menu.Button, 1)
	btn := new(menu.Button)

	// 创建click类型菜单
	btn.SetClickButton("button1", "btn_key1")
	buttons[0] = btn

	// 设置btn为二级菜单
	btn2 := new(menu.Button)
	btn2.SetSubButton("subBtn", buttons)

	buttons2 := make([]*menu.Button, 1)
	buttons2[0] = btn2

	log.Println("buttons2", buttons2)
	// 发送请求
	err := mu.SetMenu(buttons2)
	if err != nil {
		//logging.Error("err = %v", err)
		log.Printf("set menu err: %v", err)
		return
	}
}