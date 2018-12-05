package wx

import (
	"strconv"
	"time"

	"github.com/zjxpcyc/wechat.v2/utils"
)

// GetJSTicketSignature js-sdk signature
func (t *Client) GetJSTicketSignature(url string) map[string]string {
	noncestr := utils.RandString(16)
	timestamp := strconv.FormatInt(time.Now().Local().Unix(), 10)

	signature := JSShareTicketSignature(url, noncestr, t.getJSTicket(), timestamp)

	return map[string]string{
		"noncestr":  noncestr,
		"timestamp": timestamp,
		"signature": signature,
		"appId":     t.certificate["appid"],
	}
}
