package Model

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/skip2/go-qrcode"
)

var (
	Endpoint        = "https://oss-cn-shenzhen.aliyuncs.com"
	AccessKeyID     = ""
	AccessKeySecret = ""
)

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

func UploadImgToOSS(cloudfilepath string, filepath string) bool {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("unknownx")
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	//err = bucket.PutObjectFromFile("exampledir/exampleobject.txt", "D:\\localpath\\examplefile.txt")
	err = bucket.PutObjectFromFile(cloudfilepath, filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return true
}
