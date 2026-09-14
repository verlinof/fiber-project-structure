package pkg_minio

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/verlinof/fiber-project-structure/configs/minio_config"
	"github.com/minio/minio-go/v7"
)

// UploadFile mengunggah file dari path lokal ke MinIO
func UploadFile(ctx context.Context, minioClient *minio.Client, bucketName, objectName string, src multipart.File, fileSize int64, contentType string) (minio.UploadInfo, error) {
	info, err := minioClient.PutObject(ctx, bucketName, objectName, src, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return info, nil
}

// DeleteFile menghapus file dari MinIO
func DeleteFile(ctx context.Context, minioClient *minio.Client, bucketName, objectName string) error {
	// Remove public/ in Objectname
	objectName = strings.TrimPrefix(objectName, fmt.Sprintf("%s/", minio_config.Config.PublicBucket))

	err := minioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

// GetPrivateObject generates a presigned URL for a private object
func GetPrivateObject(ctx context.Context, minioClient *minio.Client, bucketName, objectName string) (string, error) {
	// Hapus prefix "private/" jika ada di objectName
	objectName = strings.TrimPrefix(objectName, fmt.Sprintf("%s/", minio_config.Config.PrivateBucket))

	expiry := 30 * time.Minute

	presignedURL, err := minioClient.PresignedGetObject(ctx, bucketName, objectName, expiry, nil)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
