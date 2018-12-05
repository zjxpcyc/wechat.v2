package utils

import (
	"encoding/xml"
	"io"
	"regexp"
)

// 简易版本的 xml 解析, 实现将 单层 xml 解析到 map[string]string

// XMLMap  xml 转换 map
type XMLMap map[string]string

// XMLElement xml 转换 element
type XMLElement struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

// XMLCDATA xml 转换 cdata
type XMLCDATA struct {
	XMLName xml.Name
	Value   string `xml:",cdata"`
}

// UnmarshalXML xml.Unmarshaler 接口实现
// 参考 https://stackoverflow.com/questions/30928770/marshall-map-to-xml-in-go
func (m *XMLMap) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	for {
		ele := XMLElement{}
		err := d.Decode(&ele)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		// 忽略 Space
		(*m)[ele.XMLName.Local] = ele.Value
	}

	return nil
}

// MarshalXML xml.Marshaler 接口实现
// 参考 https://stackoverflow.com/questions/30928770/marshall-map-to-xml-in-go
func (m *XMLMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil || len(*m) == 0 {
		return nil
	}

	if start.Name.Local == "XMLMap" {
		start.Name.Local = "xml"
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	regx, err := regexp.Compile(`<!\[CDATA\[(.*)\]\]>`)
	if err != nil {
		return err
	}

	for k, v := range *m {
		res := regx.FindAllStringSubmatch(v, -1)
		if res != nil && len(res) > 0 {
			t := XMLCDATA{
				XMLName: xml.Name{
					Space: "",
					Local: k,
				},
				Value: res[0][1],
			}

			if err := e.Encode(&t); err != nil {
				return err
			}
		} else {
			t := XMLElement{
				XMLName: xml.Name{
					Space: "",
					Local: k,
				},
				Value: v,
			}

			if err := e.Encode(&t); err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(start.End())
}
