package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	log.Println("started oss sync task")
	log.Println("OSS Go SDK Version: ", oss.Version)

	taskCron := os.Getenv("TASK_CRON")
	ossEndpoint := os.Getenv("OSS_ENDPOINT")
	ossKeyId := os.Getenv("OSS_KEY_ID")
	ossKeySecret := os.Getenv("OSS_KEY_SECRET")
	ossBucketName := os.Getenv("OSS_BUCKET_NAME")
	ossPrefix := os.Getenv("OSS_PREFIX")
	syncDir := "/sync"
	log.Println("TASK_CRON:", taskCron)
	log.Println("ossEndpoint:", ossEndpoint)
	log.Println("ossKeyId:", ossKeyId)
	log.Println("ossKeySecret:", ossKeySecret)
	log.Println("ossBucketName:", ossBucketName)
	log.Println("sync dir:", syncDir)

	// 初始化oss
	// 创建OSSClient实例。
	client, err := oss.New(ossEndpoint, ossKeyId, ossKeySecret)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(ossBucketName)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(-1)
	}

	c := cron.New()
	c.AddFunc(taskCron, func() {
		handlerSync(ossPrefix, syncDir, *bucket)
	})

	c.Start()
	for {
		time.Sleep(time.Second)
	}
}

func handlerSync(prefix, syncDir string, bucket oss.Bucket) {
	log.Println("handler sync echo dir")
	var files []string

	err := filepath.Walk(syncDir, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		log.Println(file)
		ossUpload(prefix, strings.Replace(file, syncDir, "", 1), file, bucket)
	}
}

func ossUpload(prefix, path, localPath string, bucket oss.Bucket) {
	log.Println("uploading...,", localPath)
	objName := prefix+path
	isExist, err := bucket.IsObjectExist(objName)
    if err != nil {
        log.Println("Error:", err)
        return
    }
    if isExist {
        log.Println("upload exist:",objName)
        return
    }
	err2 := bucket.PutObjectFromFile(objName, localPath)
	if err2 != nil {
		log.Println("Error:", err2)
		return
	}
	log.Println("upload finished,", localPath,"obj name:",objName)
}
