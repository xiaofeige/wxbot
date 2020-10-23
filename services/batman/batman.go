package batman

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	wechat_work_bot "github.com/xiaofeige/wechat-work-bot"
	"github.com/xiaofeige/wxbot/server"
)

func init() {

	//viper 读取配置
	robotId := viper.GetString("batman.robot_id")
	if robotId == "" {
		panic("batman robot id could be empty")
	}

	h, err := NewBatmanHandler(robotId)
	if err != nil {
		fmt.Println("new batman handler err:", err)
		return
	}

	server.AddHandler(h)
}

type BatmanHandler struct {
	*wechat_work_bot.BaseHandler

	robotId string
}

func NewBatmanHandler(robotId string) (*BatmanHandler, error) {
	h := &BatmanHandler{
		robotId: robotId,
	}
	return h, nil
}

func (h *BatmanHandler) RobotId() string {
	// webhook: https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=da1b0595-cd46-4b43-bc3a-b04e0e3a2bca
	return h.robotId
}

func (h *BatmanHandler) OnMsgRecv(ctx context.Context, msg *wechat_work_bot.CallBackReq) (rsp *wechat_work_bot.CallBackRsp, err error) {
	rsp = &wechat_work_bot.CallBackRsp{
		StrMsgType: wechat_work_bot.WxTextMsg,
		//Markdown:      MarkdownElem{},
		Text: wechat_work_bot.TextRspElem{
			Content: wechat_work_bot.CDATA{Value: "I'm rich!!!"},
			//MentionedList: MentionedListElem{},
		},
		BIgnore: false,
	}

	return rsp, nil
}

func (h *BatmanHandler) OnEnterChat(ctx context.Context, msg *wechat_work_bot.CallBackReq) (rsp *wechat_work_bot.CallBackRsp, err error) {

	return &wechat_work_bot.CallBackRsp{BIgnore: true}, nil
}

func (h *BatmanHandler) OnAddToChat(ctx context.Context, msg *wechat_work_bot.CallBackReq) (rsp *wechat_work_bot.CallBackRsp, err error) {

	return &wechat_work_bot.CallBackRsp{BIgnore: true}, nil
}

func (h *BatmanHandler) OnDeletedFromChat(ctx context.Context, msg *wechat_work_bot.CallBackReq) (rsp *wechat_work_bot.CallBackRsp, err error) {

	return &wechat_work_bot.CallBackRsp{BIgnore: true}, nil
}

func (h *BatmanHandler) OnAttachmentEvent(ctx context.Context, msg *wechat_work_bot.CallBackReq) (rsp *wechat_work_bot.CallBackRsp, err error) {

	return &wechat_work_bot.CallBackRsp{BIgnore: true}, nil
}
