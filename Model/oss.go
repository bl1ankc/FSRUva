package Model

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

/*
func CreateQRCode(id string) string {
	filename := "./img/qrcode/" + id + ".png"
	err := qrcode.WriteFile(id, qrcode.Medium, 256, filename)
	if err != nil {
		return ""
	}
	if !UploadImgToOSS("qrcode/"+id+".png", filename) {
		fmt.Println("OSS上传失败")
	}

	return id
}

*/

func UploadImgToOSS(cloudfilepath string, file io.Reader) bool {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		fmt.Println("oss上传失败1:", err)
		return false
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("unknownx")
	if err != nil {
		fmt.Println("oss上传失败2:", err)
		return false
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	//err = bucket.PutObjectFromFile("exampledir/exampleobject.txt", "D:\\localpath\\examplefile.txt")
	err = bucket.PutObject(cloudfilepath, file)
	if err != nil {
		fmt.Println("oss上传失败3:", err)
		return false
	}
	return true
}

// GetPicUrl 获取临时url
func GetPicUrl(filename string) (string, bool) {
	// 从STS服务获取安全令牌（SecurityToken）。
	token, flag := GetSTSToken()
	if !flag {
		return "", false
	}
	// 绑定客户端
	client, err := oss.New(Endpoint, token.AccessKeyId, token.AccessKeySecret, oss.SecurityToken(token.SecurityToken))
	if err != nil {
		fmt.Println("获取临时url失败1:", err.Error())
		return "", false
	}
	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("unknownx")
	if err != nil {
		fmt.Println("获取临时url失败2:", err.Error())
		return "", false
	}
	// 获取签名url
	response, err := bucket.SignURL("img/"+filename, oss.HTTPGet, 60)
	if err != nil {
		fmt.Println("获取临时url失败3:", err.Error())
		return "", false
	}
	return response, true
}

func GetSTSToken() (sts.Credentials, bool) {

	//构建一个阿里云客户端, 用于发起请求。
	//构建阿里云客户端时，需要设置AccessKey ID和AccessKey Secret。
	client, err := sts.NewClientWithAccessKey("cn-shenzhen", AccessKeyID, AccessKeySecret)
	if err != nil {
		fmt.Print("获取STS临时密钥失败1：", err.Error())
	}

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。关于参数含义和设置方法，请参见《API参考》。
	request.RoleArn = "acs:ram::1761650696847549:role/tmp"
	request.RoleSessionName = "go"
	request.DurationSeconds = "900"

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		fmt.Print("获取STS临时密钥失败2：", err.Error())
		return response.Credentials, false
	}
	return response.Credentials, true
}
