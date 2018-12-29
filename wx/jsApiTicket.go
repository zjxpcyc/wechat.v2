package wx

import (
	"net/url"
	"time"

	"github.com/zjxpcyc/wechat.v2/utils"
)

// JT is the implementation of JSAPITicket
type JT struct {
	accessToken Scheduler
	task        *utils.ScheduleTask
	jsTicket    string
	last        time.Time
	expire      time.Duration
}

// NewJT returns new instance of JT
func NewJT(at Scheduler) *JT {
	jt := &JT{
		accessToken: at,
	}

	jt.task = utils.NewScheduleTask(func() time.Duration {
		var reTrySec int64 = 60
		ticket, expire, err := jt.getTicket()
		if err != nil {
			logger.Error("获取 JS Ticket 失败", err.Error())
			expire = reTrySec
		}

		// 不允许连续不断调用
		if expire == 0 {
			expire = reTrySec
		}

		jt.jsTicket = ticket
		jt.last = time.Now().Local()
		jt.expire = time.Duration(expire) * time.Second

		return jt.expire
	})

	jt.task.Run()

	return jt
}

// Result reutns the jsapi-ticket string
func (t *JT) Result() string {
	return t.jsTicket
}

// getTicket 获取 ticket
func (t *JT) getTicket() (string, int64, error) {
	api := API["jssdk"]["ticket"]
	params := url.Values{}
	params.Set("access_token", t.accessToken.Result())

	res := make(map[string]interface{})
	_, err := utils.Request(api, params, nil, &res)
	if err != nil {
		return "", 0, err
	}

	if err = CheckResult(res); err != nil {
		return "", 0, err
	}

	ticket, _ := res["ticket"].(string)
	expire, _ := res["expires_in"].(float64)
	return ticket, int64(expire), nil
}
