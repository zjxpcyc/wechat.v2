package wx

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/zjxpcyc/wechat.v2/utils"
)

// TplMessageData 模板消息节点
type TplMessageData struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// MiniProgramData 小程序内容
type MiniProgramData struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath,omitempty"`
}

// TplMessage 模板消息
type TplMessage struct {
	ToUser      string                    `json:"touser"`
	TemplateID  string                    `json:"template_id"`
	URL         string                    `json:"url"`
	MiniProgram *MiniProgramData          `json:"miniprogram,omitempty"`
	Data        map[string]TplMessageData `json:"data"`
}

// SendTplMessage 发送模板消息
func (t *Client) SendTplMessage(to, tplID, link string, data map[string]TplMessageData) error {
	api := API["tpl_message"]["send"]
	params := url.Values{}
	params.Set("access_token", t.getAccessToken())

	message := TplMessage{
		ToUser:     to,
		TemplateID: tplID,
		URL:        link,
		Data:       data,
	}

	b, err := json.Marshal(message)
	if err != nil {
		logger.Error("转换模板消息失败", err.Error())
		return err
	}

	var res map[string]interface{}
	_, err = utils.Request(api, params, bytes.NewBuffer(b), res)
	if err != nil {
		logger.Error("发送模板消息失败", err.Error())
		return err
	}

	if err = CheckResult(res); err != nil {
		return err
	}

	return nil
}
