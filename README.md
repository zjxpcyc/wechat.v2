# 微信公众号接口
第一版的实现了部分接口。但是不支持现有业务的拆分要求。因此这是第二个版本主要是在第一版本基础上加上了部分接口拆分的支持。同时优化了部分方法内部设计。
接口初始化与调用方式均与第一版有部分不同。

## 简易使用说明
**安装**
```bash
go get "github.com/zjxpcyc/wechat.v2"
```


```go
import wechat "github.com/zjxpcyc/wechat.v2"

...

// 微信开发者信息
certificate := map[string]string{
	"appid"  "",
	"secret" "",
	"token" "",
	"aeskey": "",
}

// 一个日志记录器
// 需要实现 github.com/zjxpcyc/tinylogger.LogService
log := &AnLogger{}
wechat.SetLogger(log)

// 初始化
wx := wechat.NewClient(certificate)

// 如果 access-token 的获取是通过一个通用的中控器实现
// 那么这里需要设置 wechat RefresAccessToken 方法
// foo = interface AccessToken
wx.SetAccessToken(foo)

// 同理, 如果 jsapi-ticket 也是这种模式
// bar = func(accessToken string) string
wx.SetJSAPITicket(bar)
```

**示例**

1. 首次接入
```golang
func SomeControllerMethod() {
  // fromRequest 代表为微信发送的 http get 请求
  signature := fromRequest.Query("signature")
  timestamp := fromRequest.Query("timestamp")
  nonce := fromRequest.Query("nonce")
  echostr := fromRequest.Query("echostr")

  if wx.Signature(timestamp, nonce) == signature {
    response(echostr)
  }
}
```

2. 获取用户 OpenID
```golang
func SomeControllerMethod() {
  // fromRequest 代表为微信发送的 http get 请求
  // code 需要前端传送过来
  code := fromRequest.Query("code")

  openID, err := wx.GetOpenID(code)
  if err != nil {
    // TODO something
  }

  response(openID)
}
```

3. 获取用户详细信息
```golang
func SomeControllerMethod() {
  // fromRequest 代表为微信发送的 http get 请求
  // code 需要前端传送过来
  code := fromRequest.Query("code")

  user, err := wx.GetUserInfo(code)
  if err != nil {
    // TODO something
  }

  response(user)
}
```
