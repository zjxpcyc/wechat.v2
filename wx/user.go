package wx

import (
	"net/url"

	"github.com/zjxpcyc/wechat.v2/utils"
)

// GetUserDetail 获取用户详情(UUID)
func (t *Client) GetUserDetail(openID string) (map[string]interface{}, error) {
	logger.Info("获取用户详情(UUID): openID=" + openID)

	api := API["user"]["detail"]
	params := url.Values{}
	params.Set("access_token", t.getAccessToken())
	params.Set("openid", openID)

	res := make(map[string]interface{})
	_, err := utils.Request(api, params, nil, &res)
	if err != nil {
		logger.Error("获取用户信息(UUID) 失败, ", err.Error())
		return nil, err
	}

	if err = CheckResult(res); err != nil {
		return nil, err
	}

	return res, nil
}
