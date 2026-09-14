package pkg_minio

import (
	"log"

	"github.com/verlinof/fiber-project-structure/configs/minio_config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func Init() {
	var err error

	endpoint := minio_config.Config.DockerURL
	accessKeyID := minio_config.Config.AccessKeyID
	secretAccessKey := minio_config.Config.SecretAccessKey
	useSSL := true

	// 1. Inisialisasi minio client object.
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Koneksi ke MinIO berhasil!")
}

func GetMinioClient() *minio.Client {
	return minioClient
}
