package utils

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/zjxpcyc/wechat.v2/declares"
)

// Request 请求数据
// 请求远程 http 服务数据, 针对微信相关接口进行了特殊的处理
// 请求的地址, 方式等通过 api 指定, url 中的 search 参数通过 query 指定
// 如果请求中需要包含 body 内容, 则传入 body 参数, 没有传 nil 即可
// result 是可选参数, 用来承载或者格式化远程请求的结果, 比如 远程返回的实际上是 json 字串, 那么 result 可以为该 json 对应的 struct 指针
// data 返回值是远程请求的原始结果内容
func Request(api declares.API, query url.Values, body io.Reader, result ...interface{}) (data []byte, err error) {
	// 请求地址
	addr := api.URI
	if query != nil {
		apiURL, _ := url.Parse(api.URI)
		qry := apiURL.Query()
		for k := range query {
			qry.Set(k, query.Get(k))
		}
		apiURL.RawQuery = qry.Encode()
		addr = apiURL.String()
	}

	logger.Info("远程请求接口 URL ", addr)
	logger.Info("远程请求方法 ", api.Method)

	// 请求 Body
	var bodyData io.Reader
	if api.Method != http.MethodGet && body != nil {
		bodyData = body

		b := &bytes.Buffer{}
		io.Copy(b, body)
		logger.Info("远程请求体内容 ", b.String())
	} else {
		bodyData = nil
	}

	// 构造 http 请求
	var req *http.Request
	var res *http.Response
	client := new(http.Client)

	req, err = http.NewRequest(api.Method, addr, bodyData)
	if err != nil {
		logger.Error("初始化 http 客户端失败 ", err.Error())
		return
	}

	if api.ResponseType == declares.ResponseXML {
		req.Header.Add("Content-type", "text/xml")
	}

	res, err = client.Do(req)
	if err != nil {
		logger.Error("http 请求数据失败 ", err.Error())
		return
	}

	data, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		logger.Error("读取 http 请求结果失败 ", err.Error())
		return
	} else {
		logger.Info("远程请求结果 ", string(data))
	}

	// 格式化结果
	if result != nil && len(result) > 0 {
		if api.ResponseType == declares.ResponseJSON {
			err = json.Unmarshal(data, result[0])
			if err != nil {
				logger.Error("格式化 http 请求结果失败 ", err.Error())
				return
			}
		} else if api.ResponseType == declares.ResponseXML {
			err = xml.Unmarshal(data, result[0])
			if err != nil {
				logger.Error("格式化 http 请求结果失败 ", err.Error())
				return
			}
		}
	}

	return
}
