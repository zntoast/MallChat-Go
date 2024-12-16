package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
)

func SaveUploadedFile(file *multipart.FileHeader, uploadDir string) (string, error) {
	// 创建上传目录
	err := os.MkdirAll(uploadDir, 0755)
	if err != nil {
		return "", err
	}

	// 生成文件名
	ext := path.Ext(file.Filename)
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// 按日期分目录
	dateDir := time.Now().Format("2006/01/02")
	fullDir := path.Join(uploadDir, dateDir)
	err = os.MkdirAll(fullDir, 0755)
	if err != nil {
		return "", err
	}

	dst := path.Join(fullDir, fileName)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}

	return path.Join(dateDir, fileName), nil
}
