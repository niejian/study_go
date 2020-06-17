package conf

// 微信消息结构体

/*
"corpId": "wx36ef368cf28caea0",
 "agentId": "1000079",
 "data": {
  "touser": "80468295",
  "msgtype": "text",
  "agentid": 1000079,
  "text": {
   "content": "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看< a href= \"http://work.weixin.qq.com\">邮件中心视频实况</ a>，聪明避开排队。"
  }
 }
*/

type MsgText struct {
	Content string `json:"content"`
}

type MsgData struct {
	Touser string `json:"touser"`
	MsgType string `json:"msgtype"`
	Agentid int32 `json:"agentid"`
	Text interface{} `json:"text"`
}

type Msg struct {
	CorpId string `json:"corpId"`
	Agentid int32 `json:"agentId"`
	Data interface{} `json:"data"`
}


