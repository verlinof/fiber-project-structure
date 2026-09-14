package pkg_validation

import (
	"fmt"
	"mime/multipart"
)

type FileType []string

var (
	ImageFileType FileType = FileType{"image/jpeg", "image/jpg", "image/png"}
)

func FileValidation(fileHeader *multipart.FileHeader, fileType FileType) error {
	contentType := fileHeader.Header.Get("Content-Type")

	for _, typefile := range fileType {
		if typefile == contentType {
			return nil
		}
	}

	return fmt.Errorf("file type '%s' is not allowed", contentType)
}
