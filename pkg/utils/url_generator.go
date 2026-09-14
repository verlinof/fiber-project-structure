package pkg_utils

import (
	"fmt"

	"github.com/verlinof/fiber-project-structure/configs/minio_config"
)

func GenerateFileUrl(objectName string) string {
	if objectName == "" {
		return ""
	}

	return fmt.Sprintf("%s/%s", minio_config.Config.Endpoint, objectName)
}
