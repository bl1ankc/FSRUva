package utils

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"io"
	"main/Const"
)

var obsClient, _ = obs.New(Const.AK, Const.SK, Const.Endpoint)

func UploadFileToObs(fileName string, data io.Reader) (string, error) {
	input := &obs.PutObjectInput{}
	input.Bucket = Const.BucketName
	input.Key = fileName
	input.Body = data
	output, err := obsClient.PutObject(input)

	if err == nil {
		fmt.Printf("RequestId:%s\n", output.RequestId)
		fmt.Printf("ETag:%s\n", output.ETag)
		return Const.Endpoint[:8] + Const.BucketName + "." + Const.Endpoint[8:] + "/" + fileName, nil
	} else if obsError, ok := err.(obs.ObsError); ok {
		fmt.Printf("Code:%s\n", obsError.Code)
		fmt.Printf("Message:%s\n", obsError.Message)
		return "", err
	}
	return "", err
}
