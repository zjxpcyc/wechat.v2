package wx

import (
	"net/http"

	"github.com/zjxpcyc/wechat.v2/declares"
)

// API 接口列表
var API = map[string]map[string]declares.API{
	"access_token": map[string]declares.API{
		"get": declares.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET",
			ResponseType: declares.ResponseJSON,
		},
	},
	"oauth2": map[string]declares.API{
		"access_token": declares.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code",
			ResponseType: declares.ResponseJSON,
		},
		"refresh_token": declares.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN",
			ResponseType: declares.ResponseJSON,
		},
		"auth": declares.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID",
			ResponseType: declares.ResponseJSON,
		},
		"userinfo": declares.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN",
			ResponseType: declares.ResponseJSON,
		},
	},
	"qrcode": map[string]declares.API{
		"create": declares.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=ACCESS_TOKEN",
			ResponseType: declares.ResponseJSON,
		},
	},
	"user": map[string]declares.API{
		"detail": declares.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN",
			ResponseType: declares.ResponseJSON,
		},
	},
	"tpl_message": map[string]declares.API{
		"send": declares.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN",
			ResponseType: declares.ResponseJSON,
		},
	},
	"menu": map[string]declares.API{
		"create": declares.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN",
			ResponseType: declares.ResponseJSON,
		},
		"delete": declares.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN",
			ResponseType: declares.ResponseJSON,
		},
		"get": declares.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN",
			ResponseType: declares.ResponseJSON,
		},
	},
	"jssdk": map[string]declares.API{
		"ticket": declares.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=jsapi",
			ResponseType: declares.ResponseJSON,
		},
	},
}
