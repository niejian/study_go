package conf

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	msgText := &MsgText{
		Content: "你的快递已到，请携带工卡前往邮件中心领取。",
	}

	msgData := &MsgData{
		Touser:  "80468295",
		MsgType: "text",
		Agentid: 1000079,
		Text:    msgText,
	}

	msg := &Msg{
		CorpId:  "wx36ef368cf28caea0",
		Agentid: 1000079,
		Data:    msgData,
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))



}
