package server

import (
	wechat_work_bot "github.com/xiaofeige/wechat-work-bot"
)

type WxBotServer struct {
	*wechat_work_bot.Robot
}

var (
	gServer *WxBotServer
)

func init() {
	gServer = &WxBotServer{}

	defaultWxWorkBotConfig := wechat_work_bot.RobotConfig{
		ReceiveId:   "",
		Token:       "",
		TokenAESKey: "",
		WebHookUrl:  "",
		BindAddr:    "127.0.0.1:7788",
		UrlBase:     "",
		Handler:     nil,
		Debugger:    nil,
		ErrLogger:   nil,
	}
	wxWorkBot, err := wechat_work_bot.NewWxWorkRobot(defaultWxWorkBotConfig)
	if err != nil {
		panic(err)
	}
	gServer.Robot = wxWorkBot
}

func AddHandler(handler wechat_work_bot.MsgHandler) {
	gServer.AddHandler(handler)
}

func Start() {

	// 注册自己的处理函数

	gServer.Start()
}
