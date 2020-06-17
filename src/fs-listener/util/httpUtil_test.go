package util

import (
	"fs-listener/conf"
	"testing"
)

const (
	URL = "http://wechat.bluemoon.com.cn/wxplatform/qyweixin/msg-push/push-msg"
)

func TestPost(t *testing.T) {
	t.Run("post", func(t *testing.T) {
		msgText := &conf.MsgText{
			Content: "你的快递已到，请携带工卡前往邮件中心领取。",
		}

		msgData := &conf.MsgData{
			Touser:  "80468295",
			MsgType: "text",
			Agentid: 1000079,
			Text:    msgText,
		}

		msg := &conf.Msg{
			CorpId:  "wx36ef368cf28caea0",
			Agentid: 1000079,
			Data:    msgData,
		}

		//bytes, err := json.Marshal(msg)
		//if err != nil {
		//	panic(err)
		//}

		Post(URL, &msg, "")
	})
}
