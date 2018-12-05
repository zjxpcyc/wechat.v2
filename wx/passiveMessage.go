package wx

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/lunny/log"
)

// cdata CDATA
type cdata struct {
	Value string `xml:",cdata"`
}

// PassiveMessageText 文本
type PassiveMessageText struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   cdata    `xml:"ToUserName"`
	FromUserName cdata    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      cdata    `xml:"MsgType"`
	Content      cdata    `xml:"Content"`
}

// PassiveMessageImage 图片
type PassiveMessageImage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   cdata    `xml:"ToUserName"`
	FromUserName cdata    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      cdata    `xml:"MsgType"`
	MediaID      cdata    `xml:"MediaId"`
}

// PassiveMessageVoice 语音
type PassiveMessageVoice struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   cdata    `xml:"ToUserName"`
	FromUserName cdata    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      cdata    `xml:"MsgType"`
	MediaID      cdata    `xml:"MediaId"`
}

// PassiveMessageVideo 视频
type PassiveMessageVideo struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   cdata    `xml:"ToUserName"`
	FromUserName cdata    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      cdata    `xml:"MsgType"`
	MediaID      cdata    `xml:"MediaId"`
	Title        cdata    `xml:"Title"`
	Description  cdata    `xml:"Description"`
}

// PassiveMessageMusic 音乐
type PassiveMessageMusic struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   cdata    `xml:"ToUserName"`
	FromUserName cdata    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      cdata    `xml:"MsgType"`
	Title        cdata    `xml:"Title"`
	Description  cdata    `xml:"Description"`
	MusicURL     cdata    `xml:"MusicURL"`
	HQMusicURL   cdata    `xml:"HQMusicUrl"`
	ThumbMediaID cdata    `xml:"ThumbMediaId"`
}

// ArticleItem 图文节点
type ArticleItem struct {
	Title       cdata `xml:"Title"`
	Description cdata `xml:"Description"`
	PicURL      cdata `xml:"PicUrl"`
	URL         cdata `xml:"Url"`
}

// PassiveMessageNews 图文
type PassiveMessageNews struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   cdata         `xml:"ToUserName"`
	FromUserName cdata         `xml:"FromUserName"`
	CreateTime   int64         `xml:"CreateTime"`
	MsgType      cdata         `xml:"MsgType"`
	ArticleCount int           `xml:"MediaId"`
	Articles     []ArticleItem `xml:"Articles>item"`
}

// ResponseMessageText 回复文本消息
func (t *Client) ResponseMessageText(to, message string) ([]byte, error) {
	logger.Info("(被动)待反馈文本消息: ", fmt.Sprintf("%s", message))

	data := PassiveMessageText{
		ToUserName:   cdata{to},
		FromUserName: cdata{t.certificate["wxid"]},
		MsgType:      cdata{"text"},
		CreateTime:   time.Now().Local().Unix(),
		Content:      cdata{message},
	}

	return t.getXMLStringOfMessage(data)
}

// ResponseMessageImage 回复图片消息
func (t *Client) ResponseMessageImage(to, media string) ([]byte, error) {
	log.Info("(被动)待反馈图片消息: ", fmt.Sprintf("%s", media))

	data := PassiveMessageImage{
		ToUserName:   cdata{to},
		FromUserName: cdata{t.certificate["wxid"]},
		MsgType:      cdata{"image"},
		CreateTime:   time.Now().Local().Unix(),
		MediaID:      cdata{media},
	}

	return t.getXMLStringOfMessage(data)
}

// ResponseMessageVoice 回复音频消息
func (t *Client) ResponseMessageVoice(to, media string) ([]byte, error) {
	log.Info("(被动)待反馈音频消息: ", fmt.Sprintf("音频ID %s", media))

	data := PassiveMessageVoice{
		ToUserName:   cdata{to},
		FromUserName: cdata{t.certificate["wxid"]},
		MsgType:      cdata{"voice"},
		CreateTime:   time.Now().Local().Unix(),
		MediaID:      cdata{media},
	}

	return t.getXMLStringOfMessage(data)
}

// ResponseMessageVideo 回复视频消息
func (t *Client) ResponseMessageVideo(to, media, title, desc string) ([]byte, error) {
	log.Info("(被动)待反馈视频消息: ", fmt.Sprintf("视频ID %s , 标题 %s , 描述 %s ", media, title, desc))

	data := PassiveMessageVideo{
		ToUserName:   cdata{to},
		FromUserName: cdata{t.certificate["wxid"]},
		MsgType:      cdata{"video"},
		CreateTime:   time.Now().Local().Unix(),
		MediaID:      cdata{media},
		Title:        cdata{title},
		Description:  cdata{desc},
	}

	return t.getXMLStringOfMessage(data)
}

// ResponseMessageMusic 回复音乐消息
func (t *Client) ResponseMessageMusic(to, music, title, desc string, others ...string) ([]byte, error) {
	log.Info("(被动)待反馈音乐消息: ", fmt.Sprintf("音乐 %s , 标题 %s , 描述 %s, 其他: %v ", music, title, desc, others))

	othLen := len(others)
	hqMusicURL := ""
	thumbMediaID := ""
	if othLen > 0 {
		hqMusicURL = others[0]

		if othLen > 1 {
			thumbMediaID = others[1]
		}
	}

	data := PassiveMessageMusic{
		ToUserName:   cdata{to},
		FromUserName: cdata{t.certificate["wxid"]},
		MsgType:      cdata{"music"},
		CreateTime:   time.Now().Local().Unix(),
		Title:        cdata{title},
		Description:  cdata{desc},
		MusicURL:     cdata{music},
		HQMusicURL:   cdata{hqMusicURL},
		ThumbMediaID: cdata{thumbMediaID},
	}

	return t.getXMLStringOfMessage(data)
}

// ResponseMessageNews 回复图文消息
func (t *Client) ResponseMessageNews(to string, articles []map[string]string) ([]byte, error) {
	log.Info("(被动)待反馈图文消息: ", fmt.Sprintf("%v", articles))

	num := len(articles)
	items := make([]ArticleItem, 0)
	for _, article := range articles {
		item := ArticleItem{
			Title:       cdata{article["title"]},
			Description: cdata{article["desc"]},
			PicURL:      cdata{article["picurl"]},
			URL:         cdata{article["url"]},
		}

		items = append(items, item)
	}

	data := PassiveMessageNews{
		ToUserName:   cdata{to},
		FromUserName: cdata{t.certificate["wxid"]},
		MsgType:      cdata{"news"},
		CreateTime:   time.Now().Local().Unix(),
		ArticleCount: num,
		Articles:     items,
	}

	return t.getXMLStringOfMessage(data)
}

// getXMLStringOfMessage 获取 xml 字串
func (t *Client) getXMLStringOfMessage(message interface{}) ([]byte, error) {
	res, err := xml.MarshalIndent(message, "", "")
	if err != nil {
		logger.Error("转换消息xml内容失败: ", err.Error())
		return nil, err
	}

	return res, nil
}
