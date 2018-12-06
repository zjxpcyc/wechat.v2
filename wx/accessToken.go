package wx

import (
	"net/url"
	"time"

	"github.com/zjxpcyc/wechat.v2/utils"
)

// AT is the implementation of AccessToken
type AT struct {
	AppID       string
	appsecret   string
	accessToken string
	expire      time.Duration
	last        time.Time
	task        *utils.ScheduleTask
}

// NewAT returns new AT instance
func NewAT(appid, appsecret string) *AT {
	at := &AT{
		AppID:     appid,
		appsecret: appsecret,
	}

	at.task = utils.NewScheduleTask(func() time.Duration {
		var reTrySec int64 = 60
		token, expire, err := at.getToken()
		if err != nil {
			logger.Error("获取 Access-Token 失败", err.Error())
			expire = reTrySec
		}

		// 不允许连续不断调用
		if expire == 0 {
			expire = reTrySec
		}

		at.accessToken = token
		at.expire = time.Duration(expire) * time.Second
		at.last = time.Now().Local()

		return at.expire
	})

	at.task.Run()

	return at
}

// Result returns access-token
func (t *AT) Result() string {
	return t.accessToken
}

// getToken 获取 token
func (t *AT) getToken() (string, int64, error) {
	api := API["access_token"]["get"]
	params := url.Values{}
	params.Set("appid", t.AppID)
	params.Set("secret", t.appsecret)

	res := map[string]interface{}{}
	_, err := utils.Request(api, params, nil, &res)
	if err != nil {
		return "", 0, err
	}

	if err := CheckResult(res); err != nil {
		return "", 0, err
	}

	token, _ := res["access_token"].(string)
	expire, _ := res["expires_in"].(float64)
	return token, int64(expire), nil
}
