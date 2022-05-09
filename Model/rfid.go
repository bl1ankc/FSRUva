package Model

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// GetRFID 获取RFID
func GetRFID() (string, bool, []string) {
	var uids []string
	//发送请求
	res, err := http.Get("https://www.jaychan.work/api/getId")
	if err != nil {
		fmt.Println("获取RFID失败 发送请求：", err.Error())
		return "发送请求失败", false, uids
	}

	//结束请求
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("获取RFID失败 结束请求：", err.Error())
		}
	}(res.Body)

	if res.StatusCode == 202 {
		return "没有读到标签", false, uids
	}

	//转换数据
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("获取RFID失败 转换数据：", err.Error())
		return "转换数据失败", false, uids
	}

	//绑定数据
	type Values struct {
		Value string `json:"value"`
	}
	type Resopnse struct {
		Code int      `json:"code"`
		Data []Values `json:"data"`
		Msg  string   `json:"msg"`
	}
	var resopnse Resopnse

	err = json.Unmarshal(body, &resopnse)
	if err != nil {
		fmt.Println("获取RFID失败 绑定数据：", err.Error())
		return "绑定数据失败", false, uids
	}
	//返回数据
	for _, uid := range resopnse.Data {
		uids = append(uids, uid.Value)
	}

	return "获取成功", true, uids
}
