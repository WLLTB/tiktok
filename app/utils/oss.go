package utils

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"mime/multipart"
	"tiktok/app/constant"
)

var OssClient *oss.Client

func OssUpload(file *multipart.FileHeader, fileName string) (string, error) {
	bucket, err := OssClient.Bucket(constant.OssBucketUrl)
	if err != nil {
		log.Println(err)
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer src.Close()

	// 上传文件
	err = bucket.PutObject(fileName, src)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// 获取访问地址
	url, err := bucket.SignURL(fileName, oss.HTTPGet, 600)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return url, nil
}
