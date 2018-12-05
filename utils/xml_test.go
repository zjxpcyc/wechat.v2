package utils_test

import (
	"encoding/xml"
	"testing"

	"github.com/zjxpcyc/wechat.v2/utils"
)

func TestMarshalXML(t *testing.T) {
	m := utils.XMLMap(map[string]string{
		"appid":            "wx2421b1c4370ec43b",
		"attach":           "支付测试",
		"body":             "H5支付测试",
		"mch_id":           "10000100",
		"nonce_str":        "1add1a30ac87aa2db72f57a2375d8fec",
		"notify_url":       "http://wxpay.wxutil.com/pub_v2/pay/notify.v2.php",
		"openid":           "oUpF8uMuAJO_M2pxb1Q9zNjWeS6o",
		"out_trade_no":     "1415659990",
		"spbill_create_ip": "14.23.150.211",
		"total_fee":        "1",
		"trade_type":       "MWEB",
		"scene_info":       `<![CDATA[{"h5_info": {"type":"IOS","app_name": "王者荣耀","package_name": "com.tencent.tmgp.sgame"}}]]>`,
		"sign":             "0CB01533B8C1EF103065174F50BCA001",
	})

	dt, err := xml.Marshal(&m)
	if err != nil {
		t.Fatalf("TestMarshalXML fail : %v", err)
	}

	xmlStr := string(dt)
	if len(xmlStr) == 0 {
		t.Fatalf("TestMarshalXML fail : %s", xmlStr)
	}
}

func TestUnmarshalXML(t *testing.T) {
	xmlStr := `
	<xml>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str>
   <sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id>
   <trade_type><![CDATA[MWEB]]></trade_type>
   <mweb_url><![CDATA[https://wx.tenpay.com/cgi-bin/mmpayweb-bin/checkmweb?prepay_id=wx2016121516420242444321ca0631331346&package=1405458241]]></mweb_url>
	</xml>
	`

	dt := make(utils.XMLMap)
	err := xml.Unmarshal([]byte(xmlStr), &dt)

	if err != nil {
		t.Fatalf("TestUnmarshalXML fail : %v", err)
	}

	if len(dt) != 10 {
		t.Fatalf("TestUnmarshalXML fail : %v", dt)
	}
}
