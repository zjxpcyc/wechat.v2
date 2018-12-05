package wechat

import (
	"encoding/xml"

	"github.com/zjxpcyc/wechat.v2/utils"
)

// TransferXMLToStringMap 转换 xml 为 map
func TransferXMLToStringMap(xmlStr string) (map[string]string, error) {
	target := new(utils.XMLMap)

	err := xml.Unmarshal([]byte(xmlStr), target)
	if err != nil {
		return nil, err
	}

	return map[string]string(*target), nil
}
