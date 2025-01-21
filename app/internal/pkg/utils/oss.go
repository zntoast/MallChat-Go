package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSSClient struct {
	client       *oss.Client
	bucket       *oss.Bucket
	bucketDomain string
}

func NewOSSClient(endpoint, accessKey, accessSecret, bucketName, bucketDomain string) (*OSSClient, error) {
	client, err := oss.New(endpoint, accessKey, accessSecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return &OSSClient{
		client:       client,
		bucket:       bucket,
		bucketDomain: bucketDomain,
	}, nil
}

// UploadFile 上传文件
func (o *OSSClient) UploadFile(fileBytes []byte, fileName string) (string, error) {
	// 生成OSS对象名
	ext := strings.ToLower(filepath.Ext(fileName))
	objectKey := fmt.Sprintf("avatars/%d%s", time.Now().UnixNano(), ext)

	// 上传文件
	err := o.bucket.PutObject(objectKey, strings.NewReader(string(fileBytes)))
	if err != nil {
		return "", err
	}

	// 返回可访问的URL
	return fmt.Sprintf("https://%s/%s", o.bucketDomain, objectKey), nil
}

// DeleteFile 删除文件
func (o *OSSClient) DeleteFile(fileURL string) error {
	// 从URL中提取对象名
	objectKey := strings.TrimPrefix(fileURL, fmt.Sprintf("https://%s/", o.bucketDomain))
	return o.bucket.DeleteObject(objectKey)
}

// ValidateFileType 验证文件类型
func (o *OSSClient) ValidateFileType(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
	return allowedExts[ext]
}
