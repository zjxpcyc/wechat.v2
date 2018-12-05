package mini

import (
	"net/http"

	"github.com/zjxpcyc/wechat.v2/declares"
)

// API 接口列表
var API = map[string]map[string]declares.API{
	"oauth2": map[string]declares.API{
		"session": declares.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code",
			ResponseType: declares.ResponseJSON,
		},
	},
}
