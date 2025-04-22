package controllers

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/hhdms/msjx/internal/models"
)

// 对象存储服务配置
const (
	AccessKey        = "6x32d2ca"
	SecretKey        = "vcn2vrkxxlt77xsd"
	InternalEndpoint = "object-storage.objectstorage-system.svc.cluster.local"
	ExternalEndpoint = "objectstorageapi.hzh.sealos.run"
	BucketName       = "6x32d2ca-msjx"
	UseSSL           = true
)

// UploadController 文件上传控制器
type UploadController struct {
	minioClient *minio.Client
}

// NewUploadController 创建文件上传控制器实例
func NewUploadController() *UploadController {
	// 初始化MinIO客户端
	minioClient, err := minio.New(ExternalEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKey, SecretKey, ""),
		Secure: UseSSL,
	})

	if err != nil {
		fmt.Printf("初始化MinIO客户端失败: %v\n", err)
		return &UploadController{}
	}

	// 检查存储桶是否存在，如果不存在则创建
	exists, err := minioClient.BucketExists(context.Background(), BucketName)
	if err != nil {
		fmt.Printf("检查存储桶是否存在失败: %v\n", err)
	}

	if !exists {
		err = minioClient.MakeBucket(context.Background(), BucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Printf("创建存储桶失败: %v\n", err)
		} else {
			fmt.Printf("存储桶 %s 创建成功\n", BucketName)
		}
	}

	return &UploadController{
		minioClient: minioClient,
	}
}

// Upload 处理文件上传请求
func (c *UploadController) Upload(ctx *gin.Context) {
	// 获取上传的文件
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Code: 0,
			Msg:  "获取上传文件失败",
			Data: nil,
		})
		return
	}
	defer file.Close()

	// 生成唯一的文件名
	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// 上传文件到对象存储
	contentType := header.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// 上传文件到MinIO
	_, err = c.minioClient.PutObject(context.Background(), BucketName, fileName, file, header.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Code: 0,
			Msg:  fmt.Sprintf("上传文件失败: %v", err),
			Data: nil,
		})
		return
	}

	// 构建文件访问URL
	fileURL := fmt.Sprintf("https://%s/%s/%s", ExternalEndpoint, BucketName, fileName)

	// 返回成功响应
	ctx.JSON(http.StatusOK, models.Response{
		Code: 1,
		Msg:  "success",
		Data: fileURL,
	})
}
