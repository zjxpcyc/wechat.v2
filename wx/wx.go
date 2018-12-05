package wx

import (
	"sort"
	"strings"

	"github.com/zjxpcyc/gen"
	"github.com/zjxpcyc/tinylogger"
)

var logger tinylogger.LogService

// Client 微信公众号接口客户端
type Client struct {
	AppID       string
	appsecret   string
	accessToken AccessToken
	jsTicket    JSAPITicket

	/*
	* certificate key 值如下:
	* 	appid 	开发者ID(AppID)
	*		secret 	开发者密码(AppSecret)
	*   token 	令牌(Token)
	*   aeskey 	消息加解密密钥 (EncodingAESKey)
	 */
	certificate map[string]string
}

// NewClient 初始客户端
func NewClient(certificate map[string]string) *Client {
	cli := &Client{
		AppID:       certificate["appid"],
		appsecret:   certificate["secret"],
		certificate: certificate,
	}

	return cli
}

// SetAccessToken 设置刷新 AccessToken 的函数
// 添加该函数的目的, 主要是为了实现可以接收第三方的 access-token 的刷新任务
// 这样可以支持 access-token 与当前库封装分离的特性
// 如果不设置, 那么系统会使用默认的 access-token 刷新机制
func (t *Client) SetAccessToken(at ...AccessToken) {
	if at == nil || len(at) == 0 {
		t.accessToken = NewAT(t.AppID, t.appsecret)
	} else {
		t.accessToken = at[0]
	}
}

// SetJSAPITicket 设置刷新 jsapi-ticket 函数
func (t *Client) SetJSAPITicket(jt ...JSAPITicket) {
	if jt == nil || len(jt) == 0 {
		t.jsTicket = NewJT(t.accessToken)
	} else {
		t.jsTicket = jt[0]
	}
}

// Signature 初始接入时校验
func (t *Client) Signature(timestamp, nonce string) string {
	token := t.certificate["token"]
	strs := sort.StringSlice{token, timestamp, nonce}
	sort.Strings(strs)

	return gen.SHA1(strings.Join([]string(strs), ""))
}

// getAccessToken 获取 access-token
func (t *Client) getAccessToken() string {
	if t.accessToken == nil {
		return ""
	}

	return t.accessToken.String()
}

// getJSTicket 获取 jsapi-ticket
func (t *Client) getJSTicket() string {
	if t.jsTicket != nil {
		return ""
	}

	return t.jsTicket.String()
}

// SetLogger 设置日志记录器
func SetLogger(l tinylogger.LogService) {
	logger = l
}
