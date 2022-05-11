package Model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	WXAccessToken string
	WXMESSAGE     map[string]string
)

const (
	APPID     = ""
	APPSECRET = ""
)

// InitWXMessage 初始化模板ID
func InitWXMessage() {
	WXMESSAGE = make(map[string]string)
	WXMESSAGE["RemindUserReturnUav"] = ""
	WXMESSAGE["RemindScheduleOK"] = ""
	WXMESSAGE["RemindCheckOK"] = ""
	WXMESSAGE["RemindAdminCheck"] = ""
}

// GetWXAccessToken 获取微信accesstoken
func GetWXAccessToken() {
	//发送请求
	res, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + APPID + "&secret=" + APPSECRET)
	if err != nil {
		fmt.Println("获取微信accesstoken失败 发送请求：", err.Error())
		return
	}
	//结束请求
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("获取微信accesstoken失败 结束请求：", err.Error())
		}
	}(res.Body)

	//解析数据
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("获取微信accesstoken请求关闭失败 解析数据：", err.Error())
		return
	}

	//转换数据
	type AccessToken struct {
		AccessToken string `json:"access_Token"`
		ExpiresIn   int64  `json:"expires_in"`
		Errcode     int    `json:"errcode"`
		Errmsg      string `json:"errmsg"`
	}

	//绑定数据
	var accesstoken AccessToken
	err = json.Unmarshal(body, &accesstoken)
	if err != nil {
		fmt.Println("获取微信accesstoken请求关闭失败 解析数据：", err.Error())
		return
	}

	//错误处理
	if accesstoken.Errcode != 0 {
		fmt.Println("获取微信accesstoken请求关闭失败 错误处理：", accesstoken.Errmsg)
	}

	//赋值
	WXAccessToken = accesstoken.AccessToken

}

// GetOpenid 获取用户OPENID和UNIONID
func GetOpenid(CODE string) (bool, string, string) {
	//发送请求
	res, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + APPSECRET + "&js_code=" + CODE + "&grant_type=authorization_code")
	if err != nil {
		fmt.Println("获取用户OPENID和UNIONID 发送请求：", err.Error())
		return false, "", ""
	}

	//结束请求
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("获取用户OPENID和UNIONID 结束请求：", err.Error())
		}
	}(res.Body)

	//解析数据
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("获取用户OPENID和UNIONID 解析数据：", err.Error())
		return false, "", ""
	}

	//转换数据
	type UserInfo struct {
		Openid    string `json:"openid"`
		ExpiresIn int64  `json:"session_key"`
		Unionid   string `json:"unionid"`
		Errcode   int    `json:"errcode"`
		Errmsg    string `json:"errmsg"`
	}

	//绑定数据
	var userinfo UserInfo
	err = json.Unmarshal(body, &userinfo)
	if err != nil {
		fmt.Println("获取用户OPENID和UNIONID 解析数据：", err.Error())
		return false, "", ""
	}

	//错误处理
	if userinfo.Errcode != 0 {
		fmt.Println("获取用户OPENID和UNIONID 错误处理：", userinfo.Errmsg)
		return false, "", ""
	}

	return true, userinfo.Openid, userinfo.Unionid
}

// SendMessage 发送通知
func SendMessage(Openid string, ID string, Page string, Data interface{}) bool {

	//定义结构体
	type Request struct {
		//AccessToken string      `json:"access_token"`
		Touser     string      `json:"touser"`
		TemplateId string      `json:"template_id"`
		Page       string      `json:"page"`
		Data       interface{} `json:"data"`
	}

	//绑定数据
	requestData := Request{
		//AccessToken: WXAccessToken,
		Touser:     Openid,
		TemplateId: WXMESSAGE[ID],
		Page:       Page,
		Data:       Data,
	}
	url := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + WXAccessToken

	//转换数据
	requestByte, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("发送通知失败 转换数据：", err.Error())
		return false
	}

	//增加请求
	requset, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestByte))
	if err != nil {
		fmt.Println("发送通知失败 增加请求：", err.Error())
		return false
	}

	requset.Header.Set("Content-Type", "application/json;charset=UTF-8")

	//发起请求
	client := http.Client{}
	response, err := client.Do(requset)
	if err != nil {
		fmt.Println("发送通知失败 发起请求：", err.Error())
		return false
	}

	//结束请求
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("发送通知失败 结束请求：", err.Error())
		}
	}(response.Body)

	//解析数据
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("发送通知失败 解析数据：", err.Error())
		return false
	}

	//绑定数据
	type Response struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}
	var responseData Response
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("发送通知失败 绑定数据：", err.Error())
		return false
	}

	return true

}
