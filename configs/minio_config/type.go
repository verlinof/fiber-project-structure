package minio_config

type MinioConfig struct {
	Endpoint        string `env:"MINIO_ABSOLUTE_URL"`
	DockerURL       string `env:"DOCKER_MINIO_URL"`
	AccessKeyID     string `env:"MINIO_ROOT_USER"`
	SecretAccessKey string `env:"MINIO_ROOT_PASSWORD"`
	UseSSL          bool   `env:"MINIO_USE_SSL"`
	PublicBucket    string `env:"MINIO_PUBLIC_BUCKET"`
	PrivateBucket   string `env:"MINIO_PRIVATE_BUCKET"`
}
