package wx

import (
	"net/url"
	"strings"

	"github.com/zjxpcyc/wechat.v2/utils"
)

// RefreshMenu 刷新菜单
func (t *Client) RefreshMenu(menu string) error {
	logger.Info("准备刷新菜单")
	logger.Info("接收到菜单信息:", menu)

	accessToken := t.getAccessToken()

	api := API["menu"]["delete"]
	params := url.Values{}
	params.Set("access_token", accessToken)

	_, err := utils.Request(api, params, nil)
	if err != nil {
		logger.Error("清除原有菜单失败", err.Error())
		return err
	}

	api = API["menu"]["create"]
	params = url.Values{}
	params.Set("access_token", accessToken)

	var res map[string]interface{}
	_, err = utils.Request(api, params, strings.NewReader(menu), &res)
	if err != nil {
		logger.Error("创建菜单失败", err.Error())
		return err
	}

	if err = CheckResult(res); err != nil {
		return err
	}

	logger.Info("创建菜单成功")

	return nil
}

// GetMenu 获取菜单
func (t *Client) GetMenu() (res map[string]interface{}, err error) {
	logger.Info("准备获取菜单")

	api := API["menu"]["get"]
	params := url.Values{}
	params.Set("access_token", t.getAccessToken())

	_, err = utils.Request(api, params, nil, &res)
	if err != nil {
		logger.Error("查询菜单失败", err.Error())
		return nil, err
	}

	return
}
