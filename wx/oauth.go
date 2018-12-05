package wx

import (
	"net/url"

	"github.com/zjxpcyc/wechat.v2/utils"
)

// GetOpenID 获取用户 OpenID
func (t *Client) GetOpenID(code string) (string, error) {
	logger.Info("获取用户 OpenID: code=" + code)

	res, err := t.getOauthToken(code)
	if err != nil {
		return "", err
	}

	return res["openid"].(string), nil
}

// GetUserInfo 获取用户详情
func (t *Client) GetUserInfo(code string) (res map[string]interface{}, err error) {
	logger.Info("获取用户详情: code=" + code)

	// 依据 code 获取 openid, access_token
	res, err = t.getOauthToken(code)
	if err != nil {
		return
	}

	openID := res["openid"].(string)
	token := res["access_token"].(string)

	// 再依据 openid, access_token 获取详情
	api := API["oauth2"]["userinfo"]
	params := url.Values{}
	params.Set("access_token", token)
	params.Set("openid", openID)
	_, err = utils.Request(api, params, nil, &res)
	if err != nil {
		logger.Error("获取 Oauth2 用户信息 失败, ", err.Error())

		// 即使失败也会返回 openid
		res["openid"] = openID
		return
	}

	if err = CheckResult(res); err != nil {
		return
	}

	return
}

// getOauthToken 获取 Oauth Token
// 没有对 Token 缓存, 官方声明有  7200s 的生存周期, 实际上是用不到
func (t *Client) getOauthToken(code string) (res map[string]interface{}, err error) {
	logger.Info("获取Oauth Token: code=" + code)

	// 依据 code 获取 openid, access_token
	api := API["oauth2"]["access_token"]
	params := url.Values{}
	params.Set("appid", t.certificate["appid"])
	params.Set("secret", t.certificate["secret"])
	params.Set("code", code)

	_, err = utils.Request(api, params, nil, &res)
	if err != nil {
		logger.Error("获取 Oauth2 Access-Token 失败, ", err.Error())
		return
	}

	if err = CheckResult(res); err != nil {
		return
	}

	return
}
