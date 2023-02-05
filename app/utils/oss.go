package utils

import (
	"app/constant"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"mime/multipart"
)

var OssClient *oss.Client

func OssUpload(file *multipart.FileHeader, fileName string) (string, error) {
	bucket, err := OssClient.Bucket(constant.OssBucketUrl)
	if err != nil {
		return "", handleError(err)
	}

	src, err := file.Open()
	if err != nil {
		return "", handleError(err)
	}
	defer src.Close()

	if err := bucket.PutObject(fileName, src); err != nil {
		return "", handleError(err)
	}

	url, err := bucket.SignURL(fileName, oss.HTTPGet, 600)
	if err != nil {
		return "", handleError(err)
	}
	return url, nil
}

func handleError(err error) error {
	log.Println(err)
	return err
}
