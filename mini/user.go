package mini

import (
	"net/url"

	"github.com/zjxpcyc/wechat.v2/utils"
	"github.com/zjxpcyc/wechat/core"
)

// GetOpenID 获取用户 OpenID
func (t *Client) GetOpenID(code string) (res map[string]interface{}, err error) {
	logger.Info("获取用户 OpenID: code=" + code)

	api := API["oauth2"]["session"]
	query := url.Values{}
	query.Set("appid", t.certificate["appid"])
	query.Set("secret", t.certificate["secret"])
	query.Set("js_code", code)

	_, err = utils.Request(api, query, nil, &res)
	if err != nil {
		logger.Error("获取 登录凭证 失败, ", err.Error())
		return nil, err
	}

	if err = CheckResult(res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetUserFromEncryptData 解析加密数据
func (t *Client) GetUserFromEncryptData(encryptedData, sessionKey, iv string) (map[string]interface{}, error) {
	res, err := core.DecodeMiniData(encryptedData, iv, sessionKey)
	if err != nil {
		logger.Error("解密小程序数据失败", err.Error())
		return nil, err
	}

	return res, nil
}
