package bilibiliapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

//LoginReply 储存扫码登录返回的json
type LoginReply struct {
	Code   int           `json:"code"`
	Status bool          `json:"status"`
	Ts     time.Duration `json:"ts"`
	Data   struct {
		URL      string `json:"url"`
		OauthKey string `json:"oauthKey"`
	} `json:"data"`
}

//LoginInfo 登录状态
type LoginInfo struct {
	Code    int           `json:"code"`
	Status  bool          `json:"status"`
	Ts      time.Duration `json:"ts"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data"`
	DaTa    struct {
		URL string `json:"url"`
	} `json:"data"`
}

//GetLoginURL 	申请二维码URL及扫码秘钥;
//*LoginReply	返回的json解析后得到的struct;
//error			其他内部错误
func GetLoginURL() (*LoginReply, error) {
	//申请二维码URL及扫码秘钥
	resp, err := http.Get("http://passport.bilibili.com/qrcode/getLoginUrl")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//提取body的内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//将body的内容解析成struct
	miao := &LoginReply{}
	err = json.Unmarshal(body, miao)
	if err != nil {
		return nil, err
	}
	return miao, nil
}

//GetLoginInfo  验证二维码登录
//bool			登录状态;
//string		登录错误信息/cookie;
//error			内部其他错误;
func GetLoginInfo(oauthKey string) (bool, string, error) {
	body := "{\"oauthKey\":" + oauthKey + "\"}"
	resp, err := http.Post("http://passport.bilibili.com/qrcode/getLoginInfo", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return false, "", err
	}
	defer resp.Body.Close()
	//提取body的内容
	Body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, "", err
	}
	miao := &LoginInfo{}
	err = json.Unmarshal(Body, miao)
	if err != nil {
		return false, "", err
	}
	if miao.Status != true {
		return miao.Status, miao.Message, nil
	}
	return miao.Status, "登陆成功~", nil
}
