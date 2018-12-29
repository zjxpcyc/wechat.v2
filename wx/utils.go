package wx

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/zjxpcyc/gen"
)

// CheckResult 校验接口返回是否正确
func CheckResult(res map[string]interface{}) error {
	logger.Info("接口返回结果: ", res)

	if res == nil {
		return nil
	}

	errcode, _ := res["errcode"]
	errmsg, _ := res["errmsg"]
	if errcode == nil {
		return nil
	}

	err, _ := errcode.(float64)
	errNum := int(err)

	if errNum == 0 {
		return nil
	}

	msg, _ := errmsg.(string)
	logger.Error("接口返回错误: " + strconv.Itoa(errNum) + "-" + msg)
	return errors.New(strconv.Itoa(errNum) + "-" + msg)
}

// JSShareTicketSignature 计算 js-ticket signature
func JSShareTicketSignature(url, noncestr, ticket, timestamp string) string {
	willSign := []string{
		"noncestr=" + noncestr,
		"timestamp=" + timestamp,
		"url=" + url,
		"jsapi_ticket=" + ticket,
	}
	sort.Strings(willSign)
	str2Sign := strings.Join(willSign, "&")

	return gen.SHA1(str2Sign)
}
