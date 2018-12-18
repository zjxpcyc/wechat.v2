package mini

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/zjxpcyc/gen"
)

// CheckResult 校验接口返回结果
func CheckResult(res map[string]interface{}) error {
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
	return errors.New(msg)
}

// DecryptData 小程序解密数据
// https://developers.weixin.qq.com/miniprogram/dev/api/signature.html
func DecryptData(encryptedData, sessionKey, iv string) (data map[string]interface{}, err error) {
	rawData, e := gen.Base64Decode(encryptedData)
	if e != nil {
		err = e
		logger.Error("小程序加密数据 Base64 Decode 失败", e.Error())
		return
	}

	rawKey, e := gen.Base64Decode(sessionKey)
	if e != nil {
		err = e
		logger.Error("小程序 session_key Base64 Decode 失败", e.Error())
		return
	}

	rawIV, e := gen.Base64Decode(iv)
	if e != nil {
		err = e
		logger.Error("小程序解密数据 iv Base64 Decode 失败", e.Error())
		return
	}

	decodeData, e := gen.CBC7Decrypt(rawData, rawKey, rawIV)
	if e != nil {
		err = e
		logger.Error("小程序解密数据失败", e.Error())
		return
	}

	err = json.Unmarshal(decodeData, &data)
	return
}
