package minio_config

import "github.com/caarlos0/env/v8"

var Config *MinioConfig

func LoadConfig() *MinioConfig {
	cfg := new(MinioConfig)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
