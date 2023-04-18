package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var response []string
var listener net.Listener
var wg sync.WaitGroup

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

// init
func init() {
	var err error
	log.Printf("-------rfid监听初始化成功-------")
	listener, err = net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}

}

// ParseData 读卡器数据解析
func ParseData() []string {
	var re []string
	var l = len(response)
	if l < 5 {
		return []string{} //, errors.New("Parse data fail,check received data ")
	}
	log.Print("命令长度:", l)

	//序列个数
	n, _ := strconv.ParseInt(response[4], 16, 64)
	log.Print("序列个数:", n)
	if int(5+2*n) > l {
		return []string{}
	}

	//数据段
	body := response[5 : l-2]
	fmt.Println("body size:", len(body))

	//数据段获取
	for i := 0; i < len(body); i++ {
		var Str string
		//uid长度
		length, _ := strconv.ParseInt(body[i], 16, 64)
		fmt.Println("uid长度:", length)
		//i=0 j=1 len=12 end=13
		for j := i + 1; j <= int(length)+i; j++ {
			Str += body[j]
		}
		//录入结果
		re = append(re, Str)
		//跳过检索段落
		i += int(length)
	}
	//CRC校验
	return re //, nil
}

// GetID uid获取接口
func GetID(c *gin.Context) {
	//var code int
	//var data string

	fmt.Println("----try to connect...----")
	//wait connect
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("accept fail,err: %v\n", err)
		//continue
	} else {
		fmt.Println("----access----")
	}

	//创建协程进行收发数据
	wg.Add(1)
	go process(conn)

	//等待接受数据
	fmt.Println("---receiving...----")
	wg.Wait()

	//判断返回
	if response == nil {
		fmt.Println("---receive information fail---")
		c.JSON(200, gin.H{"data": "---receive information fail---"})
	} else {
		c.JSON(200, gin.H{"data": ParseData()})
	}

	fmt.Println("after number of go:", runtime.NumGoroutine())
}

// 监听协程
func process(conn net.Conn) {
	defer func() {
		remoteAddr := conn.RemoteAddr().String()
		log.Print("discard remove add:", remoteAddr)
		conn.Close()
		fmt.Println("pre number of go:", runtime.NumGoroutine())
		wg.Done()
	}()
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	fmt.Println("----access to con----")
	for {
		src := []byte{0x04, 0x00, 0x01, 0xdb, 0x4b}
		//发送数据
		if _, err := conn.Write(src); err != nil {
			fmt.Printf("write to client failed, err: %v\n", err)
			break
		}

		//读取数据
		var buf [128]byte
		len, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("Read" + err.Error())
			break
		} else if len == 0 {
			continue
		}
		//fmt.Println("data len:", len)
		var res []string
		for i := 0; i < len; i++ {
			res = append(res, fmt.Sprintf("%02x", buf[i]))
		}
		fmt.Println(res)

		response = res
		break

	}
}
