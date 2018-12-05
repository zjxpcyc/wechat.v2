package declares

// API API 接口信息
type API struct {
	Method       string
	URI          string
	ResponseType string
}

// ResponseJSON 返回 json
const ResponseJSON = "json"

// ResponseXML 返回 xml
const ResponseXML = "xml"
