package mini

import (
	"github.com/zjxpcyc/tinylogger"
)

var logger tinylogger.LogService

// Client 微信小程序接口客户端
type Client struct {
	AppID string

	/*
	* certificate key 值如下:
	* 	appid 	开发者ID(AppID)
	*		secret 	开发者密码(AppSecret)
	 */
	certificate map[string]string
}

// NewClient 初始客户端
func NewClient(certificate map[string]string) *Client {
	cli := &Client{
		AppID:       certificate["appid"],
		certificate: certificate,
	}

	return cli
}

// SetLogger 设置日志记录器
func SetLogger(l tinylogger.LogService) {
	logger = l
}

// func init() {
// 	logger = utils.GetLogger()
// }
